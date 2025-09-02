package com.cloudshop.user.service.impl;

import com.cloudshop.user.service.VerificationCodeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.stereotype.Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Random;
import java.util.concurrent.TimeUnit;

@Service
public class VerificationCodeServiceImpl implements VerificationCodeService {
    
    private static final Logger logger = LoggerFactory.getLogger(VerificationCodeServiceImpl.class);
    
    @Autowired
    private StringRedisTemplate redisTemplate;
    
    private static final String VERIFICATION_CODE_PREFIX = "verification_code:";
    private static final int VERIFICATION_CODE_LENGTH = 6;
    private static final int VERIFICATION_CODE_EXPIRE_MINUTES = 10;
    
    @Override
    public String generateVerificationCode() {
        Random random = new Random();
        StringBuilder code = new StringBuilder();
        for (int i = 0; i < VERIFICATION_CODE_LENGTH; i++) {
            code.append(random.nextInt(10));
        }
        return code.toString();
    }
    
    @Override
    public void storeVerificationCode(String email, String verificationCode) {
        String key = VERIFICATION_CODE_PREFIX + email;
        redisTemplate.opsForValue().set(key, verificationCode, VERIFICATION_CODE_EXPIRE_MINUTES, TimeUnit.MINUTES);
        logger.info("验证码已存储到Redis: {} 有效期: {}分钟", email, VERIFICATION_CODE_EXPIRE_MINUTES);
    }
    
    @Override
    public boolean verifyCode(String email, String verificationCode) {
        String key = VERIFICATION_CODE_PREFIX + email;
        String storedCode = redisTemplate.opsForValue().get(key);
        
        if (storedCode != null && storedCode.equals(verificationCode)) {
            logger.info("验证码验证成功: {}", email);
            return true;
        } else {
            logger.warn("验证码验证失败: {} 输入: {} 存储: {}", email, verificationCode, storedCode);
            return false;
        }
    }
    
    @Override
    public void removeVerificationCode(String email) {
        String key = VERIFICATION_CODE_PREFIX + email;
        redisTemplate.delete(key);
        logger.info("验证码已从Redis删除: {}", email);
    }
}
