package com.cloudshop.user.service;

public interface JwtTokenService {
    
    /**
     * 存储刷新令牌到 Redis
     */
    void storeRefreshToken(String userId, String refreshToken);
    
    /**
     * 验证刷新令牌是否有效
     */
    boolean isValidRefreshToken(String userId, String refreshToken);
    
    /**
     * 撤销刷新令牌
     */
    void revokeRefreshToken(String userId);
    
    /**
     * 将访问令牌加入黑名单
     */
    void blacklistToken(String token, long expirationTime);
    
    /**
     * 检查令牌是否在黑名单中
     */
    boolean isTokenBlacklisted(String token);
    
    /**
     * 存储用户会话信息
     */
    void storeUserSession(String userId, String username, String role);
    
    /**
     * 获取用户会话信息
     */
    String getUserSession(String userId);
    
    /**
     * 清除用户会话
     */
    void clearUserSession(String userId);
}
