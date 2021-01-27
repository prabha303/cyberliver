package routes

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/services/patientaccesscodeservice"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func patientAccessCode(router *httprouter.Router) {
	router.POST("/v1/verifyPatientAccessCode", VerifyPatientAccessCode)
}

func VerifyPatientAccessCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rd := logAndGetContext(w, r)

	reqBody := dtos.PatientAccessCodeReq{}
	if !parseJSON(w, r.Body, &reqBody) {
		return
	}

	pac := patientaccesscodeservice.NewPatientAccessCode(rd.l, rd.dbConn)
	res, errW := pac.VerifyPatientByAccessCode(reqBody)
	if errW != nil {
		rd.l.Errorf("issue with VerifyPatientAccessCode ", errW.Error())
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)

}
