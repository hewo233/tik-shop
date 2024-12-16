
# 临时 API 约定

先用这个 MD 描述下，因为现在项目啥都没有。后续会改用 Swagger 生成的。

GPT 加人工了。注意 `查询参数` 是 URL Param 就是那个 `/api/foo?bar=111` 这样的问号后面的东西。参数 `/api/user/:id` 这里 `:id` 就是参数，直接换成对应的值即可。

以下是电商项目 API 设计的扩展版本，加入了管理员功能、异常管理和错误码，并符合 RESTful API 设计规范。

### 1. 鉴权模块（Auth Module）
- **POST /api/auth/login**：用户登录，返回 JWT token。
  - 请求体：
    ```json5
    { "username": "string", "password": "string" }
    ```
  - 响应：
    ```json5
    { "token": "JWT string" }
    ```
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：用户名或密码错误。

- **POST /api/auth/admin/login**：管理员登录，返回 JWT token。
  - 请求体：
    ```json5
    { "username": "string", "password": "string" }
    ```
  - 响应：
    ```json5
    { "token": "JWT string" }
    ```
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：用户名或密码错误。
    - `403 Forbidden`：普通用户尝试管理员登录。

- **GET /api/auth/verify**：验证用户权限。
  - Headers：`Authorization: Bearer <token>`
  - 响应：
    ```json5
    { "authorized": true/false }
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。

### 2. 用户模块（User Module）
- **GET /api/user/:id**：获取用户信息。
  - Headers：`Authorization: Bearer <token>`
  - 响应：
    ```json5
    { 
        "id": "number", 
        "username": "string", 
        "email": "string", 
        "createdAt": "date" 
    }
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。

