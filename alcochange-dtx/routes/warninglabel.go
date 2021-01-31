package routes

import (
	"net/http"

	"cyberliver/alcochange-dtx/internals/services/warniglableservice"
	"cyberliver/alcochange-dtx/sentryaccounts"

	"github.com/julienschmidt/httprouter"
)

func warningLabel(router *httprouter.Router) {
	router.GET("/v1/alcochange/warningLabel", GetWarningLabel)
}

func GetWarningLabel(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	wl := warniglableservice.NewWarning(rd.l, rd.dbConn)
	res, errW := wl.GetWarniglableMessage()
	if errW != nil {
		rd.l.Errorf("issue with GetWarningLabel", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
