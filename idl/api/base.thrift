namespace go hewo.tikshop.route.base

struct BaseResponse {
    1: i64 code;
    2: string message;  // 描述信息
}

struct MessageResponse {
    1: string message, // Return status description
}

struct NilResponse {}

// user begin::
enum UserStatus {
    DELETED = 0,
    ACTIVE = 1,
    BANNED = 2,
}

// 定义公共的用户信息结构
struct User {
    1: i64    id; // 用户 ID
    2: string username; // 用户名
    3: string email; // 邮箱
    4: string role;
    5: UserStatus status;
}

struct Customer {
    1: string address;
    2: string phone;
}

struct Merchant {
    1: string address;
    2: string shop_name;
}

struct Admin {
    1: i32 level;
}
// user end::

// product begin::
struct Product {
    1: i64    id,
    2: i64    merchant_id,
    3: string name,
    4: string description,
    5: i64    price,        // 价格(分为单位)
    6: i64    stock,
    7: i8     status,       // 0=删除, 1=上架, 2=下架, 3=售罄
}
// product end::

// cart begin::
// ========== 购物车项结构(独立定义) ==========
struct CartItem {
    1: i64 cart_item_id,
    2: i64 product_id,
    3: i64 merchant_id,
    4: i64 quantity,
    5: i8 selected,
    6: Product product,  // 商品详情
}

struct MerchantGroup {
    1: i64 merchant_id,
    2: list<CartItem> items,
    3: i64 subtotal,
}
// cart end::

// order begin::
struct OrderItem {
    1: i64 productId
    2: i64 quantity
    3: i64 price
}

struct Order { 
    1: i64 orderId
    2: i64 userId
    3: string status
    4: i64 totalAmount
    5: string createdAt
    6: list<OrderItem> items
}

struct Address {
    1: string street
    2: string city
    3: string postalCode
    4: string country
}

struct PaymentDetails {
    1: string cardNumber 
    2: string expiryDate 
    3: string cvv 
}
// order end::