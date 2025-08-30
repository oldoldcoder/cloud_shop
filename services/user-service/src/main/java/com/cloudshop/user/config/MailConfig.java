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
        // 可由 application.yml 自动装配覆盖，这里给默认值避免空指针
        sender.setHost("smtp.example.com");
        sender.setPort(587);
        sender.setUsername("your_email@example.com");
        sender.setPassword("your_password");
        Properties props = sender.getJavaMailProperties();
        props.put("mail.smtp.auth", true);
        props.put("mail.smtp.starttls.enable", true);
        return sender;
    }
}
