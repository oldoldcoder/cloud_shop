package com.cloudshop.user.dto;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Email;

public class SendVerificationCodeRequest {
    
    @NotBlank(message = "邮箱不能为空")
    @Email(message = "邮箱格式不正确")
    private String email;
    
    // Getters and Setters
    public String getEmail() {
        return email;
    }
    
    public void setEmail(String email) {
        this.email = email;
    }
}
