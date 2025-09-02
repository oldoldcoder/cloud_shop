package com.cloudshop.user.service;

public interface EmailService {
    
    /**
     * 发送验证码邮件
     * @param to 收件人邮箱
     * @param verificationCode 验证码
     */
    void sendVerificationCode(String to, String verificationCode);
    
    /**
     * 发送密码重置成功邮件
     * @param to 收件人邮箱
     * @param username 用户名
     */
    void sendPasswordResetSuccess(String to, String username);
}
