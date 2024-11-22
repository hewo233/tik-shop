namespace go hewo.tikshop.route.auth

include "base.thrift"

struct LoginRequest {
  1: string username
  2: string password
}

struct LoginResponse {
  1: string token
}

struct VerifyResponse {
  1: bool authorized
}

service AuthService {
  LoginResponse login(1: LoginRequest req) throws (1: base.BaseErrorResponse error) (api.post="/api/auth/login")
  LoginResponse adminLogin(1: LoginRequest req) throws (1: base.BaseErrorResponse error) (api.post="/api/auth/admin/login")
  VerifyResponse verify() throws (1: base.BaseErrorResponse error) (api.get="/api/auth/verify")
}
