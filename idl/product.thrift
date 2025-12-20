namespace go hewo.tikshop.product

include "base.thrift"

// 商品信息结构
struct Product {
    1: i64    id,
    2: i64    merchant_id,
    3: string name,
    4: string description,
    5: i64    price,        // 价格(分为单位)
    6: i64    stock,
    7: i8     status,       // 0=删除, 1=上架, 2=下架, 3=售罄
}

// 创建商品请求
struct CreateProductRequest {
    1: required i64    merchant_id,
    2: required string name,
    3: optional string description,
    4: required i64    price,
    5: optional i64    stock = 0,
}

struct CreateProductResponse {
    1: Product product,
}

// 查询单个商品
struct GetProductRequest {
    1: required i64 product_id,
}

struct GetProductResponse {
    1: Product product,
}

// 更新商品
struct UpdateProductRequest {
    1: required i64    product_id,
    2: optional string name,
    3: optional string description,
    4: optional i64    price,
    5: optional i8     status,
}

struct UpdateProductResponse {
    1: Product product,
}

// 商品列表查询
struct ListProductsRequest {
    1: i64 merchant_id,
    2: i8  status,
    3: i64 page = 1,
    4: i64 page_size = 10,
}

struct ListProductsResponse {
    1: list<Product> products,
    2: i64 total,
}

// 删除商品
struct DeleteProductRequest {
    1: required i64 product_id,
}

struct DeleteProductResponse {
    1: bool success,
}

// 修改库存
struct ModifyStockRequest {
    1: required i64 product_id,
    2: required i64 delta,  // 正数增加,负数减少
}

struct ModifyStockResponse {
    1: i64 stock,  // 返回修改后的库存
}

// 商品服务
service ProductService {
    CreateProductResponse CreateProduct(1: CreateProductRequest req) throws (1: base.ErrorResponse err);
    GetProductResponse GetProduct(1: GetProductRequest req) throws (1: base.ErrorResponse err);
    UpdateProductResponse UpdateProduct(1: UpdateProductRequest req) throws (1: base.ErrorResponse err);
    ListProductsResponse ListProducts(1: ListProductsRequest req) throws (1: base.ErrorResponse err);
    DeleteProductResponse DeleteProduct(1: DeleteProductRequest req) throws (1: base.ErrorResponse err);
    ModifyStockResponse ModifyStock(1: ModifyStockRequest req) throws (1: base.ErrorResponse err);
}