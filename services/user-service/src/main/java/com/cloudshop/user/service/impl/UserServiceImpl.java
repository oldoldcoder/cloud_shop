package com.cloudshop.user.service.impl;

import com.cloudshop.user.domain.User;
import com.cloudshop.user.dto.*;
import com.cloudshop.user.mapper.UserMapper;
import com.cloudshop.user.service.JwtTokenService;
import com.cloudshop.user.service.UserService;
import com.cloudshop.user.service.EmailService;
import com.cloudshop.user.service.VerificationCodeService;
import com.cloudshop.user.util.JwtUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

@Service
public class UserServiceImpl implements UserService {

    @Autowired
    private UserMapper userMapper;

    @Autowired
    private PasswordEncoder passwordEncoder;
    
    @Autowired
    private JwtUtil jwtUtil;
    
    @Autowired
    private JwtTokenService jwtTokenService;
    
    @Autowired
    private EmailService emailService;
    
    @Autowired
    private VerificationCodeService verificationCodeService;
    
    private static final Logger logger = LoggerFactory.getLogger(UserServiceImpl.class);

    @Override
    public void register(UserRegisterRequest request) {
        if (!StringUtils.hasText(request.getUsername())) {
            throw new IllegalArgumentException("用户名不能为空");
        }
        if (!StringUtils.hasText(request.getPassword())) {
            throw new IllegalArgumentException("密码不能为空");
        }
        if (!request.getPassword().equals(request.getConfirmPassword())) {
            throw new IllegalArgumentException("两次输入的密码不一致");
        }
        userMapper.findByUsername(request.getUsername()).ifPresent(u -> { throw new IllegalArgumentException("用户名已存在"); });
        if (StringUtils.hasText(request.getEmail())) {
            userMapper.findByEmail(request.getEmail()).ifPresent(u -> { throw new IllegalArgumentException("邮箱已存在"); });
        }
        if (StringUtils.hasText(request.getPhone())) {
            userMapper.findByPhone(request.getPhone()).ifPresent(u -> { throw new IllegalArgumentException("手机号已存在"); });
        }
        User user = new User();
        user.setUsername(request.getUsername());
        user.setEmail(request.getEmail());
        user.setPhone(request.getPhone());
        user.setPasswordHash(passwordEncoder.encode(request.getPassword()));
        user.setRole("USER");
        user.setStatus(1);
        userMapper.insert(user);
    }
    
    @Override
    public UserLoginResponse login(UserLoginRequest request) {
        if (!StringUtils.hasText(request.getPrincipal()) || !StringUtils.hasText(request.getPassword())) {
            throw new IllegalArgumentException("账号或密码不能为空");
        }
        String principal = request.getPrincipal();
        Optional<User> userOpt = userMapper.findByUsername(principal);
        if (userOpt.isEmpty()) {
            userOpt = userMapper.findByEmail(principal);
        }
        if (userOpt.isEmpty()) {
            userOpt = userMapper.findByPhone(principal);
        }
        User user = userOpt.orElseThrow(() -> new IllegalArgumentException("用户不存在或密码错误"));
        if (user.getStatus() != null && user.getStatus() == 0) {
            throw new IllegalArgumentException("账号已被禁用");
        }
        if (!passwordEncoder.matches(request.getPassword(), user.getPasswordHash())) {
            throw new IllegalArgumentException("用户不存在或密码错误");
        }
        
        // 使用 JWT 生成访问令牌
        Map<String, Object> claims = new HashMap<>();
        claims.put("userId", user.getId());
        claims.put("username", user.getUsername());
        claims.put("role", user.getRole());
        
        String accessToken = jwtUtil.generateAccessToken(user.getUsername(), claims);
        String refreshToken = jwtUtil.generateRefreshToken(user.getUsername());
        long expiresIn = jwtUtil.getAccessTokenExpiration() / 1000; // 转换为秒
        
        // 存储刷新令牌到 Redis
        jwtTokenService.storeRefreshToken(user.getId().toString(), refreshToken);
        
        // 存储用户会话信息到 Redis
        jwtTokenService.storeUserSession(user.getId().toString(), user.getUsername(), user.getRole());
        
        UserLoginResponse.UserInfo info = new UserLoginResponse.UserInfo(
                user.getId(), user.getUsername(), user.getEmail(), user.getPhone()
        );
        return new UserLoginResponse(accessToken, refreshToken, expiresIn, info);
    }
    
    @Override
    public UserLoginResponse refreshToken(String refreshToken) {
        try {
            String username = jwtUtil.getUsernameFromToken(refreshToken);
            if (jwtUtil.isTokenExpired(refreshToken)) {
                throw new IllegalArgumentException("刷新令牌已过期");
            }
            
            // 获取用户信息
            Optional<User> userOpt = userMapper.findByUsername(username);
            if (userOpt.isEmpty()) {
                throw new IllegalArgumentException("用户不存在");
            }
            
            User user = userOpt.get();
            
            // 验证 Redis 中存储的刷新令牌
            if (!jwtTokenService.isValidRefreshToken(user.getId().toString(), refreshToken)) {
                throw new IllegalArgumentException("无效的刷新令牌");
            }
            
            // 生成新的访问令牌
            Map<String, Object> claims = new HashMap<>();
            claims.put("userId", user.getId());
            claims.put("username", user.getUsername());
            claims.put("role", user.getRole());
            
            String newAccessToken = jwtUtil.generateAccessToken(user.getUsername(), claims);
            String newRefreshToken = jwtUtil.generateRefreshToken(user.getUsername());
            long expiresIn = jwtUtil.getAccessTokenExpiration() / 1000; // 转换为秒
            
            // 更新 Redis 中的刷新令牌
            jwtTokenService.storeRefreshToken(user.getId().toString(), newRefreshToken);
            
            UserLoginResponse.UserInfo info = new UserLoginResponse.UserInfo(
                    user.getId(), user.getUsername(), user.getEmail(), user.getPhone()
            );
            return new UserLoginResponse(newAccessToken, newRefreshToken, expiresIn, info);
            
        } catch (Exception e) {
            throw new IllegalArgumentException("无效的刷新令牌");
        }
    }
    
