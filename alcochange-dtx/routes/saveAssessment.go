package routes

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/services/saveAssessmentService"
	"cyberliver/alcochange-dtx/sentryaccounts"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func saveAssessment(router *httprouter.Router) {
	router.POST("/v1/acdtx/saveAssessment/:uid", SaveAssessment)
}

func SaveAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)
	reqBody := dtos.SaveAssessmentRequest{}

	userID, isErr := GetIDFromParams(w, r, "uid")
	if isErr != nil {
		rd.l.Error("SaveAssessment userID --", userID, isErr)
		sentryaccounts.SentryLogExceptions(isErr)
		writeJSONMessage(isErr.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	isJSON, jErr := parseJSONWithError(w, r.Body, &reqBody)
	if !isJSON {
		rd.l.Error("SaveAssessment json Error --", isJSON, jErr)
		sentryaccounts.SentryLogExceptions(jErr)
		writeJSONMessage(jErr.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	sa := saveAssessmentService.NewSaveAssessment(rd.l, rd.dbConn)
	errW := sa.SaveAssessmentDetails(reqBody, userID)
	if errW != nil {
		rd.l.Errorf("SaveAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}

	writeJSONStruct(res, http.StatusOK, rd)
}
