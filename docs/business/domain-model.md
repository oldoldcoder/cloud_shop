# 领域模型图

## 核心业务对象
- User（用户）
- Product（商品）
- Order（订单）
- Inventory（库存）
- Payment（支付）

## 关系图

```mermaid
classDiagram
    class User {
        +id
        +name
        +email
    }

    class Product {
        +id
        +name
        +price
    }

    class Order {
        +id
        +status
        +createdAt
    }

    class Inventory {
        +productId
        +stock
    }

    class Payment {
        +id
        +orderId
        +status
    }

    User "1" --> "many" Order
    Order "1" --> "many" Product : contains
    Order "1" --> Payment : has
    Product "1" --> Inventory : manages
