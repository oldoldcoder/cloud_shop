package com.cloudshop.user.dto;

import jakarta.validation.constraints.NotBlank;

public class UserLoginRequest {
    
    @NotBlank(message = "用户名/邮箱/手机号不能为空")
    private String principal;
    
    @NotBlank(message = "密码不能为空")
    private String password;
    
    // Getters and Setters
    public String getPrincipal() {
        return principal;
    }
    
    public void setPrincipal(String principal) {
        this.principal = principal;
    }
    
    public String getPassword() {
        return password;
    }
    
    public void setPassword(String password) {
        this.password = password;
    }
}
