package com.cloudshop.user.service.impl;

import com.cloudshop.user.service.JwtTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import java.time.Duration;
import java.util.concurrent.TimeUnit;

@Service
public class JwtTokenServiceImpl implements JwtTokenService {

    @Autowired
    private RedisTemplate<String, Object> redisTemplate;

    // Redis key 前缀
    private static final String REFRESH_TOKEN_PREFIX = "refresh_token:";
    private static final String BLACKLIST_PREFIX = "blacklist:";
    private static final String SESSION_PREFIX = "session:";
    
    // 过期时间
    private static final long REFRESH_TOKEN_EXPIRE_DAYS = 30;
    private static final long SESSION_EXPIRE_HOURS = 24;

    @Override
    public void storeRefreshToken(String userId, String refreshToken) {
        String key = REFRESH_TOKEN_PREFIX + userId;
        redisTemplate.opsForValue().set(key, refreshToken, REFRESH_TOKEN_EXPIRE_DAYS, TimeUnit.DAYS);
    }

    @Override
    public boolean isValidRefreshToken(String userId, String refreshToken) {
        String key = REFRESH_TOKEN_PREFIX + userId;
        Object storedToken = redisTemplate.opsForValue().get(key);
        return storedToken != null && storedToken.toString().equals(refreshToken);
    }

    @Override
    public void revokeRefreshToken(String userId) {
        String key = REFRESH_TOKEN_PREFIX + userId;
        redisTemplate.delete(key);
    }

    @Override
    public void blacklistToken(String token, long expirationTime) {
        String key = BLACKLIST_PREFIX + token;
        // 计算剩余过期时间
        long remainingTime = Math.max(0, expirationTime - System.currentTimeMillis());
        if (remainingTime > 0) {
            redisTemplate.opsForValue().set(key, "1", remainingTime, TimeUnit.MILLISECONDS);
        }
    }

    @Override
    public boolean isTokenBlacklisted(String token) {
        String key = BLACKLIST_PREFIX + token;
        return Boolean.TRUE.equals(redisTemplate.hasKey(key));
    }

    @Override
    public void storeUserSession(String userId, String username, String role) {
        String key = SESSION_PREFIX + userId;
        String sessionInfo = String.format("{\"username\":\"%s\",\"role\":\"%s\"}", username, role);
        redisTemplate.opsForValue().set(key, sessionInfo, SESSION_EXPIRE_HOURS, TimeUnit.HOURS);
    }

    @Override
    public String getUserSession(String userId) {
        String key = SESSION_PREFIX + userId;
        Object session = redisTemplate.opsForValue().get(key);
        return session != null ? session.toString() : null;
    }

    @Override
    public void clearUserSession(String userId) {
        String key = SESSION_PREFIX + userId;
        redisTemplate.delete(key);
    }
}
