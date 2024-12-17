// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "github.com/hewo/tik-shop/route/biz/handler/hewo/tikshop/route/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_auth := _api.Group("/auth", _authMw()...)
			_auth.POST("/login", append(_loginMw(), user.Login)...)
			_auth.POST("/register", append(_registerMw(), user.Register)...)
			_auth.GET("/verify", append(_verifyMw(), user.Verify)...)
			{
				_admin := _auth.Group("/admin", _adminMw()...)
				_admin.POST("/login", append(_adminloginMw(), user.AdminLogin)...)
			}
		}
		{
			_user := _api.Group("/user", _userMw()...)
			_user.GET("/:id", append(_getuserinfobyidMw(), user.GetUserInfoByID)...)
			_id := _user.Group("/:id", _idMw()...)
			_id.PUT("/password", append(_updatepasswordMw(), user.UpdatePassword)...)
			_user.PUT("/:id", append(_updateuserMw(), user.UpdateUser)...)
		}
	}
}
