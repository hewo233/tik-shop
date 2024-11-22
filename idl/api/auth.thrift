namespace go hewo.tikshop.route.auth

include "./base.thrift"

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

service AuthService {
  LoginResponse login(1: LoginRequest req)(api.post="/api/auth/login")
  LoginResponse adminLogin(1: LoginRequest req) (api.post="/api/auth/admin/login")
  VerifyResponse verify(1: VerifyRequest req) (api.get="/api/auth/verify")
}
