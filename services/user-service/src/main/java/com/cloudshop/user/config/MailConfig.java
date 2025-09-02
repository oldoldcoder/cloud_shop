package com.cloudshop.user.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.JavaMailSenderImpl;

import java.util.Properties;

@Configuration
public class MailConfig {

    @Bean
    public JavaMailSender javaMailSender() {
        JavaMailSenderImpl sender = new JavaMailSenderImpl();
        // 163 邮箱 SMTP 配置
        sender.setHost("smtp.163.com");
        sender.setPort(465);  // 163 推荐 SSL 方式
        sender.setUsername("13679342025@163.com");
        sender.setPassword("LZRUs38T6F3sSL7v");  // 163 授权码
        
        Properties props = sender.getJavaMailProperties();
        props.put("mail.smtp.auth", true);
        props.put("mail.smtp.ssl.enable", true);  // 启用 SSL
        props.put("mail.smtp.ssl.trust", "smtp.163.com");  // 信任 163 SMTP 服务器
        props.put("mail.smtp.ssl.protocols", "TLSv1.2");  // 使用 TLS 1.2
        props.put("mail.smtp.connectiontimeout", 10000);  // 连接超时 10 秒
        props.put("mail.smtp.timeout", 10000);  // 读取超时 10 秒
        props.put("mail.smtp.writetimeout", 10000);  // 写入超时 10 秒
        
        return sender;
    }
}
