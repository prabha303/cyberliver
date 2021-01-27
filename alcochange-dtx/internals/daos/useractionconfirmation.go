package daos

import (
	"ecargoware/alcochange-dtx/models"
	"encoding/json"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type UserActionCon struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewUserAction(l *log.Logger, dbConn *pg.DB) *UserActionCon {
	return &UserActionCon{
		l:      l,
		dbConn: dbConn,
	}
}

type UserActionDao interface {
	GetUserActionByUUID(uuid string) *models.UserActionConfirmation
	GetUserActionByEmailID(uuid string) *models.UserActionConfirmation
	UpdateWarningLabelRead(userAction models.UserActionConfirmation) error
	UpdateTermsAndPrivacyRead(userAction models.UserActionConfirmation) error
	UpdateAccessCodeVerified(userAction models.UserActionConfirmation) error
	CreateUserActionConfirm(userAction models.UserActionConfirmation) (*models.UserActionConfirmation, error)
}

func (u *UserActionCon) GetUserActionByUUID(uuid string) *models.UserActionConfirmation {
	userAction := models.UserActionConfirmation{}
	u.dbConn.Model(&userAction).Where("device_uuid	 = ? ", uuid).Select()
	return &userAction
}

func (u *UserActionCon) GetUserActionByEmailID(emaiID string) *models.UserActionConfirmation {
	userAction := models.UserActionConfirmation{}
	u.dbConn.Model(&userAction).Where("email_id	 = ? ", emaiID).Select()
	return &userAction
}

func (u *UserActionCon) UpdateWarningLabelRead(userAction models.UserActionConfirmation) error {
	res, uptErr := u.dbConn.Model(&userAction).Column("warning_label_read", "version", "is_active", "updated_at").Where("id=?", userAction.ID).Update()
	if uptErr != nil {
		u.l.Errorf(" UpdateWarningLabelRead Error--", uptErr)
		return uptErr
	}
	u.l.Debug("UpdateWarningLabelRead - ", userAction.DeviceUUID, res.RowsAffected())
	return nil
}

func (u *UserActionCon) UpdateTermsAndPrivacyRead(userAction models.UserActionConfirmation) error {
	res, uptErr := u.dbConn.Model(&userAction).Column("terms_and_privacy_read", "version", "is_active", "updated_at").Where("id=?", userAction.ID).Update()
	if uptErr != nil {
		u.l.Errorf(" UpdateTermsAndPrivacyRead Error--", uptErr)
		return uptErr
	}
	u.l.Debug("UpdateTermsAndPrivacyRead - ", userAction.DeviceUUID, res.RowsAffected())
	return nil
}

func (u *UserActionCon) UpdateAccessCodeVerified(userAction models.UserActionConfirmation) error {
	res, uptErr := u.dbConn.Model(&userAction).Column("access_code_verified", "version", "is_active", "updated_at").Where("id=?", userAction.ID).Update()
	if uptErr != nil {
		u.l.Errorf(" UpdateAccessCodeVerified Error--", uptErr)
		return uptErr
	}
	u.l.Debug("UpdateAccessCodeVerified - ", userAction.DeviceUUID, res.RowsAffected())
	return nil
}

func (k *UserActionCon) CreateUserActionConfirm(userAction models.UserActionConfirmation) (*models.UserActionConfirmation, error) {
	_, insErr := k.dbConn.Model(&userAction).Insert()
	if insErr != nil {
		k.l.Error("CreateUserActionConfirm Error--", insErr)
		return nil, insErr
	}
	dataBytes, _ := json.Marshal(userAction)
	k.l.Debug("CreateUserActionConfirm json : %q", string(dataBytes))
	return &userAction, nil
}
