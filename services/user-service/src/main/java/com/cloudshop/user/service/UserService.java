package com.cloudshop.user.service;

import com.cloudshop.user.dto.*;

public interface UserService {
    
    /**
     * 用户注册
     */
    void register(UserRegisterRequest request);
    
    /**
     * 用户登录
     */
    UserLoginResponse login(UserLoginRequest request);
    
    /**
     * 刷新访问令牌
     */
    UserLoginResponse refreshToken(String refreshToken);
    
    /**
     * 用户登出
     */
    void logout(String refreshToken);
    
    /**
     * 忘记密码 - 发送重置邮件
     */
    void forgotPassword(ForgotPasswordRequest request);
    
    /**
     * 发送验证码
     */
    void sendVerificationCode(SendVerificationCodeRequest request);
    
    /**
     * 验证验证码
     */
    boolean verifyVerificationCode(VerifyCodeRequest request);
    
    /**
     * 重置密码
     */
    void resetPassword(ResetPasswordRequest request);
    
    /**
     * 获取当前用户信息
     */
    UserLoginResponse.UserInfo getCurrentUserInfo();
}
