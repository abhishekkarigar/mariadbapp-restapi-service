package router

import (
	"github.com/gorilla/mux"
	"mariadb-service/controller"
)

func InitializeRoutes(controller *controller.UserController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", controller.Post).Methods("POST")
<<<<<<< HEAD
=======
	r.HandleFunc("/api/users/{email}", controller.Delete).Methods("DELETE")
	r.HandleFunc("/api/users/{email}", controller.GetUser).Methods("GET")
>>>>>>> 11ef93d (first commit)
	return r
}
