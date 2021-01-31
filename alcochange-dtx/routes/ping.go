package routes

import (
	"net/http"

	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/internals/services/ping"

	"github.com/julienschmidt/httprouter"
)

func setPingRoutes(router *httprouter.Router) {
	router.GET("/ping", Ping)
}

var res dtos.ResStruct

// Ping godoc
// @Summary ping api
// @Description do ping
// @Tags ping
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResStruct
// @Failure 500 {object} dtos.Res500Struct
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	p := ping.New(rd.l, rd.dbConn)
	err := p.Ping()
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
	} else {
		writeJSONMessage("pong", MSG, http.StatusOK, rd)
	}
}
