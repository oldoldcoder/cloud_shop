package com.cloudshop.user.aspect;

import com.cloudshop.user.domain.LoginLog;
import com.cloudshop.user.dto.UserLoginRequest;
import com.cloudshop.user.service.LoginLogService;
import jakarta.servlet.http.HttpServletRequest;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.time.LocalDateTime;

@Aspect
@Component
public class LoginLogAspect {

    @Autowired
    private LoginLogService loginLogService;

    @Autowired
    private HttpServletRequest request;

    @Pointcut("execution(* com.cloudshop.user.controller.UserController.login(..))")
    public void loginPointcut() {}

    @Around("loginPointcut()")
    public Object aroundLogin(ProceedingJoinPoint pjp) throws Throwable {
        Object result = null;
        Integer status = 0;
        String message = null;
        Long userId = null;
        try {
            result = pjp.proceed();
            status = 1;
            message = "login success";
        } catch (Throwable ex) {
            message = ex.getMessage();
            throw ex;
        } finally {
            LoginLog log = new LoginLog();
            log.setUserId(userId); // 如需记录用户ID，可在登录成功后从返回体或上下文中获取
            log.setLoginTime(LocalDateTime.now());
            log.setIpAddress(getClientIp(request));
            log.setUserAgent(request.getHeader("User-Agent"));
            log.setStatus(status);
            log.setMessage(message);
            loginLogService.record(log);
        }
        return result;
    }

    private String getClientIp(HttpServletRequest request) {
        String ip = request.getHeader("X-Forwarded-For");
        if (ip != null && ip.length() != 0 && !"unknown".equalsIgnoreCase(ip)) {
            return ip.split(",")[0].trim();
        }
        ip = request.getHeader("X-Real-IP");
        if (ip != null && ip.length() != 0 && !"unknown".equalsIgnoreCase(ip)) {
            return ip;
        }
        return request.getRemoteAddr();
    }
}
