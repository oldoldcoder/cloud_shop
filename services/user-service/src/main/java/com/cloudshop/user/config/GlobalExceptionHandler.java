package com.cloudshop.user.config;

import com.cloudshop.user.dto.ApiResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@RestControllerAdvice
public class GlobalExceptionHandler {

    private static final Logger log = LoggerFactory.getLogger(GlobalExceptionHandler.class);

    @ExceptionHandler(MethodArgumentNotValidException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ApiResponse<List<Map<String, String>>> handleValidation(MethodArgumentNotValidException ex) {
        log.warn("参数校验失败: {}", ex.getMessage());
        List<Map<String, String>> errors = ex.getBindingResult().getFieldErrors()
                .stream()
                .map(this::toMap)
                .collect(Collectors.toList());
        return ApiResponse.errorWithData(400, "参数校验失败", errors);
    }

    @ExceptionHandler(IllegalArgumentException.class)
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ApiResponse<Void> handleIllegalArg(IllegalArgumentException ex) {
        // 仅返回通用提示，详细原因记录在服务端日志
        log.warn("业务校验异常: {}", ex.getMessage());
        return ApiResponse.error(400, "请求不合法或业务校验未通过");
    }

    @ExceptionHandler(Exception.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public ApiResponse<Void> handleUnknown(Exception ex) {
        // 未知异常统一拦截，避免堆栈或敏感信息返回前端
        log.error("系统内部错误", ex);
        return ApiResponse.error(500, "系统开小差了，请稍后再试");
    }

    private Map<String, String> toMap(FieldError fe) {
        Map<String, String> m = new HashMap<>();
        m.put("field", fe.getField());
        m.put("message", fe.getDefaultMessage());
        return m;
    }
}
