package routes

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/services/signUp"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func signINAndUp(router *httprouter.Router) {
	router.POST("/v1/user/registration", SignUp)
	router.POST("/v1/acdtx/login", Login)
}

// GetTermsAndPrivacy func to send the terms and privacy to the client
func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	reqBody := dtos.SignUpRequest{}
	isJSON, jErr := parseJSONWithError(w, r.Body, &reqBody)
	if !isJSON {
		rd.l.Error("SignUp json Error --", isJSON, jErr)
		writeJSONMessage(jErr.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	sp := signUp.NewSignUp(rd.l, rd.dbConn)
	res, errW := sp.UserSignUp(reqBody)
	if errW != nil {
		rd.l.Errorf("SignUp - Error : ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	reqBody := dtos.SignInRequest{}
	isJSON, jErr := parseJSONWithError(w, r.Body, &reqBody)
	if !isJSON {
		rd.l.Error("Login json Error --", isJSON, jErr)
		writeJSONMessage(jErr.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	sp := signUp.NewSignUp(rd.l, rd.dbConn)
	res, errW := sp.UserLogin(reqBody)
	if errW != nil {
		rd.l.Errorf("Login - Error : ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
