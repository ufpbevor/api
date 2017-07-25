package core

import "github.com/gorilla/mux"

//Router with NewRelic
func Router(h *Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/zen", h.HandleZen).Methods("GET")
	router.HandleFunc("/version", h.HandleVersion).Methods("GET")

	return router
}
