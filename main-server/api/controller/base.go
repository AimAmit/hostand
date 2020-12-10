package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(addr string) *mux.Router {
	router := InitializeRoutes()

	fmt.Printf("Listening to port %s\n", addr)
	defer log.Fatal(http.ListenAndServe(addr, router))
	return router
}
