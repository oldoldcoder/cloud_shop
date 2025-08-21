> 事件流（OrderCreated → InventoryReserved → PaymentPending → PaymentSuccess → OrderConfirmed） 就是电商系统最核心的 MVP 工作流，几乎所有分布式系统都会遇到：幂等、事务一致性、补偿机制、超时处理等。这个流是最佳练习场。

```mermaid
sequenceDiagram
    participant User
    participant OrderService
    participant InventoryService
    participant PaymentService

    User->>OrderService: Create Order
    OrderService->>InventoryService: Reserve Inventory
    InventoryService-->>OrderService: InventoryReserved
    OrderService->>PaymentService: Request Payment
    PaymentService-->>OrderService: PaymentSuccess
    OrderService-->>User: Order Confirmed
