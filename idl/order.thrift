namespace go hewo.tikshop.order

include "base.thrift"
// 订单商品项
struct OrderItem {
    1: i64 productId,    // 商品ID
    2: i64 quantity,     // 商品数量
    3: double price      // 商品单价
}

// 地址结构
struct Address {
    1: string street,    // 街道
    2: string city,      // 城市
    3: string postalCode,// 邮编
    4: string country    // 国家
}

// 支付信息结构
struct PaymentDetails {
    1: string cardNumber, // 卡号（根据支付方式动态使用）
    2: string expiryDate, // 有效期
    3: string cvv         // CVV码
}

// 订单状态枚举
enum OrderStatus {
    PENDING,             // 待支付
    PAID,                // 已支付
    CANCELLED            // 已取消
}

// 提交订单请求结构
struct SubmitOrderRequest {
    1: i64 userId,       // 用户ID
    2: list<OrderItem> items, // 商品列表
    3: Address address,  // 配送地址
    4: string paymentMethod // 支付方式
}

// 提交订单响应结构
struct SubmitOrderResponse {
    1: i64 orderId,      // 新生成的订单ID
    2: string message    // 成功消息
}

// 支付订单请求结构
struct PayOrderRequest {
    1: i64 orderId,      // 订单ID
    2: string paymentMethod, // 支付方式
    3: optional PaymentDetails paymentDetails // 支付详情（可选）
}

// 支付订单响应结构
struct PayOrderResponse {
    1: string message    // 成功消息
}

// 取消订单请求结构
struct CancelOrderRequest {
    1: i64 orderId       // 订单ID
}

// 取消订单响应结构
struct CancelOrderResponse {
    1: string message    // 成功消息
}

// 获取订单列表请求结构
struct GetOrdersRequest {
    1: i64 userId        // 用户ID
}

// 获取订单列表响应结构
struct GetOrdersResponse {
    1: list<Order> orders // 订单列表
}

// 获取单个订单详情请求结构
struct GetOrderByIdRequest {
    1: i64 orderId       // 订单ID
}

// 获取单个订单详情响应结构
struct GetOrderByIdResponse {
    1: Order order       // 订单详情
}

// 订单结构
struct Order {
    1: i64 orderId,      // 订单ID
    2: OrderStatus status, // 订单状态
    3: double totalAmount, // 总金额
    4: string createdAt,  // 创建时间（ISO 8601 格式）
    5: list<OrderItem> items // 订单商品列表
}

// 订单服务接口
service OrderService {
    // 提交订单
    SubmitOrderResponse submitOrder(1: SubmitOrderRequest request) throws (1: base.ErrorResponse error),

    // 支付订单
    PayOrderResponse payOrder(1: PayOrderRequest request) throws (1: base.ErrorResponse error),

    // 取消订单
    CancelOrderResponse cancelOrder(1: CancelOrderRequest request) throws (1: base.ErrorResponse error),

    // 获取用户订单列表
    GetOrdersResponse getOrders(1: GetOrdersRequest request) throws (1: base.ErrorResponse error),

    // 获取单个订单详情
    GetOrderByIdResponse getOrderById(1: GetOrderByIdRequest request) throws (1: base.ErrorResponse error)
}
