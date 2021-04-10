package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Router routes the different requests
func Router() (*mux.Router, error) {

	router := mux.NewRouter()
	router.Use(CorsHandler)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	getRouter := router.Methods(http.MethodGet).Subrouter().StrictSlash(true)

	postRouter.HandleFunc("/api/patient_profile", nil)
	getRouter.HandleFunc("/api/patient_profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "teleHealth")
	})

	return router, nil
}
