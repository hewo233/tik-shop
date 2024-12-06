amespace go hewo.tikshop.route.user

include "./base.thrift"

struct UserResponse {
  1: i64 id
  2: string username
  3: string email
  4: string createdAt
}

struct GetUserRequest {
  1: i64 id (api.path = "id", api.vd = "$>0")
}

struct RegisterRequest {
  1: string username (api.body = "name", api.form = "name", api.vd = "(len($) > 0 && len($) < 128)")
  2: string email (api.body = "email", api.form = "email", api.vd = "(len($) > 0 && len($) < 512)")
  3: string password (api.body = "password", api.form = "password", api.vd = "(len($) > 0 && len($) < 128)")
}
struct UpdatePasswordRequest {
  1: i64 id (api.path = "id", api.vd = "$>0")
  2: string oldPassword (api.body = "password", api.form = "password", api.vd = "(len($) > 0 && len($) < 128)")
  3: string newPassword (api.body = "password", api.form = "password", api.vd = "(len($) > 0 && len($) < 128)")
}


struct UpdateUserRequest {
  1: i64 id (api.path = "id", api.vd = "$>0")
  2: string name (api.body = "name", api.form = "name", api.vd = "(len($) > 0 && len($) < 128)")
  3: string email (api.body = "email", api.form = "email", api.vd = "(len($) > 0 && len($) < 512)")
}

struct LoginRequest {
  1: string username (api.body = "name", api.form = "name", api.vd = "(len($) > 0 && len($) < 128)")
  2: string password (api.body = "password", api.form = "password", api.vd = "(len($) > 0 && len($) < 128)")
}

struct LoginResponse {
  1: string token
}

struct VerifyRequest {}

struct VerifyResponse {
  1: bool authorized
}

service UserService {
  UserResponse getUser(1: GetUserRequest req) (api.get="/api/user/:id")
  UserResponse updateUser(1: UpdateUserRequest req) (api.put="/api/user/:id")
  base.MessageResponse register(1: RegisterRequest req) (api.post="/api/user/register")
  base.MessageResponse updatePassword(1: UpdatePasswordRequest req) (api.put="/api/user/:id/password")

  // auth
  LoginResponse login(1: LoginRequest req)(api.post="/api/auth/login")
  LoginResponse adminLogin(1: LoginRequest req) (api.post="/api/auth/admin/login")
  VerifyResponse verify(1: VerifyRequest req) (api.get="/api/auth/verify")
}
