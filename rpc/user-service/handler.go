package main

import (
	"context"
	"github.com/hertz-contrib/paseto"
	"github.com/hewo/tik-shop/rpc/user-service/pkg/hash"
	"github.com/hewo/tik-shop/shared/consts"
	"log"
	"strconv"
	"time"

	"github.com/hewo/tik-shop/db/model"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	LoginSqlManage
	TokenGenerator
}

type LoginSqlManage interface {
	Login(username, password string) (authed bool, id int64, role string, err error)
	AdminLogin(username, password string) (authed bool, id int64, err error)
	GetUserInfoByID(id int64) (usrRet *user.User, err error)
	UpdateUser(usr *user.User) (usrRet *user.User, err error)
	Register(usr *model.User) (usrRet int64, err error)
	DeleteUserByID(id int64) error
	GetCustomerInfoByID(id int64) (usrRet *user.User, cusRet *user.Customer, err error)
	GetMerchantInfoByID(id int64) (usrRet *user.User, merRet *user.Merchant, err error)
	GetAdminInfoByID(id int64) (usrRet *user.User, admRet *user.Admin, err error)
}

type TokenGenerator interface {
	CreateToken(claims *paseto.StandardClaims) (token string, err error)
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.RegisterRequest) (resp *user.RegisterResponse, err error) {

	resp = new(user.RegisterResponse)

	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	usr := &model.User{
		Username:       request.Username,
		Email:          request.Email,
		HashedPassword: hashedPassword,
		Role:           request.Role,
		Status:         consts.UserStatusActive,
	}

	switch request.Role {
	case consts.RoleCustomer:
		usr.Customer = &model.Customer{
			Address: *request.Address,
			Phone:   *request.Phone,
		}
	case consts.RoleMerchant:
		usr.Merchant = &model.Merchant{
			Address:  *request.Address,
			ShopName: *request.ShopName,
		}
	case consts.RoleAdmin:
		usr.Admin = &model.Admin{
			Level: int(*request.Level),
		}
	}

	usrID, err := s.LoginSqlManage.Register(usr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp = &user.RegisterResponse{
		UserId: usrID,
	}

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.LoginRequest) (resp *user.LoginResponse, err error) {

	resp = new(user.LoginResponse)

	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	authed, id, role, err := s.LoginSqlManage.Login(request.Email, hashedPassword)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if authed == false {
		return nil, err
	}

	idStr := strconv.FormatInt(id, 10)

	nowTime := time.Now()
	resp.Token, err = s.TokenGenerator.CreateToken(&paseto.StandardClaims{
		ID:        idStr,
		Issuer:    consts.Issuer,
		Audience:  role,
		IssuedAt:  nowTime,
		NotBefore: nowTime,
		ExpiredAt: nowTime.Add(consts.SevenDays),
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil
}

// GetUserInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfoByID(ctx context.Context, request *user.GetUserInfoByIDRequest) (resp *user.GetUserInfoByIDResponse, err error) {

	resp = new(user.GetUserInfoByIDResponse)

	usr, err := s.LoginSqlManage.GetUserInfoByID(request.UserId)
	if err != nil {
		log.Println("GetUserInfoByID error: ", err)
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

	usr, err := s.LoginSqlManage.UpdateUser(request.User)

	if err != nil {
		log.Println("UpdateUser error: ", err)
		return nil, err
	}

	resp = &user.UpdateUserResponse{
		User: usr,
	}

	return resp, nil
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {

	resp = new(user.DeleteUserResponse)

	err = s.LoginSqlManage.DeleteUserByID(req.UserId)

	if err != nil {
		return nil, err
	}

	resp = &user.DeleteUserResponse{
		Success: true,
	}

	return resp, nil
}

// GetCustomerInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetCustomerInfoByID(ctx context.Context, req *user.GetCustomerInfoByIDRequest) (resp *user.GetCustomerInfoByIDResponse, err error) {
	resp = new(user.GetCustomerInfoByIDResponse)

	usr, cus, err := s.LoginSqlManage.GetCustomerInfoByID(req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &user.GetCustomerInfoByIDResponse{
		User:     usr,
		Customer: cus,
	}
	return resp, nil
}

// UpdateCustomerInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateCustomerInfoByID(ctx context.Context, req *user.UpdateCustomerInfoByIDRequest) (resp *user.UpdateCustomerInfoByIDResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMerchantInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMerchantInfoByID(ctx context.Context, req *user.GetMerchantInfoByIDRequest) (resp *user.GetMerchantInfoByIDResponse, err error) {
	resp = new(user.GetMerchantInfoByIDResponse)

	usr, mer, err := s.LoginSqlManage.GetMerchantInfoByID(req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &user.GetMerchantInfoByIDResponse{
		User:     usr,
		Merchant: mer,
	}
	return resp, nil
}

// UpdateMerchantInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateMerchantInfoByID(ctx context.Context, req *user.UpdateMerchantInfoByIDRequest) (resp *user.UpdateMerchantInfoByIDResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAdminInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetAdminInfoByID(ctx context.Context, req *user.GetAdminInfoByIDRequest) (resp *user.GetAdminInfoByIDResponse, err error) {
	resp = new(user.GetAdminInfoByIDResponse)

	usr, adm, err := s.LoginSqlManage.GetAdminInfoByID(req.UserId)
	if err != nil {
		return nil, err
	}
	resp = &user.GetAdminInfoByIDResponse{
		User:  usr,
		Admin: adm,
	}
	return resp, nil
}

// UpdateAdminInfoByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateAdminInfoByID(ctx context.Context, req *user.UpdateAdminInfoByIDRequest) (resp *user.UpdateAdminInfoByIDResponse, err error) {
	// TODO: Your code here...
	return
}

// ListUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) ListUsers(ctx context.Context, req *user.ListUsersRequest) (resp *user.ListUsersResponse, err error) {
	// TODO: Your code here...
	return
}
