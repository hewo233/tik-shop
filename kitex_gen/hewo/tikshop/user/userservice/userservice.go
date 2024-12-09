// Code generated by Kitex v0.11.3. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	user "github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AdminLogin": kitex.NewMethodInfo(
		adminLoginHandler,
		newUserServiceAdminLoginArgs,
		newUserServiceAdminLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newUserServiceGetUserInfoArgs,
		newUserServiceGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateUser": kitex.NewMethodInfo(
		updateUserHandler,
		newUserServiceUpdateUserArgs,
		newUserServiceUpdateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdatePasswordByID": kitex.NewMethodInfo(
		updatePasswordByIDHandler,
		newUserServiceUpdatePasswordByIDArgs,
		newUserServiceUpdatePasswordByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func adminLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAdminLoginArgs)
	realResult := result.(*user.UserServiceAdminLoginResult)
	success, err := handler.(user.UserService).AdminLogin(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceAdminLoginArgs() interface{} {
	return user.NewUserServiceAdminLoginArgs()
}

func newUserServiceAdminLoginResult() interface{} {
	return user.NewUserServiceAdminLoginResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetUserInfoArgs)
	realResult := result.(*user.UserServiceGetUserInfoResult)
	success, err := handler.(user.UserService).GetUserInfo(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceGetUserInfoArgs() interface{} {
	return user.NewUserServiceGetUserInfoArgs()
}

func newUserServiceGetUserInfoResult() interface{} {
	return user.NewUserServiceGetUserInfoResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdateUserArgs)
	realResult := result.(*user.UserServiceUpdateUserResult)
	success, err := handler.(user.UserService).UpdateUser(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceUpdateUserArgs() interface{} {
	return user.NewUserServiceUpdateUserArgs()
}

func newUserServiceUpdateUserResult() interface{} {
	return user.NewUserServiceUpdateUserResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func updatePasswordByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdatePasswordByIDArgs)
	realResult := result.(*user.UserServiceUpdatePasswordByIDResult)
	success, err := handler.(user.UserService).UpdatePasswordByID(ctx, realArg.Request)
	if err != nil {
		switch v := err.(type) {
		case *base.ErrorResponse:
			realResult.Err = v
		default:
			return err
		}
	} else {
		realResult.Success = success
	}
	return nil
}
func newUserServiceUpdatePasswordByIDArgs() interface{} {
	return user.NewUserServiceUpdatePasswordByIDArgs()
}

func newUserServiceUpdatePasswordByIDResult() interface{} {
	return user.NewUserServiceUpdatePasswordByIDResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, request *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Request = request
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AdminLogin(ctx context.Context, request *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceAdminLoginArgs
	_args.Request = request
	var _result user.UserServiceAdminLoginResult
	if err = p.c.Call(ctx, "AdminLogin", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, request *user.GetUserInfoByIDRequest) (r *user.GetUserInfoByIDResponse, err error) {
	var _args user.UserServiceGetUserInfoArgs
	_args.Request = request
	var _result user.UserServiceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (r *user.UpdateUserResponse, err error) {
	var _args user.UserServiceUpdateUserArgs
	_args.Request = request
	var _result user.UserServiceUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, request *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Request = request
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePasswordByID(ctx context.Context, request *user.UpdatePasswordByIDRequest) (r *user.UpdatePasswordByIDResponse, err error) {
	var _args user.UserServiceUpdatePasswordByIDArgs
	_args.Request = request
	var _result user.UserServiceUpdatePasswordByIDResult
	if err = p.c.Call(ctx, "UpdatePasswordByID", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}
