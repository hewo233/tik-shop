package main

import (
	"context"
	"time"

	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Mock data for users
var mockUsers = map[int64]*user.User{
	1: {Id: 1, Username: "john_doe", Email: "john.doe@example.com", Role: "user", CreatedAt: time.Now().Format(time.RFC3339)},
	2: {Id: 2, Username: "admin_user", Email: "admin@example.com", Role: "admin", CreatedAt: time.Now().Format(time.RFC3339)},
}

// Auth implements the UserServiceImpl interface.
func (s *UserServiceImpl) Auth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	// TODO: Your code here...
	if request.Username == "john_doe" && request.Password == "password123" {
		resp = &user.AuthResponse{Authorized: true}
	} else {
		err = &base.ErrorResponse{Code: 401, Message: "Unauthorized: Invalid credentials"}
	}
	return resp, err
}

// AdminAuth implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminAuth(ctx context.Context, request *user.AuthRequest) (resp *user.AuthResponse, err error) {
	// TODO: Your code here...
	if request.Username == "admin_user" && request.Password == "adminpass" {
		resp = &user.AuthResponse{Authorized: true}
	} else {
		err = &base.ErrorResponse{Code: 401, Message: "Unauthorized: Invalid admin credentials"}
	}
	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (resp *user.GetUserInfoResponse, err error) {
	// TODO: Your code here...
	usr, ok := mockUsers[request.Id]
	if !ok {
		err = &base.ErrorResponse{Code: 404, Message: "User not found"}
		return nil, err
	}
	resp = &user.GetUserInfoResponse{User: usr}
	return resp, nil
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	// TODO: Your code here...
	usr, ok := mockUsers[request.Id]
	if !ok {
		err = &base.ErrorResponse{Code: 404, Message: "User not found"}
		return nil, err
	}
	usr.Username = request.Username
	usr.Email = request.Email
	resp = &user.UpdateUserResponse{User: usr}
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	for _, u := range mockUsers {
		if u.Username == request.Username || u.Email == request.Email {
			err = &base.ErrorResponse{Code: 409, Message: "Username or email already exists"}
			return nil, err
		}
	}
	newID := int64(len(mockUsers) + 1)
	mockUsers[newID] = &user.User{
		Id:        newID,
		Username:  request.Username,
		Email:     request.Email,
		Role:      "user",
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	resp = &user.RegisterResponse{Message: "User registered successfully"}
	return resp, nil
}

// UpdatePassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (resp *user.UpdatePasswordResponse, err error) {
	// TODO: Your code here...
	_, ok := mockUsers[request.Id]
	if !ok {
		err = &base.ErrorResponse{Code: 404, Message: "User not found"}
		return nil, err
	}
	if request.OldPassword != "passwd" { // Simulating a password check
		err = &base.ErrorResponse{Code: 403, Message: "Incorrect old password"}
		return nil, err
	}
	resp = &user.UpdatePasswordResponse{Message: "Password updated successfully"}
	return resp, nil
}
