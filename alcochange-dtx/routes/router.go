package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

// Prevent abnormal shutdown while panic
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				log.Print(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Put params in context for sharing them between handlers
func wrapHandler(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

//RouterConfig function
func RouterConfig() (router *httprouter.Router) {
	router = httprouter.New()
	router.PanicHandler = panicHandler

	//indexHandlers := alice.New(recoverHandler)

	setPingRoutes(router)
	warningLabel(router)
	updateUserActionConfirmation(router)
	termsAndPrivacy(router)
	patientAccessCode(router)

	return
}

func panicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	fmt.Println("(alcochange-dtx)Recovering from panic-Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
