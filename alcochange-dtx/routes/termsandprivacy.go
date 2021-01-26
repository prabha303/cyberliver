package routes

import (
	"ecargoware/alcochange-dtx/internals/services/termsandprivacyservice"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func termsAndPrivacy(router *httprouter.Router) {
	router.GET("/v1/alcochange/termsAndPrivacy", GetTermsAndPrivacy)
}

// GetTermsAndPrivacy func to send the terms and privacy to the client
func GetTermsAndPrivacy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := termsandprivacyservice.NewTermsAndPrivacy(rd.l, rd.dbConn)
	res, errW := wl.GetTermsAndPrivacyMessage()
	if errW != nil {
		rd.l.Errorf("GetTermsAndPrivacy - Error : ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