- **PUT /api/user/:id**：修改用户信息。
  - Headers：`Authorization: Bearer <token>`
  - 响应：
    ```json5
    {
        "id": "number", 
        "username": "string",
        "email": "string",
        "createdAt": "date" 
    }
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。

- **POST /api/user/register**：用户注册。
  - 请求体：
    ```json5
    {
        "id": "number", 
        "username": "string",
        "email": "string",
        "createdAt": "date" 
    }
    ```
  - 响应：
    ```json5
    { "message": "User registered successfully." }
    ```
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `409 Conflict`：用户已存在。

- **PUT /api/user/:id/password**：修改用户密码。
  - Headers：`Authorization: Bearer <token>`
  - 请求体：`{ "oldPassword": "string", "newPassword": "string" }`
  - 响应：`{ "message": "Password updated successfully." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：旧密码错误。

### 3. 商品模块（Product Module）
- **GET /api/products**：获取商品列表，支持分页。
  - 查询参数：`page=number`，`limit=number`
  - 响应：
    ```json5
   [
        { 
            "id": "number", 
            "name": "string", 
            "price": "number", 
            "stock": "number", 
            "description": "string" 
        }
    ]
    ```
  - 错误码：
    - `400 Bad Request`：请求参数不正确。

- **GET /api/product/:id**：获取单个商品的详细信息。
  - 参数：`id`（商品ID）
  - 响应：
    ```json5
    { 
        "id": "number", 
        "name": "string", 
        "price": "number", 
        "stock": "number", 
        "description": "string" 
    }
    ```
  - 错误码：
    - `404 Not Found`：商品不存在。

- **POST /api/product**（管理员）：添加新商品。
  - Headers：`Authorization: Bearer <token>`
  - 请求体：`{ "name": "string", "price": "number", "stock": "number", "description": "string" }`
  - 响应：`{ "message": "Product created successfully.", "productId": "number" }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足。

- **PUT /api/product/:id**（管理员）：更新商品信息。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`id`（商品ID）
  - 请求体：`{ "name": "string", "price": "number", "stock": "number", "description": "string" }`
  - 响应：`{ "message": "Product updated successfully." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足。
    - `404 Not Found`：商品不存在。

- **DELETE /api/product/:id**（管理员）：删除商品。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`id`（商品ID）
  - 响应：`{ "message": "Product deleted successfully." }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足。
    - `404 Not Found`：商品不存在。

### 4. 购物车模块（Cart Module）
- **GET /api/cart**：获取购物车商品列表。
  - Headers：`Authorization: Bearer <token>`
  - 响应：
    ```json5
    [
        { 
            "productId": "number", 
            "quantity": "number" 
        }
    ]
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。

- **POST /api/cart**：向购物车添加商品。
  - Headers：`Authorization: Bearer <token>`
  - 请求体：`{ "productId": "number", "quantity": "number" }`
  - 响应：`{ "message": "Product added to cart." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `404 Not Found`：商品不存在。
    - `409 Conflict`：商品已在购物车中。

- **PUT /api/cart/:productId**：更新购物车中商品的数量。
  - Headers：`Authorization: Bearer <token>`
  - 请求体：`{ "quantity": "number" }`
  - 响应：`{ "message": "Cart updated successfully." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `404 Not Found`：商品不存在于购物车中。

- **DELETE /api/cart/:productId**：从购物车中删除商品。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`productId`
  - 响应：`{ "message": "Product removed from cart." }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `404 Not Found`：商品不存在于购物车中。

- **DELETE /api/cart**：清空购物车。
  - Headers：`Authorization: Bearer <token>`
  - 响应：`{ "message": "roduct removed from cart." }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。


### 5. 订单模块（Order Module）
为了支持**直接购买商品**并**提交订单**的功能，可以设计一个新的 API，允许用户在购物车或选择商品后直接提交订单。该操作通常包括选择商品、设置配送地址（可选）、支付方式等信息，并生成订单。

- **POST /api/order**：直接购买商品并提交订单。
  - 请求头：`Authorization: Bearer <token>`
  - 请求体：
    ```json5
    {
      "items": [
        {
          "productId": "number",
          "quantity": "number"
        }
      ],
      "address": {
        "street": "string",
        "city": "string",
        "postalCode": "string",
        "country": "string"
      },
      "paymentMethod": "string" // 支付方式：如"credit_card", "paypal"等
    }
    ```
  - 响应：
    ```json
    {
      "orderId": "string", // 新生成的订单 ID
      "message": "Order placed successfully."
    }
    ```
  - 错误码：
    - `400 Bad Request`：请求参数无效（如商品库存不足）。
    - `401 Unauthorized`：用户未登录或 JWT token 无效。
    - `404 Not Found`：商品不存在或已下架。
    - `422 Unprocessable Entity`：商品库存不足，无法完成订单。
    - `500 Internal Server Error`：服务器错误。

- **POST /api/order/:orderId/pay**：支付订单
  - 请求头：`Authorization: Bearer <token>`
  - 参数：`orderId`
  - 请求体：
    ```json5
    {
      "paymentMethod": "string", // 支付方式：如"credit_card", "paypal"等
      "paymentDetails": { // 支付详情（根据支付方式而定）
        "cardNumber": "string",
    2   "expiryDate": "string",
        "cvv": "string"
      }
    }
    ```
  - 响应：
    ```json
    {
      "message": "Payment successful. Order is confirmed."
    }
    ```
  - 错误码：
    - `400 Bad Request`：请求参数不正确（如支付信息缺失）。
    - `401 Unauthorized`：用户未登录或 JWT token 无效。
    - `404 Not Found`：订单未找到。
    - `422 Unprocessable Entity`：支付失败（如支付信息无效）。
    - `500 Internal Server Error`：支付过程出错。

- **POST /api/orders/:orderId/cancel**：取消订单。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`orderId`
  - 响应：`{ "message": "Order cancelled successfully." }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `404 Not Found`：订单不存在。
    - `409 Conflict`：订单已支付或已取消。

- **GET /api/orders**：获取用户的订单列表。
  - Headers：`Authorization: Bearer <token>`
  - 响应：
    ```json5
    [
        { 
            "orderId": "number", 
            "status": "string", 
            "totalAmount": "number", 
            "createdAt": "date", 
            "items": [
                { 
                    "productId": "number", 
                    "quantity": "number", 
                    "price": "number"
                }
            ]
        }
    ]
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。

- **GET /api/orders/:orderId**：获取单个订单的详情。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`orderId`
  - 响应：
    ```json5
    { 
        "orderId": "number", 
        "status": "string", 
        "totalAmount": "number", 
        "createdAt": "date", 
        "items": [
            { 
                "productId": "number", 
                "quantity": "number", 
                "price": "number" 
            }
        ]
    }
    ```
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `404 Not Found`：订单不存在。

以下是管理员模块中查看所有用户订单记录的 API 完整设计：

### 6. 管理员模块（Admin Module）
- **GET /api/admin/orders**：查看所有用户的订单记录。
  - Headers：`Authorization: Bearer <token>`
  - 响应：`[{ "orderId": "number", "userId": "number", "status": "string", "totalAmount": "number", "createdAt": "date", "items": [{ "productId": "number", "quantity": "number", "price": "number" }] }]`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。
  
- **GET /api/admin/orders/:orderId**：查看指定订单的详细信息。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`orderId`
  - 响应：`{ "orderId": "number", "userId": "number", "status": "string", "totalAmount": "number", "createdAt": "date", "items": [{ "productId": "number", "quantity": "number", "price": "number" }] }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。
    - `404 Not Found`：订单不存在。

- **GET /api/admin/users**：获取所有用户信息。
  - Headers：`Authorization: Bearer <token>`
  - 响应：`[{ "userId": "number", "username": "string", "email": "string", "createdAt": "date" }]`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。

- **GET /api/admin/users/:userId**：查看指定用户的详细信息。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`userId`
  - 响应：`{ "userId": "number", "username": "string", "email": "string", "createdAt": "date", "lastLogin": "date" }`
  - 错误码：
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。
    - `404 Not Found`：用户不存在。

- **PUT /api/admin/users/:userId**：更新指定用户的个人信息。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`userId`
  - 请求体：`{ "username": "string", "email": "string" }`
  - 响应：`{ "message": "User information updated successfully." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。
    - `404 Not Found`：用户不存在。

- **PUT /api/admin/users/:userId/password**：更改指定用户的密码。
  - Headers：`Authorization: Bearer <token>`
  - 参数：`userId`
  - 请求体：`{ "newPassword": "string" }`
  - 响应：`{ "message": "User password updated successfully." }`
  - 错误码：
    - `400 Bad Request`：请求参数不正确。
    - `401 Unauthorized`：无效的 token。
    - `403 Forbidden`：权限不足，普通用户无权访问。
    - `404 Not Found`：用户不存在。

### 错误响应结构

为了统一 API 错误处理，每个错误响应应该为纯文本。即使用 `c.String(CODE, MESSAGE)` 返回


### 全局错误码
- **400 Bad Request**：请求参数错误。
- **401 Unauthorized**：未经授权，token 无效或缺失。
- **403 Forbidden**：无权限访问。
- **404 Not Found**：资源未找到。
- **409 Conflict**：操作冲突（如商品已在购物车中，订单已支付或取消）。
- **500 Internal Server Error**：服务器内部错误。

此设计确保了各模块的权限分隔，满足 RESTful 设计规范，并支持管理员的特定操作需求。
