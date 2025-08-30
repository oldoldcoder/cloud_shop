package com.cloudshop.user.service;

import com.cloudshop.user.domain.LoginLog;

public interface LoginLogService {
    void record(LoginLog log);
}
