package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DavidHODs/HNGx/StageTwo/db"
	"github.com/DavidHODs/HNGx/StageTwo/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	portStr, err := db.LoadEnv("APP_PORT")
	if err != nil {
		log.Fatal(err)
	}

	port, _ := strconv.Atoi(portStr)

	router := mux.NewRouter()
	router.HandleFunc("/api", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/{id:[0-9]+}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/api", handlers.GetAllPersons).Methods("GET")
	router.HandleFunc("/api/{id:[0-9]+}", handlers.UpdatePerson).Methods("PATCH")
	router.HandleFunc("/api/soft-delete/{id:[0-9]+}", handlers.SoftDeletePerson).Methods("PATCH")
	router.HandleFunc("/api/hard-delete/{id:[0-9]+}", handlers.HardDeletePerson).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Printf("app running on port %d\n", port)
	log.Fatal(srv.ListenAndServe())
}
