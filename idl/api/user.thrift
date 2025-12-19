namespace go hewo.tikshop.route.user

include "./base.thrift"

enum UserStatus {
    DELETED = 0,
    ACTIVE = 1,
    BANNED = 2,
}

// 定义公共的用户信息结构
struct User {
    1: i64    id; // 用户 ID
    2: string username; // 用户名
    3: string email; // 邮箱
    4: string role;
    5: UserStatus status;
}

struct Customer {
    1: string address;
    2: string phone;
}

struct Merchant {
    1: string address;
    2: string shop_name;
}

struct Admin {
    1: i32 level;
}

// ========== Common APIs ==========

// Register
struct RegisterRequest {
    1: required string username (api.body="username", api.vd="len($) > 0 && len($) <= 50; msg:'用户名长度必须在1-50之间'");
    2: required string password (api.body="password", api.vd="len($) >= 8 && len($) <= 100; msg:'密码长度必须在8-100之间'");
    3: required string email (api.body="email", api.vd="len($) > 0 && len($) <= 255 && regexp('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$', $); msg:'邮箱格式不正确'");
    4: required string role (api.body="role", api.vd="$ == 'customer' || $ == 'merchant' || $ == 'admin'; msg:'角色必须是customer/merchant/admin'");

    // 可选的扩展字段，根据 role 决定
    5: optional string address (api.body="address", api.vd="len($) <= 512; msg:'地址长度不能超过512'");
    6: optional string phone (api.body="phone", api.vd="len($) <= 20 && regexp('^[0-9+-]+$', $); msg:'手机号格式不正确'");
    7: optional string shop_name (api.body="shop_name", api.vd="len($) > 0 && len($) <= 255; msg:'店铺名称长度必须在1-255之间'");
    8: optional i32 level (api.body="level", api.vd="$ >= 1 && $ <= 10; msg:'管理员等级必须在1-10之间'");
}

struct RegisterResponse {
    1: base.BaseResponse base;
    2: optional i64 user_id;
}

// Login
struct LoginRequest {
    1: string email (api.body="email", api.vd="len($) > 0 && regexp('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$', $); msg:'邮箱格式不正确'");
    2: string password (api.body="password", api.vd="len($) >= 8 && len($) <= 100; msg:'密码长度必须在8-100之间'");
}

struct LoginResponse {
    1: base.BaseResponse base;
    2: optional string token;
}

struct GetUserInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
}

struct GetUserInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional User user;
}

struct UpdateUserRequest {
    1: User user (api.body="user");
}

struct UpdateUserResponse {
    1: base.BaseResponse base;
    2: optional User user;
}

struct DeleteUserRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
}

struct DeleteUserResponse {
    1: base.BaseResponse base;
    2: optional bool success;
}

// ========== Customer APIs ==========
struct GetCustomerInfoByIDRequest {
    1: optional i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
}

struct GetCustomerInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional User user;
    3: optional Customer customer;
}

struct UpdateCustomerInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
    2: optional string address (api.body="address", api.vd="len($) <= 512; msg:'地址长度不能超过512'");
    3: optional string phone (api.body="phone", api.vd="len($) <= 20 && regexp('^[0-9+-]+$', $); msg:'手机号格式不正确'");
}

struct UpdateCustomerInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional Customer customer;
}

// ========== Merchant APIs ==========
struct GetMerchantInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
}

struct GetMerchantInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional User user;
    3: optional Merchant merchant;
}

struct UpdateMerchantInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
    2: optional string address (api.body="address", api.vd="len($) <= 512; msg:'地址长度不能超过512'");
    3: optional string shop_name (api.body="shop_name", api.vd="len($) > 0 && len($) <= 255; msg:'店铺名称长度必须在1-255之间'");
}

struct UpdateMerchantInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional Merchant merchant;
}

// ========== Admin APIs ==========

struct GetAdminInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
}

struct GetAdminInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional User user;
    3: optional Admin admin;
}

struct UpdateAdminInfoByIDRequest {
    1: i64 user_id (api.path="user_id", api.vd="$ > 0; msg:'用户ID必须大于0'");
    2: optional i32 level (api.body="level", api.vd="$ >= 1 && $ <= 10; msg:'管理员等级必须在1-10之间'");
}

struct UpdateAdminInfoByIDResponse {
    1: base.BaseResponse base;
    2: optional i32 level;
}

struct ListUsersRequest {
    1: required i32 page_number (api.query="page_number", api.vd="$ > 0; msg:'页码必须大于0'");
    2: required i32 page_size (api.query="page_size", api.vd="$ > 0 && $ <= 100; msg:'每页大小必须在1-100之间'");
}

struct ListUsersResponse {
    1: base.BaseResponse base;
    2: optional list<User> users;
    3: optional i32 total_count;
}


service UserService {
    RegisterResponse Register(1: RegisterRequest req) (api.post="/auth/register");
    LoginResponse Login(1: LoginRequest req) (api.post="/auth/login");

    GetUserInfoByIDResponse GetUserInfoByID(1: GetUserInfoByIDRequest req) (api.get="/user/:user_id");
    UpdateUserResponse UpdateUser(1: UpdateUserRequest req) (api.patch="/user/:user_id");
    DeleteUserResponse DeleteUser(1: DeleteUserRequest req) (api.delete="/user/:user_id");

    GetCustomerInfoByIDResponse GetCustomerInfoByID(1: GetCustomerInfoByIDRequest req) (api.get="/customer/:user_id");
    UpdateCustomerInfoByIDResponse UpdateCustomerInfoByID(1: UpdateCustomerInfoByIDRequest req) (api.patch="/customer/:user_id");

    GetMerchantInfoByIDResponse GetMerchantInfoByID(1: GetMerchantInfoByIDRequest req) (api.get="/merchant/:user_id");
    UpdateMerchantInfoByIDResponse UpdateMerchantInfoByID(1: UpdateMerchantInfoByIDRequest req) (api.patch="/merchant/:user_id");

    GetAdminInfoByIDResponse GetAdminInfoByID(1: GetAdminInfoByIDRequest req) (api.get="/admin/:user_id");
    UpdateAdminInfoByIDResponse UpdateAdminInfoByID(1: UpdateAdminInfoByIDRequest req) (api.patch="/admin/:user_id");

    ListUsersResponse ListUsers(1: ListUsersRequest req) (api.get="/users");
}