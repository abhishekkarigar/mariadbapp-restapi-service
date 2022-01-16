package controller

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
=======
	"github.com/gorilla/mux"
>>>>>>> 11ef93d (first commit)
	"mariadb-service/domain"
	"net/http"
)

type UserController struct {
	Db domain.UserStore
}

func (user *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	all, err := user.Db.GetAll()
	if err != nil {
		return
	}
	marshal, err := json.Marshal(all)
	if err != nil {
		return
	}
<<<<<<< HEAD
=======
	w.Header().Set("Content-Type", "application/json")
>>>>>>> 11ef93d (first commit)
	w.Write(marshal)
}

func (user *UserController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newuser domain.User
	_ = json.NewDecoder(r.Body).Decode(&newuser)

	fmt.Println(newuser.Name)
	fmt.Println(newuser.Email)
	err := user.Db.CreateUser(newuser)
	if err != nil {
<<<<<<< HEAD
=======
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("duplicate email"))
>>>>>>> 11ef93d (first commit)
		return
	}
	w.Write([]byte("record added successfully"))
}
<<<<<<< HEAD
=======

func (user *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	if email == "" {
		w.WriteHeader(400)
		w.Write([]byte("email is missing"))
		return
	}
	err := user.Db.DeleteUser(email)
	if err != nil {
		return
	}
	w.Write([]byte("record deleted successfully"))
}

func (user *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	var getUser domain.User
	params := mux.Vars(r)
	email := params["email"]
	fmt.Println(email)
	if email == "" {
		w.WriteHeader(400)
		w.Write([]byte("email is missing"))
		return
	}
	user1, err := user.Db.GetUser(getUser, email)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(user1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}
>>>>>>> 11ef93d (first commit)
