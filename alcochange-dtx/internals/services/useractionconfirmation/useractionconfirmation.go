package useractionconfirmation

import (
	"errors"

	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/models"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type UserActionConfirm struct {
	dbConn        *pg.DB
	l             *log.Logger
	userActionDao daos.UserActionDao
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func NewUserActionConfirm(l *log.Logger, dbConn *pg.DB) *UserActionConfirm {
	return &UserActionConfirm{
		l:             l,
		dbConn:        dbConn,
		userActionDao: daos.NewUserAction(l, dbConn),
	}
}

func (u *UserActionConfirm) UserActionConfirmation(req dtos.UserActionConfirmationReq) (*dtos.UserActionConfirmationResponse, error) {
	userAction := &models.UserActionConfirmation{}

	if req.DeviceUUID == "" && req.EmailID == "" {
		return nil, errors.New("Invalid input")
	}

	if req.DeviceUUID != "" {
		userAction = u.userActionDao.GetUserActionByUUID(req.DeviceUUID)
	}
	if userAction.ID == 0 && req.EmailID != "" {
		userAction = u.userActionDao.GetUserActionByEmailID(req.EmailID)
	}

	if userAction.ID > 0 {
		userAction.BeforeUpdate(req.TimeZone)
		userAction.WarningLabelRead = req.WarningLabelRead
		userAction.TermsAndPrivacyRead = req.TermsAndPrivacyRead
		userAction.AccessCodeVerified = req.AccessCodeVerified
		if req.WarningLabelRead {
			errW := u.userActionDao.UpdateWarningLabelRead(*userAction)
			if errW != nil {
				u.l.Error("UpdateWarningLabelRead Error - ", errW)
				return nil, errW
			}
		}
		if req.TermsAndPrivacyRead {
			errT := u.userActionDao.UpdateTermsAndPrivacyRead(*userAction)
			if errT != nil {
				u.l.Error("UpdateTermsAndPrivacyRead Error - ", errT)
				return nil, errT
			}
		}
		if req.AccessCodeVerified {
			errA := u.userActionDao.UpdateAccessCodeVerified(*userAction)
			if errA != nil {
				u.l.Error("UpdateAccessCodeVerified Error - ", errA)
				return nil, errA
			}
		}
	} else {
		userAction.BeforeInsert(req.TimeZone)
		userAction.DeviceUUID = req.DeviceUUID
		userAction.EmailID = req.EmailID
		userAction.AccessCodeVerified = req.AccessCodeVerified
		userAction.WarningLabelRead = req.WarningLabelRead
		userAction.TermsAndPrivacyRead = req.TermsAndPrivacyRead
		_, errT := u.userActionDao.CreateUserActionConfirm(*userAction)
		if errT != nil {
			u.l.Error("CreateUserActionConfirm Error - ", errT)
			return nil, errT
		}
	}
	res := dtos.UserActionConfirmationResponse{}
	res.Message = "Saved succssfully"
	return &res, nil
}
