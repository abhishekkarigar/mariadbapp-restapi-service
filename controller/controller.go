package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"mariadb-service/domain"
	kafkaservice "mariadb-service/kafka-service"
	"net/http"
)

type UserController struct {
	Db       domain.UserStore
	Producer sarama.SyncProducer
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
	w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("duplicate email"))
		return
	}
	kafkaservice.Publish("new record added succesfully", user.Producer)
	w.Write([]byte("record added successfully"))
}

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
	kafkaservice.Publish("record deleted succesfully", user.Producer)
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
