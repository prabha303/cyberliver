package routes

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/services/useractionconfirmation"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func updateuserActionConfirmation(router *httprouter.Router) {
	router.POST("/v1/action/confirm", UpdateUserActionConfirmation)
}

func UpdateUserActionConfirmation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	ua := useractionconfirmation.NewUserActionConfirm(rd.l, rd.dbConn)
	reqBody := dtos.UserActionConfirmationReq{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rd.l.Errorf("UpdateuserActionConfirmation req error: ", err.Error())
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		rd.l.Errorf("UpdateuserActionConfirmation unmarshal error: ", err.Error())
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	res, errW := ua.UserActionConfirmation(reqBody)
	if errW != nil {
		rd.l.Errorf("issue with GetWarningLabel ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
