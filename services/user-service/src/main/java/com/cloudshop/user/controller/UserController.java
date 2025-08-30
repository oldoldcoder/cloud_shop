package com.cloudshop.user.controller;

import com.cloudshop.user.dto.*;
import com.cloudshop.user.service.UserService;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/users")
@CrossOrigin(origins = "*")
public class UserController {
    
    @Autowired
    private UserService userService;
    
    /**
     * 用户注册
     */
    @PostMapping("/register")
    public ApiResponse<String> register(@Valid @RequestBody UserRegisterRequest request) {
        try {
            userService.register(request);
            return ApiResponse.success("注册成功", "用户注册成功，请查收激活邮件");
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 用户登录
     */
    @PostMapping("/login")
    public ApiResponse<UserLoginResponse> login(@Valid @RequestBody UserLoginRequest request) {
        try {
            UserLoginResponse response = userService.login(request);
            return ApiResponse.success("登录成功", response);
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 刷新访问令牌
     */
    @PostMapping("/token/refresh")
    public ApiResponse<UserLoginResponse> refreshToken(@RequestParam String refreshToken) {
        try {
            UserLoginResponse response = userService.refreshToken(refreshToken);
            return ApiResponse.success("令牌刷新成功", response);
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 用户登出
     */
    @PostMapping("/logout")
    public ApiResponse<String> logout(@RequestParam String refreshToken) {
        try {
            userService.logout(refreshToken);
            return ApiResponse.success("登出成功", "用户已成功登出");
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 忘记密码 - 发送重置邮件
     */
    @PostMapping("/password/forgot")
    public ApiResponse<String> forgotPassword(@Valid @RequestBody ForgotPasswordRequest request) {
        try {
            userService.forgotPassword(request);
            return ApiResponse.success("重置邮件已发送", "密码重置链接已发送到您的邮箱，请查收");
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 重置密码
     */
    @PostMapping("/password/reset")
    public ApiResponse<String> resetPassword(@Valid @RequestBody ResetPasswordRequest request) {
        try {
            userService.resetPassword(request);
            return ApiResponse.success("密码重置成功", "密码已成功重置，请使用新密码登录");
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
    
    /**
     * 获取当前用户信息
     */
    @GetMapping("/profile")
    public ApiResponse<UserLoginResponse.UserInfo> getProfile() {
        try {
            UserLoginResponse.UserInfo userInfo = userService.getCurrentUserInfo();
            return ApiResponse.success("获取成功", userInfo);
        } catch (Exception e) {
            return ApiResponse.error(e.getMessage());
        }
    }
}
