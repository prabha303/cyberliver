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
	GetRoleForSignup(signupCode string) models.Role
	GetProductAccessFor(productAccess string) models.ProductAccess
	CreateUser(user models.Users) (*models.Users, error)
	CreateUserAccess(user models.UserAccess) (*models.UserAccess, error)
	DeleteUser(ID int64) bool
}

func (sp *SignUp) EmailIDExists(email string) bool {
	users := []models.Users{}
	err := sp.dbConn.Model(&users).Where("LOWER(email_id) = LOWER(?)", email).Select()
	if err != nil {
		sp.l.Error("EmailIDExists Error", err.Error())
	}
	return len(users) > 0
}

func (sp *SignUp) GetRoleForSignup(signupCode string) models.Role {
	role := models.Role{}
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

func (sp *SignUp) CreateUser(user models.Users) (*models.Users, error) {
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
	users := models.Users{}
	_, insErr := sp.dbConn.Model(&users).Where("id = ? ", ID).Delete()
	if insErr != nil {
		sp.l.Error("DeleteUser Error", insErr.Error())
	}
	return true
}
