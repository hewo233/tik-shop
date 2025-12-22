namespace go hewo.tikshop.route.product

include "./base.thrift"

// 获取单个商品
struct GetProductByIDRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0", api.body = '-');
}

struct GetProductByIDResponse {
    1: base.BaseResponse base;
    2: optional base.Product product;
}

// 创建商品
struct CreateProductRequest {
    1: string name (api.body = "name", api.vd = "len($) > 0 && len($) < 255");
    2: optional string description (api.body = "description");
    3: i64 price (api.body = "price", api.vd = "$>0");
    4: optional i64 stock (api.body = "stock");
}

struct CreateProductResponse {
    1: base.BaseResponse base;
    2: optional i64 product_id;
}

// 更新商品
struct UpdateProductByIDRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0", api.body = '-');
    2: optional string name (api.body = "name", api.vd = "len($) > 0 && len($) < 255");
    3: optional string description (api.body = "description");
    4: optional i64 price (api.body = "price", api.vd = "$>0");
    5: optional i8 status (api.body = "status", api.vd = "$>0 && $<3");
}

struct UpdateProductByIDResponse{
    1: base.BaseResponse base;
    2: optional base.Product product;
}

// 获取商品列表
struct ListProductsRequest {
    1: i64 merchant_id (api.query="merchant_id", api.vd = "$>0");
    2: i64 page (api.query="page", api.vd = "$>0");
    3: i64 page_size (api.query="page_size", api.vd = "$>0 && $<=100");
}

struct ListProductsResponse {
    1: base.BaseResponse base;
    2: optional list<base.Product> products;
    3: optional i64 total;
}

// 删除商品
struct DeleteProductByIDRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0", api.body = '-');
}

struct DeleteProductByIDResponse {
    1: base.BaseResponse base;
    2: optional bool success;
}

// 修改库存
struct ModifyStockByIDRequest {
    1: i64 product_id (api.path = "id", api.vd = "$>0", api.body = '-');
    2: i64 delta (api.body = "delta", api.vd = "$!=0");
    3: i64 currentStock (api.body = "current_stock", api.vd = "$>=0");
}

struct ModifyStockByIDResponse {
    1: base.BaseResponse base;
    2: optional i64 stock;
}

service ProductService {
    CreateProductResponse CreateProduct(1: CreateProductRequest req) (api.post="/product");
    GetProductByIDResponse GetProductByID(1: GetProductByIDRequest req) (api.get="/product/:id");
    UpdateProductByIDResponse UpdateProductByID(1: UpdateProductByIDRequest req) (api.put="/product/:id");
    ListProductsResponse ListProducts(1: ListProductsRequest req) (api.get="/product/list");
    DeleteProductByIDResponse DeleteProductByID(1: DeleteProductByIDRequest req) (api.delete="/product/:id");
    ModifyStockByIDResponse ModifyStockByID(1: ModifyStockByIDRequest req) (api.put="/product/:id/stock");
}