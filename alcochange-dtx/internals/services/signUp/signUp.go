package signUp

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/models"
	"ecargoware/alcochange-dtx/utils"
	"errors"
	"strings"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

//DFMMMyyyyNoSep := "Jan2006"

type SignUp struct {
	dbConn    *pg.DB
	l         *log.Logger
	signUpDao daos.SignUpDao
}

func NewSignUp(l *log.Logger, dbConn *pg.DB) *SignUp {
	return &SignUp{
		l:         l,
		dbConn:    dbConn,
		signUpDao: daos.NewSignUp(l, dbConn),
	}
}

const layOut = "2006-01-02 15:04"

func (sp *SignUp) UserSignUp(signReq dtos.SignUpRequest) (*dtos.SignUpResponse, error) {

	emailExists := sp.signUpDao.EmailIDExists(signReq.RegisterUserRequest.EmailID)
	if emailExists {
		return nil, errors.New("EmailID Already Exists")
	}
	userType := sp.signUpDao.GetRoleForSignup(signReq.RegisterUserRequest.SignUpFor)
	if userType.ID == 0 {
		return nil, errors.New("Invalid Value signUpFor")
	}

	if signReq.RegisterUserRequest.Password == "" {
		return nil, errors.New("Password should not empty")
	}

	currentTime, _ := utils.CurrentTimeWithZone(signReq.Timezone)
	userTable := models.User{}
	userTable.FirstName = signReq.RegisterUserRequest.FirstName
	userTable.LastName = signReq.RegisterUserRequest.LastName
	userTable.CountryCode = signReq.CountryCode
	userTable.UserTypeID = userType.ID
	userTable.Lang = signReq.Lang
	userTable.Timezone = signReq.Timezone
	userTable.PatientAccessCode = signReq.RegisterUserRequest.AccessCode
	userTable.EmailID = signReq.RegisterUserRequest.EmailID
	userTable.LoggedSrc = signReq.RegisterUserRequest.LoggedSrc
	userTable.DOB = signReq.RegisterUserRequest.DOB
	userTable.SolutionType = signReq.RegisterUserRequest.SolutionType
	userTable.AppID = signReq.RegisterUserRequest.AppID
	userTable.DeviceUUID = signReq.LoginDeviceDetails.DeviceUUID
	userTable.JoinedDate = currentTime
	userTable.Gender = signReq.RegisterUserRequest.Gender
	userTable.BeforeInsert(signReq.Timezone)
	user, errUser := sp.signUpDao.CreateUser(userTable)
	if errUser != nil {
		sp.l.Error("CreateUser Error--", errUser)
		return nil, errUser
	}
	sp.l.Debug("User Created, ID- ", user.ID)

	userAccessTable := models.UserAccess{}
	//productAccess := sp.signUpDao.GetProductAccessFor(signReq.RegisterUserRequest.ProductAccessFor)
	if strings.ToUpper(strings.TrimSpace(signReq.RegisterUserRequest.ProductAccessFor)) == "ALCOCHANGE-DTX" { // Need to write logic
		sp.l.Debug("productAccess.Code Equel - ", signReq.RegisterUserRequest.ProductAccessFor)
		userAccessTable.AlcoChangeDtxAccess = true
	}
	userAccessTable.UserID = user.ID
	userAccessTable.EmailID = user.EmailID
	userAccessTable.MobileNo = user.MobileNo
	userAccessTable.Timezone = user.Timezone
	userAccessTable.PatientAccessCode = user.PatientAccessCode
	userAccessTable.UserTypeID = user.UserTypeID
	userAccessTable.DeviceUUID = user.DeviceUUID
	userAccessTable.FirstName = user.FirstName
	userAccessTable.LastName = user.LastName
	userAccessTable.Password = utils.SHAEncoding(signReq.RegisterUserRequest.Password)
	userAccessTable.BeforeInsert(signReq.Timezone)
	userAccess, errAccess := sp.signUpDao.CreateUserAccess(userAccessTable)
	if errAccess != nil {
		sp.l.Error("CreateUserAccess Error--", errAccess)
		if user.ID != 0 {
			sp.signUpDao.DeleteUser(user.ID)
		}
		return nil, errAccess
	}
	sp.l.Debug("userAccess Created, ID- ", userAccess.ID)
	loginDeviceDetailsTable := models.LoginDeviceDetails{}
	loginDeviceDetailsTable.UserID = user.ID
	loginDeviceDetailsTable.Timezone = user.Timezone
	loginDeviceDetailsTable.Latitude = signReq.RegisterUserRequest.Latitude
	loginDeviceDetailsTable.Longitude = signReq.RegisterUserRequest.Longitude
	loginDeviceDetailsTable.AppID = signReq.RegisterUserRequest.AppID
	loginDeviceDetailsTable.OsVersion = signReq.LoginDeviceDetails.OsVersion
	loginDeviceDetailsTable.OsType = signReq.LoginDeviceDetails.OsType
	loginDeviceDetailsTable.DeviceUUID = signReq.LoginDeviceDetails.DeviceUUID
	loginDeviceDetailsTable.DeviceInfo = signReq.LoginDeviceDetails.DeviceInfo
	loginDeviceDetailsTable.NetworkInfo = signReq.LoginDeviceDetails.NetworkInfo
	loginDeviceDetailsTable.UserAppVersion = signReq.LoginDeviceDetails.UserAppVersion
	loginDeviceDetailsTable.BeforeInsert(signReq.Timezone)
	_, errDD := sp.signUpDao.CreateLoginDeviceDetails(loginDeviceDetailsTable)
	if errDD != nil {
		sp.l.Error("CreateLoginDeviceDetails Error--", errDD)
		if user.ID != 0 { // If Any issues  RollBack
			sp.signUpDao.DeleteUser(user.ID)
			sp.signUpDao.DeleteUserAccess(userAccess.ID)
		}
		return nil, errDD
	}
	userAccessConfirmation, _ := sp.signUpDao.GetUserActionConfirmationUUID(signReq.LoginDeviceDetails.DeviceUUID)
	if userAccessConfirmation.ID != 0 {
		userAccessConfirmation.UserID = user.ID
		userAccessConfirmation.IsSignedUp = true
		userAccessConfirmation.BeforeUpdate(user.Timezone)
		sp.signUpDao.UpdateUserActionConfirmation(*userAccessConfirmation)
	}
	//Response
	signUp := dtos.SignUpResponse{}
	signUp.AccessToken = ""
	signUp.LoggedSrc = user.LoggedSrc
	signUp.DeviceUUID = user.DeviceUUID
	signUp.FirstName = user.FirstName
	signUp.LastName = user.LastName
	signUp.JoinedDate = ""
	signUp.TokenID = ""
	signUp.Email = user.EmailID

	return &signUp, nil
}

