package router

import (
	"github.com/gorilla/mux"
	"mariadb-service/controller"
)

func InitializeRoutes(controller *controller.UserController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", controller.Post).Methods("POST")
	return r
}
