package main

import (
	"context"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Auth implements the UserServiceImpl interface.
func (s *UserServiceImpl) Auth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	// TODO: Your code here...
	return
}

// AdminAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminAuth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	// TODO: Your code here...
	return
}
