namespace go hewo.tikshop.route.product

include "./base.thrift"

// 获取单个商品
struct GetProductRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0");
}

struct GetProductResponse {
    1: base.Product product;
}

// 创建商品
struct CreateProductRequest {
    1: i64 merchant_id (api.body = "merchant_id", api.vd = "$>0");
    2: string name (api.body = "name", api.vd = "len($) > 0 && len($) < 255");
    3: optional string description (api.body = "description");
    4: i64 price (api.body = "price", api.vd = "$>0");
    5: optional i64 stock (api.body = "stock");
}

struct CreateProductResponse {
    1: base.Product product;
}

// 更新商品
struct UpdateProductRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0");
    2: optional string name (api.body = "name", api.vd = "len($) > 0 && len($) < 255");
    3: optional string description (api.body = "description");
    4: optional i64 price (api.body = "price", api.vd = "$>0");
    5: optional i8 status (api.body = "status", api.vd = "$>=0 && $<=3");
}

struct UpdateProductResponse {
    1: base.Product product;
}

// 获取商品列表
struct ListProductsRequest {
    1: i64 merchant_id (api.query="merchant_id");
    2: i8  status (api.query="status");
    3: i64 page (api.query="page", api.vd = "$>0");
    4: i64 page_size (api.query="page_size", api.vd = "$>0 && $<=100");
}

struct ListProductsResponse {
    1: list<base.Product> products;
    2: i64 total;
}

// 删除商品
struct DeleteProductRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0");
}

struct DeleteProductResponse {
    1: bool success;
}

// 修改库存
struct ModifyStockRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0");
    2: i64 delta (api.body = "delta", api.vd = "$!=0");
}

struct ModifyStockResponse {
    1: i64 stock;
}

service ProductService {
    CreateProductResponse CreateProduct(1: CreateProductRequest req) (api.post="/product");
    GetProductResponse GetProduct(1: GetProductRequest req) (api.get="/product/:id");
    UpdateProductResponse UpdateProduct(1: UpdateProductRequest req) (api.put="/product/:id");
    ListProductsResponse ListProducts(1: ListProductsRequest req) (api.get="/product/list");
    DeleteProductResponse DeleteProduct(1: DeleteProductRequest req) (api.delete="/product/:id");
    ModifyStockResponse ModifyStock(1: ModifyStockRequest req) (api.put="/product/:id/stock");
}