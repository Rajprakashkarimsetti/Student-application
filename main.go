package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rajprakash/student/handler"
	"github.com/rajprakash/student/metrics"
	"github.com/rajprakash/student/service"
	store2 "github.com/rajprakash/student/store"
	"log"
	"net/http"
)

func main() {
	dns := "root:Raj@12345@tcp(127.0.0.1:3306)/learning"

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("could not able to connect to db", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := store2.NewStudentStore(db)
	service := service.NewStudentService(store)
	newHandler := handler.NewHandler(service)

	metrics.Init() // just initialize counters, don't register /metrics here

	// Create router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/student/post", newHandler.Post).Methods("POST")
	r.HandleFunc("/student", newHandler.GetAll).Methods("GET")
	r.HandleFunc("/student/{id}", newHandler.GetById).Methods("GET")
	r.HandleFunc("/student/{id}", newHandler.Delete).Methods("DELETE")
	r.HandleFunc("/student/{id}", newHandler.Put).Methods("PUT")

	// Register Prometheus metrics route here ONLY ONCE
	r.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
