package com.cloudshop.user.mapper;

import com.cloudshop.user.domain.User;
import org.apache.ibatis.annotations.Param;

import java.util.Optional;

public interface UserMapper {
    int insert(User user);

    Optional<User> findByUsername(@Param("username") String username);

    Optional<User> findByEmail(@Param("email") String email);

    Optional<User> findByPhone(@Param("phone") String phone);
}
