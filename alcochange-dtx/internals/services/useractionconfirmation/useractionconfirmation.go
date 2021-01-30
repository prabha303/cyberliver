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
		userAction.WarningLabelRedeemed = req.WarningLabelRedeemed
		userAction.TermsAndConditionsRedeemed = req.TermsAndConditionsRedeemed
		userAction.AccessCodeVerified = req.AccessCodeVerified
		if req.WarningLabelRedeemed {
			errW := u.userActionDao.UpdateWarningLabelRedeemed(*userAction)
			if errW != nil {
				u.l.Error("UpdateWarningLabelRedeemed Error - ", errW)
				return nil, errW
			}
		}
		if req.TermsAndConditionsRedeemed {
			errT := u.userActionDao.UpdateTermsAndConditionsRedeemed(*userAction)
			if errT != nil {
				u.l.Error("UpdateTermsAndConditionsRedeemed Error - ", errT)
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
		userAction.AccessCodeVerified = req.AccessCodeVerified
		userAction.WarningLabelRedeemed = req.WarningLabelRedeemed
		userAction.TermsAndConditionsRedeemed = req.TermsAndConditionsRedeemed
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
