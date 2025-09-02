package com.cloudshop.user.service;

public interface VerificationCodeService {
    
    /**
     * 生成验证码
     * @return 6位数字验证码
     */
    String generateVerificationCode();
    
    /**
     * 存储验证码到Redis
     * @param email 邮箱
     * @param verificationCode 验证码
     */
    void storeVerificationCode(String email, String verificationCode);
    
    /**
     * 验证验证码
     * @param email 邮箱
     * @param verificationCode 验证码
     * @return 是否验证成功
     */
    boolean verifyCode(String email, String verificationCode);
    
    /**
     * 删除验证码
     * @param email 邮箱
     */
    void removeVerificationCode(String email);
}
