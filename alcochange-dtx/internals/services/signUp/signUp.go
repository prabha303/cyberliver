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
	role := sp.signUpDao.GetRoleForSignup(signReq.RegisterUserRequest.SignUpFor)
	if role.ID == 0 {
		return nil, errors.New("Invalid Value signUpFor")
	}

	currentTime, _ := utils.CurrentTimeWithZone(signReq.Timezone)
	userTable := models.Users{}
	userTable.FirstName = signReq.RegisterUserRequest.FirstName
	userTable.LastName = signReq.RegisterUserRequest.LastName
	userTable.CountryCode = signReq.CountryCode
	userTable.RoleID = role.ID
	userTable.Lang = signReq.Lang
	userTable.Timezone = signReq.Timezone
	userTable.PatientAccessCode = signReq.RegisterUserRequest.AccessCode
	userTable.EmailID = signReq.RegisterUserRequest.EmailID
	userTable.LoggedSrc = signReq.RegisterUserRequest.LoggedSrc
	userTable.DOB = signReq.RegisterUserRequest.DOB
	userTable.SolutionType = signReq.RegisterUserRequest.SolutionType
	userTable.AppID = signReq.RegisterUserRequest.AppID
	userTable.UUID = signReq.LoginDeviceDetails.DeviceUUID
	userTable.JoinedDate = currentTime
	userTable.Gender = signReq.RegisterUserRequest.Gender
	userTable.BeforeInsert(signReq.Timezone)
	users, errUser := sp.signUpDao.CreateUser(userTable)
	if errUser != nil {
		sp.l.Error("CreateUser Error--", errUser)
		return nil, errUser
	}
	sp.l.Debug("User Created, ID- ", users.ID)

	userAccessTable := models.UserAccess{}
	//productAccess := sp.signUpDao.GetProductAccessFor(signReq.RegisterUserRequest.ProductAccessFor)
	if strings.ToUpper(strings.TrimSpace(signReq.RegisterUserRequest.ProductAccessFor)) == "ALCOCHANGE-DTX" { // Need to write logic
		sp.l.Debug("productAccess.Code Equel - ", signReq.RegisterUserRequest.ProductAccessFor)
		userAccessTable.AlcoChangeDtxAccess = true
	}
	userAccessTable.UsersID = users.ID
	userAccessTable.EmailID = users.EmailID
	userAccessTable.MobileNo = users.MobileNo
	userAccessTable.Timezone = users.Timezone
	userAccessTable.PatientAccessCode = users.PatientAccessCode
	userAccessTable.RoleID = users.RoleID
	userAccessTable.UUID = users.UUID
	userAccessTable.FirstName = users.FirstName
	userAccessTable.LastName = users.LastName
	userAccessTable.Password = utils.SHAEncoding(signReq.RegisterUserRequest.Password)
	userAccessTable.BeforeInsert(signReq.Timezone)

	userAccess, errAccess := sp.signUpDao.CreateUserAccess(userAccessTable)
	if errAccess != nil {
		sp.l.Error("CreateUserAccess Error--", errAccess)
		if users.ID != 0 {
			sp.signUpDao.DeleteUser(users.ID)
		}
		return nil, errAccess
	}
	sp.l.Debug("userAccess Created, ID- ", userAccess.ID)

	//Response
	signUp := dtos.SignUpResponse{}
	signUp.AccessToken = ""
	signUp.LoggedSrc = users.LoggedSrc
	signUp.DeviceUUID = users.UUID
	signUp.FirstName = users.FirstName
	signUp.LastName = users.LastName
	signUp.JoinedDate = ""
	signUp.TokenID = ""
	signUp.Email = users.EmailID

	return &signUp, nil
}
