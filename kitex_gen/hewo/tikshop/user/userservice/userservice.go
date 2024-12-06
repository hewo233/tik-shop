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
	"Auth": kitex.NewMethodInfo(
		authHandler,
		newUserServiceAuthArgs,
		newUserServiceAuthResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AdminAuth": kitex.NewMethodInfo(
		adminAuthHandler,
		newUserServiceAdminAuthArgs,
		newUserServiceAdminAuthResult,
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
	"UpdatePassword": kitex.NewMethodInfo(
		updatePasswordHandler,
		newUserServiceUpdatePasswordArgs,
		newUserServiceUpdatePasswordResult,
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

func authHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAuthArgs)
	realResult := result.(*user.UserServiceAuthResult)
	success, err := handler.(user.UserService).Auth(ctx, realArg.Request)
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
func newUserServiceAuthArgs() interface{} {
	return user.NewUserServiceAuthArgs()
}

func newUserServiceAuthResult() interface{} {
	return user.NewUserServiceAuthResult()
}

func adminAuthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAdminAuthArgs)
	realResult := result.(*user.UserServiceAdminAuthResult)
	success, err := handler.(user.UserService).AdminAuth(ctx, realArg.Request)
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
func newUserServiceAdminAuthArgs() interface{} {
	return user.NewUserServiceAdminAuthArgs()
}

func newUserServiceAdminAuthResult() interface{} {
	return user.NewUserServiceAdminAuthResult()
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

func updatePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUpdatePasswordArgs)
	realResult := result.(*user.UserServiceUpdatePasswordResult)
	success, err := handler.(user.UserService).UpdatePassword(ctx, realArg.Request)
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
func newUserServiceUpdatePasswordArgs() interface{} {
	return user.NewUserServiceUpdatePasswordArgs()
}

func newUserServiceUpdatePasswordResult() interface{} {
	return user.NewUserServiceUpdatePasswordResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Auth(ctx context.Context, request *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceAuthArgs
	_args.Request = request
	var _result user.UserServiceAuthResult
	if err = p.c.Call(ctx, "Auth", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AdminAuth(ctx context.Context, request *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceAdminAuthArgs
	_args.Request = request
	var _result user.UserServiceAdminAuthResult
	if err = p.c.Call(ctx, "AdminAuth", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, request *user.GetUserInfoRequest) (r *user.GetUserInfoResponse, err error) {
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

func (p *kClient) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest) (r *user.UpdatePasswordResponse, err error) {
	var _args user.UserServiceUpdatePasswordArgs
	_args.Request = request
	var _result user.UserServiceUpdatePasswordResult
	if err = p.c.Call(ctx, "UpdatePassword", &_args, &_result); err != nil {
		return
	}
	switch {
	case _result.Err != nil:
		return r, _result.Err
	}
	return _result.GetSuccess(), nil
}
