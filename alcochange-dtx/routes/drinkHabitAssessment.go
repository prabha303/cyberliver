package routes

import (
	"cyberliver/alcochange-dtx/internals/services/drinkHabitAssessmentService"
	"cyberliver/alcochange-dtx/sentryaccounts"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func drinkHabitAssessment(router *httprouter.Router) {
	router.GET("/v1/acdtx/drinkHabitAssessment", GetDrinkHabitAssessment)
}

// GetDrinkHabitAssessment func to send the Drink Habit Assessment to the client
func GetDrinkHabitAssessment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// userID, isErr := controllers.GetIDFromParams(w, r, "id")
	// if !isErr {
	// 	return
	// }
	rd := logAndGetContext(w, r)
	wl := drinkHabitAssessmentService.NewDrinkHabitAssessment(rd.l, rd.dbConn)
	res, errW := wl.GetDrinkHabitAssessmentMessage()
	if errW != nil {
		rd.l.Errorf("GetDrinkHabitAssessment - Error : ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
