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
	log.Println("Register called with user:", usr)

	_, err = u.Where(u.Email.Eq(usr.Email)).First()
	if err == nil {
		log.Println("Email already exists:", usr.Email)
		return -1, &base.ErrorResponse{Code: errno.StatusConflictCode, Message: "Email already exists"}
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Error querying user by email:", err)
		return -1, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	log.Println("Creating new user:", usr)
	err = u.Create(usr)
	if err != nil {
		log.Println("Error creating user:", err)
		return -1, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	log.Println("User registered successfully with ID:", usr.ID)
	return usr.ID, nil
}

func (m *LoginSqlManageImpl) Login(email, password string) (authed bool, id int64, role string, err error) {
	log.Println("Login called with email:", email)

	usr, err := u.Where(u.Email.Eq(email)).First()
	if err != nil {
		log.Println("Error finding user by email:", err)
		return false, -1, "", &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	log.Println("User found:", usr.ID)
	hashed := usr.HashedPassword
	checked := hash.CheckPassword(hashed, password)

	if !checked {
		log.Println("Password check failed for user ID:", usr.ID)
		return false, -1, "", &base.ErrorResponse{Code: errno.StatusUnauthorizedCode, Message: "Incorrect Password"}
	}

	if usr.Status != consts.UserStatusActive {
		log.Println("User status is not active:", usr.Status)
		if usr.Status == consts.UserStatusBanned {
			return false, -1, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "User is banned"}
		}
		if usr.Status == consts.UserStatusDeleted {
			return false, -1, "", &base.ErrorResponse{Code: errno.ForbiddenCode, Message: "User is deleted"}
		}
	}

	log.Println("Login successful for user ID:", usr.ID)
	return true, usr.ID, usr.Role, nil
}

func (m *LoginSqlManageImpl) GetUserInfoByID(id int64) (usrRet *user.User, err error) {
	log.Println("GetUserInfoByID called with ID:", id)

	usr, err := u.Where(u.ID.Eq(id)).First()
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}

	err = copier.Copy(usrRet, usr)
	if err != nil {
		log.Println("Error copying user data:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	log.Println("GetUserInfoByID successful for user ID:", usr.ID)
	return usrRet, nil
}

func (m *LoginSqlManageImpl) UpdateUser(usr *user.User) (usrRet *user.User, err error) {
	log.Println("UpdateUser called with user ID:", usr.Id)

	exist, err := u.Where(u.ID.Eq(usr.Id)).First()
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	if usr.Username != "" {
		log.Println("Updating username for user ID:", usr.Id)
		exist.Username = usr.Username
	}

	_, err = u.Where(u.ID.Eq(usr.Id)).Updates(exist)
	if err != nil {
		log.Println("Error updating user:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	newUser, err := u.Where(u.ID.Eq(usr.Id)).First()
	if err != nil {
		log.Println("Error fetching updated user:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(usrRet, newUser)
	if err != nil {
		log.Println("Error copying updated user data:", err)
		return nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	log.Println("UpdateUser successful for user ID:", usr.Id)
	return usrRet, nil
}

func (m *LoginSqlManageImpl) DeleteUserByID(id int64) error {
	log.Println("DeleteUserByID called with ID:", id)

	_, err := u.Where(u.ID.Eq(id)).Delete()
	if err != nil {
		log.Println("Error deleting user by ID:", err)
		return &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	log.Println("User deleted successfully with ID:", id)
	return nil
}

func (m *LoginSqlManageImpl) GetCustomerInfoByID(id int64) (usrRet *user.User, cusRet *user.Customer, err error) {
	log.Println("GetCustomerInfoByID called with ID:", id)

	usr, err := u.Preload(u.Customer).Where(u.ID.Eq(id)).First()
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(usrRet, usr)
	if err != nil {
		log.Println("Error copying user data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	cusRet = &user.Customer{}
	err = copier.Copy(cusRet, usr.Customer)
	if err != nil {
		log.Println("Error copying customer data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	log.Println("GetCustomerInfoByID successful for user ID:", usr.ID)
	return usrRet, cusRet, nil
}

func (m *LoginSqlManageImpl) GetMerchantInfoByID(id int64) (usrRet *user.User, merRet *user.Merchant, err error) {
	log.Println("GetMerchantInfoByID called with ID:", id)

	usr, err := u.Preload(u.Merchant).Where(u.ID.Eq(id)).First()
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(usrRet, usr)
	if err != nil {
		log.Println("Error copying user data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	merRet = &user.Merchant{}
	err = copier.Copy(merRet, usr.Merchant)
	if err != nil {
		log.Println("Error copying merchant data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	log.Println("GetMerchantInfoByID successful for user ID:", usr.ID)
	return usrRet, merRet, nil
}

func (m *LoginSqlManageImpl) GetAdminInfoByID(id int64) (usrRet *user.User, admRet *user.Admin, err error) {
	log.Println("GetAdminInfoByID called with ID:", id)

	usr, err := u.Preload(u.Admin).Where(u.ID.Eq(id)).First()
	if err != nil {
		log.Println("Error finding user by ID:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusNotFoundCode, Message: err.Error()}
	}

	usrRet = &user.User{}
	err = copier.Copy(usrRet, usr)
	if err != nil {
		log.Println("Error copying user data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}

	admRet = &user.Admin{}
	err = copier.Copy(admRet, usr.Admin)
	if err != nil {
		log.Println("Error copying admin data:", err)
		return nil, nil, &base.ErrorResponse{Code: errno.StatusInternalServerErrorCode, Message: err.Error()}
	}
	log.Println("GetAdminInfoByID successful for user ID:", usr.ID)
	return usrRet, admRet, nil
}
