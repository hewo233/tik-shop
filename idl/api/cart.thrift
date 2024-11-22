namespace go hewo.tikshop.route.cart

include "./base.thrift"

struct MessageResponse {
    1: string message
}

struct GetCartRequest {}

struct AddToCartRequest {
    1: i64 productId (api.body = "productId", api.form = "productId", api.vd = "$>0")
    2: i64 quantity (api.body = "quantity", api.form = "quantity", api.vd = "$>0")
}

struct UpdateCartItemRequest {
    1: i64 productId (api.path = "productId", api.vd = "$>0")
    2: i64 quantity (api.body = "quantity", api.form = "quantity", api.vd = "$>0")
}

struct DeleteCartItemRequest {
    1: i64 productId (api.body = "productId", api.form = "productId", api.vd = "$>0")
}

struct DeleteCartRequest { }

service ProductService {
    list<base.CartItem> getCart(1: GetCartRequest req) (api.get="/api/cart")
    base.MessageResponse addToCart(1: AddToCartRequest req) (api.post="/api/cart")
    base.MessageResponse updateCartItem(1: UpdateCartItemRequest req) (api.put="/api/cart/:productId")
    base.MessageResponse deleteCartItem(1: DeleteCartItemRequest req) (api.delete="/api/cart/:productId")
    base.MessageResponse deleteCart(1: DeleteCartRequest req) (api.delete="/api/cart")
}
