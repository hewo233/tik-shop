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
	"log"
)

var u = &query.Q.Users

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

func (m *LoginSqlManageImpl) AdminLogin(username, password string) (authed bool, id string, err error) {
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

func (m *LoginSqlManageImpl) GetUserInfoByID(id int64) (usrRet *user.User, err error) {
	usr, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}
	err = copier.Copy(&usrRet, usr)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return usrRet, nil
}

func (m *LoginSqlManageImpl) Register(username, email, password, role string) (usrRet *user.User, err error) {

	log.Println("superQuery u: ", u)

	log.Println("superQuery Register: ", username, email, password, role)

	tmpUsr, err := u.Where(u.Username.Eq(username)).First()
	if err != nil {
		log.Println("superQuery Register temUsr: ", err)
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	if tmpUsr != nil {
		// 不重名
		return nil, &base.ErrorResponse{Code: errno.StatusConflictCode, Message: "Username already exists"}
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	usr := &model.Users{
		Username:       username,
		Email:          email,
		HashedPassword: hash,
		Role:           role,
	}

	log.Println("superQuery usr: ", usr)

	err = u.Create(usr)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	err = copier.Copy(&usrRet, usr)

	return usrRet, nil
}

func (m *LoginSqlManageImpl) UpdateUser(usr *model.Users) error {
	_, err := u.Where(u.Id.Eq(usr.Id)).Updates(usr)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	return nil
}

func (m *LoginSqlManageImpl) UpdatePasswordByID(id int64, oldPassword, newPassword string) error {
	temUsr, err := u.Where(u.Id.Eq(id)).First()
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	ok := utils.CheckPassword(temUsr.HashedPassword, oldPassword)
	if !ok {
		return &base.ErrorResponse{Code: errno.StatusNotAcceptableCode, Message: "old pass not match"}
	}

	hashNew, err := utils.HashPassword(newPassword)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	_, err = u.Where(u.Id.Eq(id)).Update(u.HashedPassword, hashNew)
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return nil
}
