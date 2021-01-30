package daos

import (
	"ecargoware/alcochange-dtx/models"
	"encoding/json"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SignUp struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewSignUp(l *log.Logger, dbConn *pg.DB) *SignUp {
	return &SignUp{
		l:      l,
		dbConn: dbConn,
	}
}

type SignUpDao interface {
	EmailIDExists(email string) bool
	GetRoleForSignup(signupCode string) models.UserType
	GetProductAccessFor(productAccess string) models.ProductAccess
	CreateUser(user models.User) (*models.User, error)
	CreateUserAccess(user models.UserAccess) (*models.UserAccess, error)
	CreateLoginDeviceDetails(user models.LoginDeviceDetails) (*models.LoginDeviceDetails, error)
	DeleteUser(ID int64) bool
	DeleteUserAccess(ID int64) bool
	GetUserActionConfirmationUUID(uudi string) (*models.UserActionConfirmation, error)
	UpdateUserActionConfirmation(userAction models.UserActionConfirmation) bool
	UserValidation(email string, password string) (*models.UserAccess, error)
	UpdateUserAccess(userAccess models.UserAccess) bool
	GetLoginDeviceDetails(userID int64) (*models.LoginDeviceDetails, error)
	UpdateLoginDeviceDetails(user models.LoginDeviceDetails) bool
	CreateLoginLog(loginLogs models.LoginLogs) bool
}

func (sp *SignUp) EmailIDExists(email string) bool {
	user := []models.User{}
	err := sp.dbConn.Model(&user).Where("LOWER(email_id) = LOWER(?)", email).Select()
	if err != nil {
		sp.l.Error("EmailIDExists Error", err.Error())
	}
	return len(user) > 0
}

func (sp *SignUp) GetRoleForSignup(signupCode string) models.UserType {
	role := models.UserType{}
	err := sp.dbConn.Model(&role).Where("LOWER(code) = LOWER(?)", signupCode).Select()
	if err != nil {
		sp.l.Error("GetRoleForSignup Error", err.Error())
	}
	return role
}

func (sp *SignUp) GetProductAccessFor(productAccess string) models.ProductAccess {
	productAcc := models.ProductAccess{}
	err := sp.dbConn.Model(&productAcc).Where("LOWER(code) = LOWER(?)", productAccess).Select()
	if err != nil {
		sp.l.Error("GetProductAccessFor Error", err.Error())
	}
	return productAcc
}

func (sp *SignUp) CreateUser(user models.User) (*models.User, error) {
	_, insErr := sp.dbConn.Model(&user).Insert()
	if insErr != nil {
		sp.l.Error("CreateUser Error--", insErr)
		return nil, insErr
	}
	dataBytes, _ := json.Marshal(user)
	sp.l.Debug("User table json : %q", string(dataBytes))
	return &user, nil
}

func (sp *SignUp) CreateUserAccess(userAccess models.UserAccess) (*models.UserAccess, error) {
	_, insErr := sp.dbConn.Model(&userAccess).Insert()
	if insErr != nil {
		sp.l.Error("CreateUserAccess Error--", insErr)
		return nil, insErr
	}
	dataBytes, _ := json.Marshal(userAccess)
	sp.l.Debug("CreateUserAccess table json : %q", string(dataBytes))
	return &userAccess, nil
}

func (sp *SignUp) DeleteUser(ID int64) bool {
	user := models.User{}
	_, insErr := sp.dbConn.Model(&user).Where("id = ? ", ID).Delete()
	if insErr != nil {
		sp.l.Error("DeleteUser Error", insErr.Error())
	}
	return true
}

func (sp *SignUp) DeleteUserAccess(ID int64) bool {
	user := models.UserAccess{}
	_, insErr := sp.dbConn.Model(&user).Where("id = ? ", ID).Delete()
	if insErr != nil {
		sp.l.Error("DeleteUserAccess Error", insErr.Error())
	}
	return true
}

func (sp *SignUp) CreateLoginDeviceDetails(loginDeviceDetails models.LoginDeviceDetails) (*models.LoginDeviceDetails, error) {
	_, insErr := sp.dbConn.Model(&loginDeviceDetails).Insert()
	if insErr != nil {
		sp.l.Error("CreateLoginDeviceDetails Error--", insErr)
		return nil, insErr
	}
	dataBytes, _ := json.Marshal(loginDeviceDetails)
	sp.l.Debug("CreateLoginDeviceDetails table json : %q", string(dataBytes))
	return &loginDeviceDetails, nil
}

func (sp *SignUp) GetUserActionConfirmationUUID(uuid string) (*models.UserActionConfirmation, error) {
	userAction := models.UserActionConfirmation{}
	err := sp.dbConn.Model(&userAction).Where("LOWER(device_uuid) = LOWER(?)", uuid).Select()
	if err != nil {
		sp.l.Error("GetUserActionConfirmationUUID Error", err.Error())
	}
	return &userAction, nil
}

func (sp *SignUp) UpdateUserActionConfirmation(userActionCon models.UserActionConfirmation) bool {
	_, uptErr := sp.dbConn.Model(&userActionCon).Column("user_id", "is_signed_up", "is_active", "version", "updated_at").Where("id=?", userActionCon.ID).Update()
	if uptErr != nil {
		sp.l.Error("UpdateUserActionConfirmation Error--", uptErr)
		return false
	}
	sp.l.Debug("UpdateUserActionConfirmation Affected", "Success")
	return true
}

func (sp *SignUp) UserValidation(email string, password string) (*models.UserAccess, error) {
	user := models.UserAccess{}
	err := sp.dbConn.Model(&user).Where("LOWER(email_id)=LOWER(?) AND password=? ", email, password).Select()
	if err != nil {
		sp.l.Error("UserValidation Error", email, password, err.Error())
		return nil, err
	}
	return &user, nil
}

func (sp *SignUp) UpdateUserAccess(userAccess models.UserAccess) bool {
	_, uptErr := sp.dbConn.Model(&userAccess).Column("alco_change_dtx_access", "is_active", "last_login", "version", "updated_at").Where("id=?", userAccess.ID).Update()
	if uptErr != nil {
		sp.l.Error("UpdateUserAccess Login Error--", uptErr)
		return false
	}
	sp.l.Debug("UpdateUserAccess Login Affected", "Success")
	return true
}

func (sp *SignUp) GetLoginDeviceDetails(userID int64) (*models.LoginDeviceDetails, error) {
	loginDeviceDetails := models.LoginDeviceDetails{}
	err := sp.dbConn.Model(&loginDeviceDetails).Where("user_id = ?", userID).Select()
	if err != nil {
		sp.l.Error("GetLoginDeviceDetails Error", err.Error())
		return nil, err
	}
	return &loginDeviceDetails, nil
}

func (sp *SignUp) UpdateLoginDeviceDetails(loginDeviceDetails models.LoginDeviceDetails) bool {
	_, uptErr := sp.dbConn.Model(&loginDeviceDetails).Column("timezone", "latitude", "longitude", "app_id", "last_login", "os_version",
		"user_app_version", "os_type", "device_uuid", "device_info", "network_info", "is_active", "version", "updated_at").Where("id=?", loginDeviceDetails.ID).Update()
	if uptErr != nil {
		sp.l.Error("UpdateLoginDeviceDetails Login Error--", uptErr)
		return false
	}
	sp.l.Debug("UpdateLoginDeviceDetails Login Affected", "Success")
	return true
}

func (sp *SignUp) CreateLoginLog(loginLogs models.LoginLogs) bool {
	_, insErr := sp.dbConn.Model(&loginLogs).Insert()
	if insErr != nil {
		sp.l.Error("CreateLoginLog Error--", insErr)
		return false
	}
	dataBytes, _ := json.Marshal(loginLogs)
	sp.l.Debug("CreateLoginLog table json : %q", string(dataBytes))
	return true
}
