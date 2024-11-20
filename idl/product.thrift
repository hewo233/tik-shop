namespace go hewo.tikshop.product

include "base.thrift"
// 商品信息结构
struct Product {
    1: i64    id, // 商品ID
    2: string name, // 商品名称
    3: double price, // 商品价格
    4: i64    stock, // 库存数量
    5: string description // 商品描述
}

// 分页请求参数
struct GetProductsRequest {
    1: i64 page  = 1, // 页码，默认值为1
    2: i64 limit = 10 // 每页数量，默认值为10
}

struct GetProductsReqsponse {
    1: list<Product> products;
}

struct GetProductByIdRequest {
    1: i64 id;
}

struct GetProductByIdResponse {
    1: Product product;
}

// 商品创建请求结构
struct CreateProductRequest {
    1: string name, // 商品名称
    2: double price, // 商品价格
    3: i64    stock, // 库存数量
    4: string description // 商品描述
}

// 商品创建响应结构
struct CreateProductResponse {
    1: string message, // 成功消息
    2: i64    productId // 创建的商品ID
}

// 商品更新请求结构
struct UpdateProductRequest {
    1: optional string name, // 商品名称（可选）
    2: optional double price, // 商品价格（可选）
    3: optional i64    stock, // 库存数量（可选）
    4: optional string description // 商品描述（可选）
    5: i64      id
}

// 商品更新响应结构
struct UpdateProductResponse {
    1: string message // 成功消息
}

// 商品删除响应结构
struct DeleteProductRequset {
    1: i64 id;
}

struct DeleteProductResponse {
    1: string message // 成功消息
}

// 商品服务接口
service ProductService {
    // 获取商品列表，支持分页
    GetProductsReqsponse getProducts(1: GetProductsRequest request) throws (1: base.ErrorResponse error),

    // 获取单个商品详情
    GetProductByIdResponse getProductById(1: GetProductByIdRequest request) throws (1: base.ErrorResponse error),

    // 添加新商品（管理员权限）
    CreateProductResponse createProduct(1: CreateProductRequest request) throws (1: base.ErrorResponse error),

    // 更新商品信息（管理员权限）
    UpdateProductResponse updateProduct(1: UpdateProductRequest request) throws (1: base.ErrorResponse error),

    // 删除商品（管理员权限）
    DeleteProductResponse deleteProduct(1: DeleteProductRequset request) throws (1: base.ErrorResponse error)
}