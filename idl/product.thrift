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
    1: i64 product_id,
}

// 查询单个商品
struct GetProductByIDRequest {
    1: required i64 product_id,
}

struct GetProductByIDResponse {
    1: Product product,
}

// 更新商品
struct UpdateProductByIDRequest {
    1: required i64    product_id,
    2: required i64    merchant_id,
    3: optional string name,
    4: optional string description,
    5: optional i64    price,
    6: optional i8     status,
}

struct UpdateProductByIDResponse {
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
struct DeleteProductByIDRequest {
    1: required i64 product_id,
    2: required i64 merchant_id,
}

struct DeleteProductByIDResponse {
    1: bool success,
}

// 修改库存
struct ModifyStockByIDRequest {
    1: required i64 product_id,
    2: required i64 delta,  // 正数增加,负数减少
    3: required i64 currentStock, // 当前库存
    4: required i64 merchant_id,
}

struct ModifyStockByIDResponse {
    1: i64 stock,  // 返回修改后的库存
}

// 商品服务
service ProductService {
    CreateProductResponse CreateProduct(1: CreateProductRequest req) throws (1: base.ErrorResponse err);
    GetProductByIDResponse GetProductByID(1: GetProductByIDRequest req) throws (1: base.ErrorResponse err);
    UpdateProductByIDResponse UpdateProductByID(1: UpdateProductByIDRequest req) throws (1: base.ErrorResponse err);
    ListProductsResponse ListProducts(1: ListProductsRequest req) throws (1: base.ErrorResponse err);
    DeleteProductByIDResponse DeleteProductByID(1: DeleteProductByIDRequest req) throws (1: base.ErrorResponse err);
    ModifyStockByIDResponse ModifyStock(1: ModifyStockByIDRequest req) throws (1: base.ErrorResponse err);
}