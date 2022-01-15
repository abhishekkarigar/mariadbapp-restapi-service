package controller

import (
	"encoding/json"
	"fmt"
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
		return
	}
	w.Write([]byte("record added successfully"))
}
