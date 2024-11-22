namespace go hewo.tikshop.route.user

include "base.thrift"

struct User {
  1: i64 id
  2: string username
  3: string email
  4: string createdAt
}

struct RegisterRequest {
  1: string username
  2: string email
  3: string password
}

struct UpdatePasswordRequest {
  1: string oldPassword
  2: string newPassword
}

struct UpdateUserRequest {
  1: string username
  2: string email
}


service UserService {
  User getUser(1: i64 id) throws (1: base.BaseErrorResponse error) (api.get 
  User updateUser(1: i64 id, 2: UpdateUserRequest req) throws (1: base.BaseErrorResponse error)
  void register(1: RegisterRequest req) throws (1: base.BaseErrorResponse error)
  void updatePassword(1: i64 id, 2: UpdatePasswordRequest req) throws (1: base.BaseErrorResponse error)
}
