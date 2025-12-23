namespace go hewo.tikshop.cart

include "base.thrift"
include "product.thrift"

// 购物车项(带关联商品)
struct CartItem {
    1: i64                 cart_item_id,  // 对应 db.model.CartItem.ID
    2: i64                 product_id,
    3: i64                 merchant_id,   // 冗余字段,用于前端分组
    4: i64                 quantity,
    5: bool                  selected,      // 0=未选中 1=选中
    6: product.Product     product,       // Preload 出来的商品信息
}

// 商家分组(前端按店铺展示)
struct MerchantGroup {
    1: i64             merchant_id,
    2: list<CartItem>  items,
    3: i64             subtotal,          // 该商家已选中商品小计(分)
}

// ========== 获取购物车 ==========
struct GetCartRequest {
    1: required i64 customer_id,
}

struct GetCartResponse {
    1: list<MerchantGroup> groups,        // 按商家分组
    2: i64                 total_selected, // 已选中商品总价(分)
}

// ========== 添加到购物车 ==========
struct AddToCartRequest {
    1: required i64 customer_id,
    2: required i64 product_id,
    3: required i64 quantity = 1,
}

struct AddToCartResponse {
    1: i64 cart_item_id,
}

// ========== 更新数量 ==========
struct UpdateQuantityRequest {
    1: required i64 customer_id,
    2: required i64 cart_item_id,
    3: required i64 quantity,             // 0=删除该项
}

struct UpdateQuantityResponse {
    1: bool success,
}

// ========== 切换选中状态 ==========
struct ToggleSelectRequest {
    1: required i64      customer_id,
    2: required list<i64> cart_item_ids,  // 支持批量
    3: required bool selected,        // 0=取消 1=选中
}

struct ToggleSelectResponse {
    1: bool success,
}

// ========== 删除购物车项 ==========
struct RemoveItemsRequest {
    1: required i64      customer_id,
    2: required list<i64> cart_item_ids,
}

struct RemoveItemsResponse {
    1: bool success,
}

// ========== 清空购物车 ==========
struct ClearCartRequest {
    1: required i64 customer_id,
}

struct ClearCartResponse {
    1: bool success,
}

// ========== 购物车服务 ==========
service CartService {
    GetCartResponse GetCart(1: GetCartRequest req) throws (1: base.ErrorResponse err);
    AddToCartResponse AddToCart(1: AddToCartRequest req) throws (1: base.ErrorResponse err);
    UpdateQuantityResponse UpdateQuantity(1: UpdateQuantityRequest req) throws (1: base.ErrorResponse err);
    ToggleSelectResponse ToggleSelect(1: ToggleSelectRequest req) throws (1: base.ErrorResponse err);
    RemoveItemsResponse RemoveItems(1: RemoveItemsRequest req) throws (1: base.ErrorResponse err);
    ClearCartResponse ClearCart(1: ClearCartRequest req) throws (1: base.ErrorResponse err);
}