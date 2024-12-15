package main

import (
	"context"
	"github.com/hertz-contrib/paseto"
	"github.com/hewo/tik-shop/shared/consts"
	"log"
	"time"

	"github.com/hewo/tik-shop/db/model"
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
	AdminLogin(username, password string) (authed bool, id string, err error)
	GetUserInfoByID(id int64) (usrRet *user.User, err error)
	UpdateUser(usr *model.Users) error
	Register(username, email, password, role string) (usrRet *user.User, err error)
	UpdatePasswordByID(id int64, oldPassword, newPassword string) error
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {

	resp = new(user.LoginResponse)
	authed, id, err := s.LoginSqlManage.Login(request.Username, request.Password)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if authed == true {

		nowTime := time.Now()
		resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
			ID:        id,
			Issuer:    consts.Issuer,
			Audience:  consts.User,
			IssuedAt:  nowTime,
			NotBefore: nowTime,
			ExpiredAt: nowTime.Add(consts.SevenDays),
		})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return resp, nil
}

// AdminLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) AdminLogin(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {

	resp = new(user.LoginResponse)
	authed, id, err := s.LoginSqlManage.AdminLogin(request.Username, request.Password)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if authed == true {

		nowTime := time.Now()
		resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
			ID:        id,
			Issuer:    consts.Issuer,
			Audience:  consts.Admin,
			IssuedAt:  nowTime,
			NotBefore: nowTime,
			ExpiredAt: nowTime.Add(consts.ThirtyDays),
		})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return resp, nil
}

// GetUserInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoByID(ctx context.Context, request *user.GetUserInfoByIDRequest) (resp *user.GetUserInfoByIDResponse, err error) {

	resp = new(user.GetUserInfoByIDResponse)

	usr, err := s.LoginSqlManage.GetUserInfoByID(request.Id)
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

	resp = new(user.UpdateUserResponse)

	u := &model.Users{}

	err = copier.Copy(&u, request)
	if err != nil {
		return nil, err
	}

	err = s.LoginSqlManage.UpdateUser(u)
	if err != nil {
		return nil, err
	}

	usr, err := s.LoginSqlManage.GetUserInfoByID(request.User.GetId())
	if err != nil {
		return nil, err
	}

	resp = &user.UpdateUserResponse{User: usr}

	return resp, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {

	resp = new(user.RegisterResponse)

	usr, err := s.LoginSqlManage.Register(request.Username, request.Email, request.Password, request.Role)
	if err != nil {
		return nil, err
	}

	resp = &user.RegisterResponse{User: usr}

	return resp, nil
}

// UpdatePasswordByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdatePasswordByID(ctx context.Context, request *user.UpdatePasswordByIDRequest) (resp *user.UpdatePasswordByIDResponse, err error) {
	err = s.LoginSqlManage.UpdatePasswordByID(request.Id, request.OldPassword, request.NewPassword_)
	if err != nil {
		return nil, err
	}
	// TODO
	return resp, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, request *user.GetUserInfoByIDRequest) (resp *user.GetUserInfoByIDResponse, err error) {
	// TODO: Your code here...
	return
}
