package main

import (
	"log"
	"mariadb-service/config"
	"mariadb-service/controller"
	"mariadb-service/dbstore"
	kafkaservice "mariadb-service/kafka-service"
	"mariadb-service/router"
	"net/http"
)

func main() {
	gconfig := config.ReadEnv()
	initProducer, _ := kafkaservice.InitProducer()

	cc := &controller.UserController{
		Db:       dbstore.NewGormDatabase(gconfig),
		Producer: initProducer,
	}

	r := router.InitializeRoutes(cc)
	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
