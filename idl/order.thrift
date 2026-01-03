namespace go hewo.tikshop.order

include "base.thrift"

enum OrderStatus {
    PENDING_PAYMENT = 0, // 待支付
    PAID            = 1, // 已支付
    SHIPPED         = 2, // 已发货
    COMPLETED       = 3, // 已完成
    CANCELLED       = 4, // 已取消
    REFUNDED        = 5  // 已退款
    WAITING         = 6  // 下单中
}

struct OrderAddress {
    1: string customer_name,
    2: string phone,
    3: string address,
}

// 订单商品项
    struct OrderItem {
    1: i64    product_id,
    2: i64    quantity,
    3: i64    cost,          // 单价(分)
    4: string product_name,
    5: i64    merchant_id,   // 归属商家ID
    6: i64    total_cost,    // 小计(分)
}

// 订单主结构 (支持多商家混合)
struct Order {
    1: i64          id,
    2: i64          customer_id,
    4: OrderStatus  status,
    5: list<OrderItem> order_items,
    6: i64          total_amount,  // 总金额(分)
    7: OrderAddress    address,
    8: i64          created_at,
}

struct CreateOrderItem {
    1: i64    product_id,
    2: i64    quantity,
}

struct CreateOrderRequest {
    1: required i64 customer_id,
    2: required list<CreateOrderItem> items, // 列表内可包含不同商家的商品
}

struct CreateOrderResponse {
    1: i64 order_id,
}

struct ListOrdersRequest {
    1: required i64 customer_id,
    2: optional i32 status,
    3: i64 page = 1,
    4: i64 page_size = 10,
}

struct ListOrdersResponse {
    1: list<Order> orders,
    2: i64 total,
}

struct GetOrderRequest {
    1: required i64 order_id,
    2: required i64    customer_id,
}

struct GetOrderResponse {
    1: Order order,
}

struct CancelOrderRequest {
    1: required i64 order_id,
    2: required i64    customer_id,
}

struct CancelOrderResponse {
    1: bool success,
}

service OrderService {
    CreateOrderResponse CreateOrder(1: CreateOrderRequest req) throws (1: base.ErrorResponse err);
    ListOrdersResponse ListOrders(1: ListOrdersRequest req) throws (1: base.ErrorResponse err);
    GetOrderResponse GetOrder(1: GetOrderRequest req) throws (1: base.ErrorResponse err);
    CancelOrderResponse CancelOrder(1: CancelOrderRequest req) throws (1: base.ErrorResponse err);
}