namespace go hewo.tikshop.user

include "base.thrift"
// 定义公共的用户信息结构
struct User {
    1: i64    id; // 用户 ID
    2: string username; // 用户名
    3: string email; // 邮箱
    4: string role;
    5: string createdAt; // 创建时间 (ISO8601 格式字符串)
}

// Get User
struct GetUserInfoByIDRequest {
    1: i64 id;
}

struct GetUserInfoByIDResponse {
    1: User user;
}

// Login
struct LoginRequest {
    1: string username; // 用户名
    2: string password; // 密码
}

struct LoginResponse {
    1: string token; // 分发 token
}

// 定义用户模块的请求和响应结构
struct UpdateUserRequest {
    1: User user;
}

struct UpdateUserResponse {
    1: User user;
}

struct RegisterRequest {
    1: string username; // 注册用户名
    2: string email; // 注册邮箱
    3: string password; // 注册密码
    4: string role;
}

struct RegisterResponse {
    1: User user;
}

struct UpdatePasswordByIDRequest {
    1: i64    id;
    2: string oldPassword; // 旧密码
    3: string newPassword; // 新密码
}

struct UpdatePasswordByIDResponse {
    1: bool changedFlag; // 密码修改成功标志
}

// 定义 UserService 接口
service UserService {
    // 鉴权模块
    LoginResponse Login(1: LoginRequest request) throws (1: base.ErrorResponse err);
    LoginResponse AdminLogin(1: LoginRequest request) throws (1: base.ErrorResponse err);

    // 用户模块
    GetUserInfoByIDResponse GetUserInfo(1: GetUserInfoByIDRequest request) throws (1: base.ErrorResponse err);
    UpdateUserResponse UpdateUser(1: UpdateUserRequest request) throws (1: base.ErrorResponse err);
    RegisterResponse Register(1: RegisterRequest request) throws (1: base.ErrorResponse err);
    UpdatePasswordByIDResponse UpdatePasswordByID(1: UpdatePasswordByIDRequest request) throws (1: base.ErrorResponse err);
}