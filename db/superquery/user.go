package superquery

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/db/superquery/utils"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
	"github.com/jinzhu/copier"
)

var u = query.Q.Users

func Auth(username, password string) (err error) {
	usr, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusNotFound, Message: err.Error()}
	}
	if usr.Role == "admin" {
		return &base.ErrorResponse{Code: consts.StatusBadGateway, Message: "Can't login as admin"}
	}
	hash := usr.HashedPassword
	checked := utils.CheckPassword(hash, password)
	if !checked {
		return &base.ErrorResponse{Code: consts.StatusUnauthorized, Message: "Incorrect Password"}
	}
	return nil
}

func AdminAuth(username, password string) error {
	usr, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusNotFound, Message: err.Error()}
	}
	if usr.Role != "admin" {
		return &base.ErrorResponse{Code: consts.StatusBadGateway, Message: "Can't login as normal user"}
	}
	hash := usr.HashedPassword
	checked := utils.CheckPassword(hash, password)
	if !checked {
		return &base.ErrorResponse{Code: consts.StatusUnauthorized, Message: "Incorrect Password"}
	}
	return nil

}

func GetUserInfo(id int64) (usrRet *user.User, err error) {
	usr, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusNotFound, Message: err.Error()}
	}
	err = copier.Copy(&usrRet, usr)
	if err != nil {
		return nil, &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return
}

func Register(username, email, password string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	usr := &model.Users{}
	usr.Username = username
	usr.Email = email
	usr.HashedPassword = *hash
	err = u.Create(usr)
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return nil
}

func UpdateUser(usr *model.Users) error {
	err := u.Save(usr)
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return nil
}

func UpdatePassword(id int64, oldPassword, newPassword string) error {
	us, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	ok := utils.CheckPassword(us.HashedPassword, oldPassword)
	if !ok {
		return &base.ErrorResponse{Code: consts.StatusNotAcceptable, Message: "old pass not match"}
	}

	hashnew, err := utils.HashPassword(oldPassword)
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}

	usr := &model.Users{}
	usr.Id = id
	usr.HashedPassword = *hashnew
	err = u.Save(usr)
	if err != nil {
		return &base.ErrorResponse{Code: consts.StatusInternalServerError, Message: err.Error()}
	}
	return nil
}
