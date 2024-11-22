namespace go hewo.tikshop.route.order

include "./base.thrift"

struct GetOrdersRequest {}

struct GetOrderInfoRequest {
    1: i64 orderId (api.path = "orderId", api.vd = "$>0")
}

struct PostOrderRequest {
    1: list<base.CartItem> items (api.body = "items")
    2: base.Address Address (api.body = "address")
    3: string paymentMethod (api.body = "paymentMethod")
}

struct PostOrderResponse {
    1: i64 orderId,
    2: string message
}

struct PayOrderRequest {
    1: i64 orderId (api.path = "orderId", api.vd = "$>0")
    2: string paymentMethod (api.body = "paymentMethod")
    3: base.PaymentDetails paymentDetails (api.body = "paymentDetails")
}

struct CancelOrderRequest {
    1: i64 orderId (api.path = "orderId", api.vd = "$>0")
}

service OrderService {
    list<base.Order> getOrders(1: GetOrdersRequest req) (api.get="/api/orders")
    base.Order getOrderInfo(1: GetOrderInfoRequest req) (api.get="/api/order/:orderId")
    PostOrderResponse postOrder(1: PostOrderRequest req) (api.post = "/api/order")
    base.MessageResponse payOrder(1: PayOrderRequest req) (api.post = "/api/order/:orderId/pay")
    base.MessageResponse cancelOrder(1: CancelOrderRequest req) (api.post = "/api/orders/:orderId/cancel")
}