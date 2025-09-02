Web 层：Gin

配置/存储：Viper、GORM(+mysql driver)、Redis

认证：JWT

消息：Kafka、RabbitMQ（二选一或混用，视业务而定）

文档：Swagger 生态三件套

驱动：MySQL 驱动给 GORM/database/sql 使用

如果你告诉我当前 product-service 的功能边界（是否真的需要 Kafka/RabbitMQ/Swagger），我可以给出最小化依赖建议与初始化样例。
