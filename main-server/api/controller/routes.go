package controller

import "github.com/gorilla/mux"

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/app", CreateApp).Methods("POST")
	return r
}