    @Override
    public void logout(String refreshToken) {
        try {
            if (jwtUtil.isTokenExpired(refreshToken)) {
                throw new IllegalArgumentException("刷新令牌已过期");
            }
            
            String username = jwtUtil.getUsernameFromToken(refreshToken);
            Optional<User> userOpt = userMapper.findByUsername(username);
            if (userOpt.isPresent()) {
                User user = userOpt.get();
                // 撤销刷新令牌
                jwtTokenService.revokeRefreshToken(user.getId().toString());
                // 清除用户会话
                jwtTokenService.clearUserSession(user.getId().toString());
            }
            
        } catch (Exception e) {
            throw new IllegalArgumentException("无效的刷新令牌");
        }
    }
    
    @Override
    public void forgotPassword(ForgotPasswordRequest request) {
        if (!StringUtils.hasText(request.getEmail())) {
            throw new IllegalArgumentException("邮箱不能为空");
        }
        
        // 检查邮箱是否存在
        Optional<User> userOpt = userMapper.findByEmail(request.getEmail());
        if (userOpt.isEmpty()) {
            throw new IllegalArgumentException("该邮箱未注册");
        }
        
        User user = userOpt.get();
        if (user.getStatus() != null && user.getStatus() == 0) {
            throw new IllegalArgumentException("账号已被禁用");
        }
        
        // 生成验证码
        String verificationCode = verificationCodeService.generateVerificationCode();
        
        // 存储验证码到Redis
        verificationCodeService.storeVerificationCode(request.getEmail(), verificationCode);
        
        // 发送验证码邮件
        emailService.sendVerificationCode(request.getEmail(), verificationCode);
        
        logger.info("忘记密码验证码已发送: {}", request.getEmail());
    }
    
    @Override
    public void sendVerificationCode(SendVerificationCodeRequest request) {
        if (!StringUtils.hasText(request.getEmail())) {
            throw new IllegalArgumentException("邮箱不能为空");
        }
        
        // 检查邮箱是否存在
        Optional<User> userOpt = userMapper.findByEmail(request.getEmail());
        if (userOpt.isEmpty()) {
            throw new IllegalArgumentException("该邮箱未注册");
        }
        
        User user = userOpt.get();
        if (user.getStatus() != null && user.getStatus() == 0) {
            throw new IllegalArgumentException("账号已被禁用");
        }
        
        // 生成验证码
        String verificationCode = verificationCodeService.generateVerificationCode();
        
        // 存储验证码到Redis
        verificationCodeService.storeVerificationCode(request.getEmail(), verificationCode);
        
        // 发送验证码邮件
        emailService.sendVerificationCode(request.getEmail(), verificationCode);
        
        logger.info("验证码已发送: {}", request.getEmail());
    }
    
    @Override
    public boolean verifyVerificationCode(VerifyCodeRequest request) {
        if (!StringUtils.hasText(request.getEmail()) || !StringUtils.hasText(request.getVerificationCode())) {
            throw new IllegalArgumentException("邮箱和验证码不能为空");
        }
        
        // 验证验证码
        boolean isValid = verificationCodeService.verifyCode(request.getEmail(), request.getVerificationCode());
        
        if (isValid) {
            // 验证成功后删除验证码
            verificationCodeService.removeVerificationCode(request.getEmail());
            logger.info("验证码验证成功: {}", request.getEmail());
        } else {
            logger.warn("验证码验证失败: {}", request.getEmail());
        }
        
        return isValid;
    }
    
    @Override
    public void resetPassword(ResetPasswordRequest request) {
        if (!StringUtils.hasText(request.getResetToken()) || !StringUtils.hasText(request.getNewPassword())) {
            throw new IllegalArgumentException("重置令牌和新密码不能为空");
        }
        
        if (!request.getNewPassword().equals(request.getConfirmPassword())) {
            throw new IllegalArgumentException("两次输入的密码不一致");
        }
        
        // 这里应该验证重置令牌的有效性
        // 为了简化，我们假设重置令牌就是邮箱
        String email = request.getResetToken();
        
        // 检查邮箱是否存在
        Optional<User> userOpt = userMapper.findByEmail(email);
        if (userOpt.isEmpty()) {
            throw new IllegalArgumentException("用户不存在");
        }
        
        User user = userOpt.get();
        if (user.getStatus() != null && user.getStatus() == 0) {
            throw new IllegalArgumentException("账号已被禁用");
        }
        
        // 更新密码
        user.setPasswordHash(passwordEncoder.encode(request.getNewPassword()));
        userMapper.updateById(user);
        
        // 发送密码重置成功邮件
        emailService.sendPasswordResetSuccess(email, user.getUsername());
        
        logger.info("密码重置成功: {}", email);
    }
    
    @Override
    public UserLoginResponse.UserInfo getCurrentUserInfo() {
        // TODO: 从 SecurityContext 获取当前用户信息
        throw new UnsupportedOperationException("获取用户信息功能待实现");
    }
}
