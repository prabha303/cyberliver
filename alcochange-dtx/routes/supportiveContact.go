package routes

import (
	"cyberliver/alcochange-dtx/internals/services/supportiveContactService"
	"cyberliver/alcochange-dtx/sentryaccounts"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func supportiveContact(router *httprouter.Router) {
	router.GET("/v1/acdtx/supportiveContact/relationShip", GetSupportiveContact)
}

// GetSupportiveContact func to send the Supportive Contact to the client
func GetSupportiveContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := supportiveContactService.NewSupportiveContact(rd.l, rd.dbConn)
	res, errW := wl.GetSupportiveContact()
	if errW != nil {
		rd.l.Errorf("GetSupportiveContact - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
