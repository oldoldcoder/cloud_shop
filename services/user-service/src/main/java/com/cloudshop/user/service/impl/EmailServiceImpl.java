package com.cloudshop.user.service.impl;

import com.cloudshop.user.service.EmailService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@Service
public class EmailServiceImpl implements EmailService {
    
    private static final Logger logger = LoggerFactory.getLogger(EmailServiceImpl.class);
    
    @Autowired
    private JavaMailSender mailSender;

    @Value("${spring.mail.username:}")
    private String fromAddress;
    
    @Override
    public void sendVerificationCode(String to, String verificationCode) {
        try {
            SimpleMailMessage message = new SimpleMailMessage();
            message.setTo(to);
            if (fromAddress != null && !fromAddress.isEmpty()) {
                message.setFrom(fromAddress);
            }
            message.setSubject("Cloud Shop - 密码重置验证码");
            message.setText(String.format(
                "您好！\n\n" +
                "您正在重置 Cloud Shop 账户的密码。\n\n" +
                "验证码：%s\n\n" +
                "验证码有效期为10分钟，请尽快使用。\n\n" +
                "如果这不是您的操作，请忽略此邮件。\n\n" +
                "Cloud Shop 团队", 
                verificationCode
            ));
            
            mailSender.send(message);
            logger.info("验证码邮件发送成功: {}", to);
        } catch (Exception e) {
            logger.error("验证码邮件发送失败: {} 错误: {}", to, e.getMessage());
            throw new RuntimeException("邮件发送失败: " + e.getMessage());
        }
    }
    
    @Override
    public void sendPasswordResetSuccess(String to, String username) {
        try {
            SimpleMailMessage message = new SimpleMailMessage();
            message.setTo(to);
            if (fromAddress != null && !fromAddress.isEmpty()) {
                message.setFrom(fromAddress);
            }
            message.setSubject("Cloud Shop - 密码重置成功");
            message.setText(String.format(
                "您好 %s！\n\n" +
                "您的 Cloud Shop 账户密码已成功重置。\n\n" +
                "如果这不是您的操作，请立即联系客服。\n\n" +
                "Cloud Shop 团队", 
                username
            ));
            
            mailSender.send(message);
            logger.info("密码重置成功邮件发送成功: {}", to);
        } catch (Exception e) {
            logger.error("密码重置成功邮件发送失败: {} 错误: {}", to, e.getMessage());
            // 这里不抛出异常，因为密码重置已经成功
        }
    }
}