func (login *SignUp) UserLogin(signReq dtos.SignInRequest) (*dtos.SignInResponse, error) {
	response := dtos.SignInResponse{}
	if signReq.EmailID == "" || signReq.Password == "" {
		return nil, errors.New("Invalid credentials")
	}
	signReq.Password = utils.SHAEncoding(signReq.Password)
	userAccess, errLogin := login.signUpDao.UserValidation(signReq.EmailID, signReq.Password)
	if errLogin != nil {
		login.l.Error("UserValidation Error-", errLogin)
		isExists := strings.Contains(errLogin.Error(), "no rows in result set")
		if isExists {
			errLogin = errors.New("Invalid credentials")
		}
		return nil, errLogin
	}
	if strings.ToUpper(strings.TrimSpace(signReq.ProductAccessFor)) == "ALCOCHANGE-DTX" {
		login.l.Debug("productAccess.Code Equel - ", signReq.ProductAccessFor)
		userAccess.AlcoChangeDtxAccess = true
	}
	userAccess.BeforeUpdate(signReq.Timezone)

	go login.signUpDao.UpdateUserAccess(*userAccess)

	userLoginDetails, errD := login.signUpDao.GetLoginDeviceDetails(userAccess.UserID)
	if errD != nil {
		login.l.Error("GetLoginDeviceDetails Login Error-", errD)
		return nil, errD
	}

	if userLoginDetails.ID > 0 {
		login.l.Error("userLoginDetails.ID-", userLoginDetails.ID, userLoginDetails.UserID)
		userLoginDetails.AppID = signReq.AppID
		userLoginDetails.Timezone = signReq.Timezone
		userLoginDetails.Latitude = signReq.Latitude
		userLoginDetails.Longitude = signReq.Longitude
		userLoginDetails.OsVersion = signReq.LoginDeviceDetails.OsVersion
		userLoginDetails.OsType = signReq.LoginDeviceDetails.OsType
		userLoginDetails.UserAppVersion = signReq.LoginDeviceDetails.UserAppVersion
		userLoginDetails.DeviceUUID = signReq.LoginDeviceDetails.DeviceUUID
		userLoginDetails.DeviceInfo = signReq.LoginDeviceDetails.DeviceInfo
		userLoginDetails.NetworkInfo = signReq.LoginDeviceDetails.NetworkInfo
		userLoginDetails.BeforeUpdate(signReq.Timezone)
		go login.signUpDao.UpdateLoginDeviceDetails(*userLoginDetails)
	}

	loginLogs := models.LoginLogs{}
	loginLogs.UserID = userAccess.UserID
	loginLogs.AppID = signReq.AppID
	loginLogs.UserAppVersion = signReq.LoginDeviceDetails.UserAppVersion
	loginLogs.Latitude = signReq.Latitude
	loginLogs.Longitude = signReq.Longitude
	loginLogs.OsVersion = signReq.LoginDeviceDetails.OsVersion
	loginLogs.OsType = signReq.LoginDeviceDetails.OsType
	loginLogs.DeviceUUID = signReq.LoginDeviceDetails.DeviceUUID
	loginLogs.DeviceInfo = signReq.LoginDeviceDetails.DeviceInfo
	loginLogs.NetworkInfo = signReq.LoginDeviceDetails.NetworkInfo
	loginLogs.BeforeInsert(signReq.Timezone)
	go login.signUpDao.CreateLoginLog(loginLogs)

	response.EmailID = userAccess.EmailID
	return &response, nil
}
