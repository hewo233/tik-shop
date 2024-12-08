package superquery

import (
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/db/superquery/utils"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/hewo/tik-shop/shared/errno"
	"github.com/jinzhu/copier"
)

var u = query.Q.Users

type LoginSqlManageImpl struct{}

func NewLoginSqlManageImpl() *LoginSqlManageImpl {
	return &LoginSqlManageImpl{}
}

func (m *LoginSqlManageImpl) Login(username, password string) (authed bool, id string, err error) {
	usr, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return false, "", &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	if usr.Role == consts.Admin {
		return false, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "Can't login as admin"}
	}

	hash := usr.HashedPassword
	checked := utils.CheckPassword(hash, password)

	if !checked {
		return false, "", &base.ErrorResponse{Code: errno.StatusUnauthorizedCode, Message: "Incorrect Password"}
	}

	return true, usr.Username, nil
}

func AdminAuth(username, password string) (authed bool, id string, err error) {
	usr, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		return false, "", &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	if usr.Role != consts.Admin {
		return false, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "Can't login as admin"}
	}

	hash := usr.HashedPassword
	checked := utils.CheckPassword(hash, password)

	if !checked {
		return false, "", &base.ErrorResponse{Code: errno.StatusUnauthorizedCode, Message: "Incorrect Password"}
	}

	return true, usr.Username, nil

}

func GetUserInfoByID(id int64) (usrRet *user.User, err error) {
	usr, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}
	err = copier.Copy(&usrRet, usr)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return
}

func Register(username, email, password, role string) error {
	tmpUsr, err := u.Where(u.Username.Eq(username)).First()
	if tmpUsr != nil {
		// 不重名
		return &base.ErrorResponse{Code: errno.StatusConflictCode, Message: "Username already exists"}
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	usr := &model.Users{
		Username:       username,
		Email:          email,
		HashedPassword: hash,
		Role:           role,
	}

	err = u.Create(usr)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	return nil
}

func UpdateUser(usr *model.Users) error {
	err := u.Save(usr)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return nil
}

func UpdatePassword(id int64, oldPassword, newPassword string) error {
	us, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	ok := utils.CheckPassword(us.HashedPassword, oldPassword)
	if !ok {
		return &base.ErrorResponse{Code: errno.StatusNotAcceptableCode, Message: "old pass not match"}
	}

	hashnew, err := utils.HashPassword(oldPassword)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	usr := &model.Users{}
	usr.Id = id
	usr.HashedPassword = hashnew
	err = u.Save(usr)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return nil
}
