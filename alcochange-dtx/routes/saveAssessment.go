package routes

import (
	"cyberliver/alcochange-dtx/dtos"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func saveAssessment(router *httprouter.Router) {
	router.POST("/v1/acdtx/save/assessment", SaveAssessment)
}

func SaveAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	reqBody := dtos.SaveAssessmentRequest{}
	isJSON, jErr := parseJSONWithError(w, r.Body, &reqBody)
	if !isJSON {
		rd.l.Error("SaveAssessment json Error --", isJSON, jErr)
		writeJSONMessage(jErr.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	sa := saveAssessmentService.NewSaveAssessment(rd.l, rd.dbConn)
	res, errW := sa.SaveAssessmentDetails(reqBody)
	if errW != nil {
		rd.l.Errorf("SaveAssessment - Error : ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
