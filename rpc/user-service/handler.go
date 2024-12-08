package main

import (
	"context"
	"github.com/hertz-contrib/paseto"
	"github.com/hewo/tik-shop/shared/consts"
	"log"
	"time"

	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/superquery"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
	"github.com/jinzhu/copier"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	LoginSqlManage
	TokenGenerator
}

type LoginSqlManage interface {
	Login(username, password string) (authed bool, id string, err error)
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {
	authed, id, err := s.LoginSqlManage.Login(request.Username, request.Password)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if authed == true {

		nowTime := time.Now()
		resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
			ID:        id,
			Issuer:    "tik-shop",
			Audience:  "user",
			IssuedAt:  nowTime,
			NotBefore: nowTime,
			ExpiredAt: nowTime.Add(consts.SevenDays),
		})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	// TODO
	return resp, nil
}

// AdminLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminLogin(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {
	err = superquery.AdminAuth(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	// TODO
	return resp, err
}

// GetUserInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoByID(ctx context.Context, request *user.GetUserInfoByIDRequest) (resp *user.GetUserInfoByIDResponse, err error) {
	// TODO: Your code here...
	usr, err := superquery.GetUserInfoByID(request.Id)
	if err != nil {
		return nil, err
	}
	resp = &user.GetUserInfoByIDResponse{
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
	usr, err := superquery.GetUserInfoByID(request.Id)
	if err != nil {
		return nil, err
	}
	resp = &user.UpdateUserResponse{User: usr}
	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	err = superquery.Register(request.Username, request.Email, request.Password, request.Role)
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
	// TODO
	return resp, nil
}
