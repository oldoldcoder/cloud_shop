package com.cloudshop.gateway.filter;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.factory.AbstractGatewayFilterFactory;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.stereotype.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.crypto.SecretKey;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Arrays;
import java.util.List;
import java.util.TimeZone;

@Component
public class JwtAuthenticationFilter extends AbstractGatewayFilterFactory<JwtAuthenticationFilter.Config> {

    private static final Logger logger = LoggerFactory.getLogger(JwtAuthenticationFilter.class);

    @Value("${jwt.secret}")
    private String secret;

    // 定义公开路径，不需要验证 JWT
    private static final List<String> PUBLIC_PATHS = Arrays.asList(
        "/api/users/login",
        "/api/users/register", 
        "/api/users/password/forgot",
        "/api/users/password/reset",
        "/api/users/token/refresh",
        "/api/users/verification/send",
        "/api/users/verification/verify"
    );

    public JwtAuthenticationFilter() {
        super(Config.class);
    }

    @Override
    public GatewayFilter apply(Config config) {
        return (exchange, chain) -> {
            ServerHttpRequest request = exchange.getRequest();
            String path = request.getPath().value();
            String method = request.getMethod().name();
            String remoteAddress = request.getRemoteAddress() != null ? request.getRemoteAddress().getAddress().getHostAddress() : "unknown";
            
            // 🚀 网关接收请求日志
            logger.info("🚀 [网关] 收到请求: {} {} 来自: {} 时间: {}", 
                method, path, remoteAddress, formatDate(new Date()));
            
            // 检查是否是公开路径
            if (isPublicPath(path)) {
                logger.info("✅ [网关] 公开路径，直接放行: {} {}", method, path);
                return chain.filter(exchange);
            }
            
            logger.info("🔒 [网关] 需要认证的路径: {} {}", method, path);
            
            // 获取 Authorization 头
            String authHeader = request.getHeaders().getFirst(HttpHeaders.AUTHORIZATION);
            if (authHeader == null || !authHeader.startsWith("Bearer ")) {
                logger.warn("❌ [网关] 缺少或无效的Authorization头: {} {} 来自: {}", method, path, remoteAddress);
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }
            
            // 提取 JWT 令牌
            String token = authHeader.substring(7);
            logger.info("🔑 [网关] 验证JWT令牌: {} {} 令牌长度: {}", method, path, token.length());
            
            try {
                // 验证 JWT 令牌
                if (validateToken(token)) {
                    logger.info("✅ [网关] JWT令牌验证成功: {} {}", method, path);
                    return chain.filter(exchange);
                } else {
                    logger.warn("❌ [网关] JWT令牌验证失败: {} {}", method, path);
                    exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                    return exchange.getResponse().setComplete();
                }
            } catch (Exception e) {
                logger.error("💥 [网关] JWT令牌验证异常: {} {} 错误: {}", method, path, e.getMessage());
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }
        };
    }
    
    private boolean isPublicPath(String path) {
        boolean isPublic = PUBLIC_PATHS.stream().anyMatch(path::startsWith);
        if (isPublic) {
            logger.debug("🔓 [网关] 路径匹配公开规则: {}", path);
        }
        return isPublic;
    }
    
    private boolean validateToken(String token) {
        try {
            SecretKey key = Keys.hmacShaKeyFor(secret.getBytes());
            Claims claims = Jwts.parserBuilder()
                    .setSigningKey(key)
                    .build()
                    .parseClaimsJws(token)
                    .getBody();
            
            // 检查令牌是否过期
            boolean isValid = !claims.getExpiration().before(new Date());
            logger.debug("🔍 [网关] JWT令牌验证结果: 过期={} 有效期至={}", 
                !isValid, claims.getExpiration());
            return isValid;
        } catch (Exception e) {
            logger.error("💥 [网关] JWT令牌解析异常: {}", e.getMessage());
            return false;
        }
    }
    
    public static class Config {
        // 配置属性，如果需要的话
    }
    
    /**
     * 格式化日期为ISO格式字符串
     */
    private String formatDate(Date date) {
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
        sdf.setTimeZone(TimeZone.getTimeZone("UTC"));
        return sdf.format(date);
    }
}
