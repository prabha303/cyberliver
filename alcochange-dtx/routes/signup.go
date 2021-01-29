package routes

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/services/signUp"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func signINAndUp(router *httprouter.Router) {
	router.POST("/v1/user/registration", SignUp)
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
		rd.l.Errorf("GetTermsAndPrivacy - Error : ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
