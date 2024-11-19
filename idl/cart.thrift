namespace go hewo.tikshop.cart

// 公共的错误结构
struct ErrorResponse {
    1: i64    code,      // 错误码
    2: string message    // 错误信息
}

// 商品项结构（在购物车中）
struct CartItem {
    1: i64 productId,    // 商品ID
    2: i64 quantity      // 商品数量
}

// 获取购物车请求结构
struct GetCartRequest {
    1: i64 userId        // 用户ID
}

// 获取购物车响应结构
struct GetCartResponse {
    1: list<CartItem> items // 购物车商品列表
}

// 添加商品到购物车请求结构
struct AddToCartRequest {
    1: i64 userId,       // 用户ID
    2: i64 productId,    // 商品ID
    3: i64 quantity      // 商品数量
}

// 添加商品到购物车响应结构
struct AddToCartResponse {
    1: string message    // 成功消息
}

// 更新购物车商品数量请求结构
struct UpdateCartRequest {
    1: i64 userId,       // 用户ID
    2: i64 productId,    // 商品ID
    3: i64 quantity      // 更新后的商品数量
}

// 更新购物车商品数量响应结构
struct UpdateCartResponse {
    1: string message    // 成功消息
}

// 从购物车中移除商品请求结构
struct RemoveFromCartRequest {
    1: i64 userId,       // 用户ID
    2: i64 productId     // 商品ID
}

// 从购物车中移除商品响应结构
struct RemoveFromCartResponse {
    1: string message    // 成功消息
}

// 清空购物车请求结构
struct ClearCartRequest {
    1: i64 userId        // 用户ID
}

// 清空购物车响应结构
struct ClearCartResponse {
    1: string message    // 成功消息
}

// 购物车服务接口
service CartService {
    // 获取购物车列表
    GetCartResponse getCart(1: GetCartRequest request) throws (1: ErrorResponse error),

    // 添加商品到购物车
    AddToCartResponse addToCart(1: AddToCartRequest request) throws (1: ErrorResponse error),

    // 更新购物车中商品数量
    UpdateCartResponse updateCart(1: UpdateCartRequest request) throws (1: ErrorResponse error),

    // 从购物车中移除商品
    RemoveFromCartResponse removeFromCart(1: RemoveFromCartRequest request) throws (1: ErrorResponse error),

    // 清空购物车
    ClearCartResponse clearCart(1: ClearCartRequest request) throws (1: ErrorResponse error)
}
