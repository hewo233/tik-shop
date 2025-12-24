namespace go hewo.tikshop.route.order

include "./base.thrift"

// ========== 基础结构定义 (对应 RPC 层) ==========

enum OrderStatus {
    PENDING_PAYMENT = 0, // 待支付
    PAID            = 1, // 已支付
    SHIPPED         = 2, // 已发货
    COMPLETED       = 3, // 已完成
    CANCELLED       = 4, // 已取消
    REFUNDED        = 5  // 已退款
}

struct OrderAddress {
    1: string customer_name (api.body = "customer_name", api.vd = "len($)>0"),
    2: string phone (api.body = "phone", api.vd = "len($)>0"),
    3: string address (api.body = "address", api.vd = "len($)>0"),
}

struct OrderItem {
    1: i64 product_id (api.body = "product_id", api.vd = "$>0"),
    2: i64 quantity (api.body = "quantity", api.vd = "$>0"),
    3: i64 cost (api.body = "cost", api.vd = "$>0"), // 单价(分)
    4: string product_name (api.body = "product_name", api.vd = "len($)>0"),
    5: i64 merchant_id (api.body = "merchant_id", api.vd = "$>0"),
    6: i64 total_cost (api.body = "total_cost", api.vd = "$>0"), // 小计(分)
}

// API 返回的订单视图
struct Order {
    1: i64 id,
    2: i64 customer_id,
    3: OrderStatus status,
    4: list<OrderItem> items,
    5: i64 total_amount,
    6: OrderAddress address,
    7: i64 created_at,
}

struct CreateOrderItem {
    1: i64    product_id (api.body = "product_id", api.vd = "$>0"),
    2: i64    quantity (api.body = "quantity", api.vd = "$>0"),
}

// ========== 1. 下单接口 ==========
struct CreateOrderRequest {
    1: list<CreateOrderItem> items (api.body = "items", api.vd = "len($)>0"),
}

struct CreateOrderResponse {
    1: base.BaseResponse base,
    2: optional i64 order_id,
}

// ========== 2. 订单列表接口 ==========
struct ListOrdersRequest {
    1: i64 page (api.query = "page", api.vd = "$>0", default = "1"),
    2: i64 page_size (api.query = "page_size", api.vd = "$>0", default = "10"),
    3: optional i32 status (api.query = "status"),
}

struct ListOrdersResponse {
    1: base.BaseResponse base,
    2: list<OrderItem> orders,
    3: optional i64 total,
}

// ========== 3. 订单详情接口 ==========
struct GetOrderRequest {
    1: i64 id (api.path = "id", api.vd = "$>0"),
}

struct GetOrderResponse {
    1: base.BaseResponse base,
    2: optional OrderItem order,
}

// ========== 4. 取消订单接口 ==========
struct CancelOrderRequest {
    1: i64 id (api.path = "id", api.vd = "$>0"),
}

struct CancelOrderResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

struct MarkOrderPaidRequest {
    1: i64 id (api.path = "id", api.vd = "$>0"),
}

struct MarkOrderPaidResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

// ========== Service 定义 ==========
service OrderService {
    CreateOrderResponse PlaceOrder(1: CreateOrderRequest req) (api.post="/order");
    ListOrdersResponse ListOrders(1: ListOrdersRequest req) (api.get="/order");
    GetOrderResponse GetOrder(1: GetOrderRequest req) (api.get="/order/:id");
    CancelOrderResponse CancelOrder(1: CancelOrderRequest req) (api.post="/order/:id/cancel");
    MarkOrderPaidResponse MarkOrderPaid(1: MarkOrderPaidRequest req) (api.post="/order/:id/mark_paid");

}