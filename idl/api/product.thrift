namespace go hewo.tikshop.route.product

include "./base.thrift"

struct GetProductsRequest {
  1: i64 page (api.query="page", api.vd = "$>0")
  2: i64 limit (api.query="limit", api.vd = "$>0")
}

struct GetProductRequest {
    1: i64 id (api.path = "id", api.vd = "$>0")
}

struct CreateProductRequest {
  1: string name (api.body = "name", api.form = "name", api.vd = "(len($) > 0 && len($) < 512)")
  2: i64 price (api.body = "price", api.form = "price", api.vd = "$>0")
  3: i64 stock (api.body = "stock", api.form = "stock", api.vd = "$>0")
  4: string description (api.body = "description", api.form = "description", api.vd = "(len($) > 0 && len($) < 4096)")
}

struct CreateProductResponse {
    1: string message
    2: i64 productId
}

struct UpdateProductRequest {
  1: i64 id (api.path = "id", api.vd = "$>0")
  2: string name (api.body = "name", api.form = "name", api.vd = "(len($) > 0 && len($) < 512)")
  3: i64 price (api.body = "price", api.form = "price", api.vd = "$>0")
  4: i64 stock (api.body = "stock", api.form = "stock", api.vd = "$>0")
  5: string description (api.body = "description", api.form = "description", api.vd = "(len($) > 0 && len($) < 4096)")
}

struct UpdateProductResponse {
    1: string message
    2: i64 productId
}

struct DeleteProductRequest {
    1: i64 id (api.path = "id", api.vd = "$>0")
}

struct DeleteProductResponse {
    1: string message
    2: i64 productId
}

service ProductService {
  list<base.Product> getProducts(1: GetProductsRequest req) (api.get="/api/products")
  base.Product getProduct(1: GetProductRequest req) (api.get="/api/product/:id")
  CreateProductResponse createProduct(1: CreateProductRequest req) (api.post="/api/products")
  UpdateProductResponse updateProduct(1: UpdateProductRequest req) (api.put="/api/product/:id")
  DeleteProductResponse deleteProduct(1: DeleteProductRequest req) (api.delete="/api/product/:id")
}
