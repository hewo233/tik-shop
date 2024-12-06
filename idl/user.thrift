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
struct GetUserInfoRequest {
    1: i64 id;
}

struct GetUserInfoResponse {
    1: User user;
}

// 定义鉴权模块的请求和响应结构
struct AuthRequest {
    1: string username; // 用户名
    2: string password; // 密码
}

struct AuthResponse {
    1: string token; // 是否通过验证
}

// 定义用户模块的请求和响应结构
struct UpdateUserRequest {
    1: i64    id;
    2: string username; // 新的用户名
    3: string email; // 新的邮箱
}

struct UpdateUserResponse {
    1: User user;
}

struct RegisterRequest {
    1: string username; // 注册用户名
    2: string email; // 注册邮箱
    3: string password; // 注册密码
}

struct RegisterResponse {
    1: string message; // 注册成功提示信息
}

struct UpdatePasswordRequest {
    1: i64    id;
    2: string oldPassword; // 旧密码
    3: string newPassword; // 新密码
}

struct UpdatePasswordResponse {
    1: bool changedFlag; // 密码修改成功标志
}

// 定义 UserService 接口
service UserService {
    // 鉴权模块
    AuthResponse Auth(1: AuthRequest request) throws (1: base.ErrorResponse err);
    AuthResponse AdminAuth(1: AuthRequest request) throws (1: base.ErrorResponse err);

    // 用户模块
    GetUserInfoResponse GetUserInfo(1: GetUserInfoRequest request) throws (1: base.ErrorResponse err);
    UpdateUserResponse UpdateUser(1: UpdateUserRequest request) throws (1: base.ErrorResponse err);
    RegisterResponse Register(1: RegisterRequest request) throws (1: base.ErrorResponse err);
    UpdatePasswordResponse UpdatePassword(1: UpdatePasswordRequest request) throws (1: base.ErrorResponse err);
}