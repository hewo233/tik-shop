namespace go hewo.tikshop.route.cart

include "./base.thrift"

// ========== 获取购物车 ==========
struct GetCartRequest {}

struct GetCartResponse {
    1: base.BaseResponse base,
    2: optional list<base.MerchantGroup> groups,
    3: optional i64 total_selected,
}

// ========== 添加到购物车 ==========
struct AddToCartRequest {
    1: i64 product_id (api.body = "product_id", api.vd = "$>0"),
    2: optional i64 quantity (api.body = "quantity", api.vd = "$>0"),
}

struct AddToCartResponse {
    1: base.BaseResponse base,
    2: optional i64 cart_item_id,
}

// ========== 更新数量 ==========
struct UpdateQuantityRequest {
    1: i64 cart_item_id (api.path = "id", api.vd = "$>0", api.body = '-'),
    2: i64 quantity (api.body = "quantity", api.vd = "$>=0"),  // 0=删除
}

struct UpdateQuantityResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

// ========== 切换选中状态 ==========
struct ToggleSelectRequest {
    1: list<i64> cart_item_ids (api.body = "cart_item_ids"),
    2: i8 selected (api.body = "selected", api.vd = "$>=0 && $<=1"),
}

struct ToggleSelectResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

// ========== 删除购物车项 ==========
struct RemoveItemsRequest {
    1: list<i64> cart_item_ids (api.body = "cart_item_ids"),
}

struct RemoveItemsResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

// ========== 清空购物车 ==========
struct ClearCartRequest {}

struct ClearCartResponse {
    1: base.BaseResponse base,
    2: optional bool success,
}

// ========== 购物车服务 ==========
service CartService {
    GetCartResponse GetCart(1: GetCartRequest req) (api.get="/cart");
    AddToCartResponse AddToCart(1: AddToCartRequest req) (api.post="/cart");
    UpdateQuantityResponse UpdateQuantity(1: UpdateQuantityRequest req) (api.put="/cart/:id");
    ToggleSelectResponse ToggleSelect(1: ToggleSelectRequest req) (api.put="/cart/select");
    RemoveItemsResponse RemoveItems(1: RemoveItemsRequest req) (api.delete="/cart/items");
    ClearCartResponse ClearCart(1: ClearCartRequest req) (api.delete="/cart");
}