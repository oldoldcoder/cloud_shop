package com.cloudshop.user.service.impl;

import com.cloudshop.user.domain.LoginLog;
import com.cloudshop.user.mapper.LoginLogMapper;
import com.cloudshop.user.service.LoginLogService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class LoginLogServiceImpl implements LoginLogService {

    @Autowired
    private LoginLogMapper loginLogMapper;

    @Override
    public void record(LoginLog log) {
        loginLogMapper.insert(log);
    }
}
