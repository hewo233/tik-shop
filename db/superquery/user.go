package superquery

import (
	"errors"
	"github.com/hewo/tik-shop/db/model"
	"github.com/hewo/tik-shop/db/query"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/base"
	"github.com/hewo/tik-shop/kitex_gen/hewo/tikshop/user"
	"github.com/hewo/tik-shop/rpc/user-service/pkg/hash"
	"github.com/hewo/tik-shop/shared/consts"
	"github.com/hewo/tik-shop/shared/errno"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
)

var u = &query.Q.User

type LoginSqlManageImpl struct{}

func NewLoginSqlManageImpl() *LoginSqlManageImpl {
	return &LoginSqlManageImpl{}
}

func (m *LoginSqlManageImpl) Register(usr *model.User) (usrID int64, err error) {

	_, err = u.Where(u.Email.Eq(usr.Email)).First()
	if err == nil {
		return -1, &base.ErrorResponse{Code: errno.StatusBadRequestCode, Message: "email already exists"}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	// Debug log
	log.Println("superQuery usr: ", usr)

	err = u.Create(usr)
	if err != nil {
		return -1, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	return usr.ID, nil
}

func (m *LoginSqlManageImpl) Login(email, password string) (authed bool, id int64, role string, err error) {

	usr, err := u.Where(u.Email.Eq(email)).First()
	if err != nil {
		return false, -1, "", &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	hashed := usr.HashedPassword
	checked := hash.CheckPassword(hashed, password)

	if !checked {
		return false, -1, "", &base.ErrorResponse{Code: errno.StatusUnauthorizedCode, Message: "Incorrect Password"}
	}

	if usr.Status != consts.UserStatusActive {
		if usr.Status == consts.UserStatusBanned {
			return false, -1, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "User is banned"}
		}
		if usr.Status == consts.UserStatusDeleted {
			return false, -1, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "User is deleted"}
		}
	}

	return true, usr.ID, usr.Role, nil
}

func (m *LoginSqlManageImpl) GetUserInfoByID(id int64) (usrRet *user.User, err error) {
	usr, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	//log.Println("GetUSerInfoByID superQuery usr: ", usr)
	usrRet = &user.User{}

	err = copier.Copy(&usrRet, &usr)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	return usrRet, nil
}

func (m *LoginSqlManageImpl) UpdateUser(usr *user.User) (usrRet *user.User, err error) {
	exist, err := u.Where(u.ID.Eq(usr.Id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	// 这个接口暂时只能改这个
	//if usr.Email != "" {
	//	exist.Email = usr.Email
	//}
	if usr.Username != "" {
		exist.Username = usr.Username
	}

	_, err = u.Where(u.ID.Eq(usr.Id)).Updates(exist)
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	newUser, err := u.Where(u.ID.Eq(usr.Id)).First()
	if err != nil {
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(&usrRet, &newUser)

	return usrRet, nil
}

func (m *LoginSqlManageImpl) DeleteUserByID(id int64) error {
	_, err := u.Where(u.ID.Eq(id)).Delete()
	if err != nil {
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return nil
}

func (m *LoginSqlManageImpl) GetCustomerInfoByID(id int64) (usrRet *user.User, cusRet *user.Customer, err error) {
	usr, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(&usrRet, &usr)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	cusRet = &user.Customer{}
	err = copier.Copy(&cusRet, &usr.Customer)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return usrRet, cusRet, nil
}

func (m *LoginSqlManageImpl) GetMerchantInfoByID(id int64) (usrRet *user.User, merRet *user.Merchant, err error) {
	usr, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(&usrRet, &usr)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	merRet = &user.Merchant{}
	err = copier.Copy(&merRet, &usr.Merchant)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return usrRet, merRet, nil
}

func (m *LoginSqlManageImpl) GetAdminInfoByID(id int64) (usrRet *user.User, admRet *user.Admin, err error) {
	usr, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(&usrRet, &usr)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	admRet = &user.Admin{}
	err = copier.Copy(&admRet, &usr.Admin)
	if err != nil {
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	return usrRet, admRet, nil
}
