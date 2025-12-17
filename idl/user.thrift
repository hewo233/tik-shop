namespace go hewo.tikshop.user

include "base.thrift"

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
    1: required string username,
    2: required string password,
    3: required string email,
    4: required string role,  // "customer", "merchant", "admin"

    // 可选的扩展字段，根据 role 决定
    5: optional string address,
    6: optional string phone,      // customer 使用
    7: optional string shop_name,  // merchant 使用
    8: optional i32 level,         // admin 使用
}

struct RegisterResponse {
    1: i64 user_id,
}

// Login
struct LoginRequest {
    1: string email; // 用户名
    2: string password; // 密码
}

struct LoginResponse {
    1: string token; // 分发 token
}

struct GetUserInfoByIDRequest {
    1: i64 user_id; // 使用 user_id 查询
}

struct GetUserInfoByIDResponse {
    1: User user;
}

struct UpdateUserRequest {
    1: User user;
}

struct UpdateUserResponse {
    1: User user;
}

struct DeleteUserRequest {
    1: i64 user_id;
}

struct DeleteUserResponse {
    1: bool success;
}

// ========== Customer APIs ==========
struct GetCustomerInfoByIDRequest {
    1: i64 user_id;  // 使用 user_id 查询
}

struct GetCustomerInfoByIDResponse {
    1: User user;
    2: Customer customer;
}

struct UpdateCustomerInfoByIDRequest {
    1: i64 user_id;
    2: optional string address;
    3: optional string phone;
}

struct UpdateCustomerInfoByIDResponse {
    1: Customer customer;
}

// ========== Merchant APIs ==========
struct GetMerchantInfoByIDRequest {
    1: i64 user_id;
}

struct GetMerchantInfoByIDResponse {
    1: User user;
    2: Merchant merchant;
}

struct UpdateMerchantInfoByIDRequest {
    1: i64 user_id;
    2: optional string address;
    3: optional string shop_name;
}

struct UpdateMerchantInfoByIDResponse {
    1: Merchant merchant;
}

// ========== Admin APIs ==========

struct GetAdminInfoByIDRequest {
    1: i64 user_id;
}

struct GetAdminInfoByIDResponse {
    1: User user;
    2: Admin admin;
}

struct UpdateAdminInfoByIDRequest {
    1: i64 user_id;
    2: optional i32 level;
}

struct UpdateAdminInfoByIDResponse {
    1: i32 level;
}

struct ListUsersRequest {
    1: required i32 page_number;
    2: required i32 page_size;
}

struct ListUsersResponse {
    1: list<User> users;
    2: i32 total_count;
}


service UserService {
    RegisterResponse Register(1: RegisterRequest req) throws (1: base.ErrorResponse err);
    LoginResponse Login(1: LoginRequest req) throws (1: base.ErrorResponse err);

    GetUserInfoByIDResponse GetUserInfoByID(1: GetUserInfoByIDRequest req) throws (1: base.ErrorResponse err);
    UpdateUserResponse UpdateUser(1: UpdateUserRequest req) throws (1: base.ErrorResponse err);
    DeleteUserResponse DeleteUser(1: DeleteUserRequest req) throws (1: base.ErrorResponse err);

    GetCustomerInfoByIDResponse GetCustomerInfoByID(1: GetCustomerInfoByIDRequest req) throws (1: base.ErrorResponse err);
    UpdateCustomerInfoByIDResponse UpdateCustomerInfoByID(1: UpdateCustomerInfoByIDRequest req) throws (1: base.ErrorResponse err);

    GetMerchantInfoByIDResponse GetMerchantInfoByID(1: GetMerchantInfoByIDRequest req) throws (1: base.ErrorResponse err);
    UpdateMerchantInfoByIDResponse UpdateMerchantInfoByID(1: UpdateMerchantInfoByIDRequest req) throws (1: base.ErrorResponse err);

    GetAdminInfoByIDResponse GetAdminInfoByID(1: GetAdminInfoByIDRequest req) throws (1: base.ErrorResponse err);
    UpdateAdminInfoByIDResponse UpdateAdminInfoByID(1: UpdateAdminInfoByIDRequest req) throws (1: base.ErrorResponse err);

    ListUsersResponse ListUsers(1: ListUsersRequest req) throws (1: base.ErrorResponse err);
}