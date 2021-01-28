package routes

import (
	"ecargoware/alcochange-dtx/dtos"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func signINAndUp(router *httprouter.Router) {
	router.GET("/v1/user/registration", SignUp)
}

// GetTermsAndPrivacy func to send the terms and privacy to the client
func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//rd := logAndGetContext(w, r)

	reqBody := dtos.SignUpRequest{}
	if !parseJSON(w, r.Body, &reqBody) {
		return
	}

}
