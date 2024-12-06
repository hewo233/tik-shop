package main

import (
	"context"

	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/superquery"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
	"github.com/jinzhu/copier"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Auth implements the UserServiceImpl interface.
func (s *UserServiceImpl) Auth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	err = superquery.Auth(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	resp = &user.AuthResponse{Authorized: true}
	return resp, err
}

// AdminAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminAuth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	err = superquery.AdminAuth(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	resp = &user.AuthResponse{Authorized: true}
	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	usr, err := superquery.GetUserInfo(request.Id)
	if err != nil {
		return nil, err
	}
	resp = &user.GetUserInfoResponse{
		User: usr,
	}
	return resp, nil
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	u := &model.Users{}
	err = copier.Copy(&u, request)
	if err != nil {
		return nil, err
	}
	err = superquery.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	usr, err := superquery.GetUserInfo(request.Id)
	if err != nil {
		return nil, err
	}
	resp = &user.UpdateUserResponse{User: usr}
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	err = superquery.Register(request.Username, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	resp = &user.RegisterResponse{Message: "User registered successfully"}
	return resp, nil
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	err = superquery.UpdatePassword(request.Id, request.OldPassword, request.NewPassword_)
	if err != nil {
		return nil, err
	}
	resp = &user.UpdatePasswordResponse{Message: "Password updated successfully"}
	return resp, nil
}
