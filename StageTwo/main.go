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

	// creates person's record
	router.HandleFunc("/api", handlers.CreatePerson).Methods("POST")

	// retrieves a person's record
	router.HandleFunc("/api/{user_id:[0-9]+}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/api/{user_name}", handlers.GetPersonByName).Methods("GET")

	// retrieves all person record
	router.HandleFunc("/api", handlers.GetAllPersons).Methods("GET")

	// updates a person 's record
	router.HandleFunc("/api/{user_id:[0-9]+}", handlers.UpdatePerson).Methods("PATCH")
	router.HandleFunc("/api/{user_name}", handlers.UpdatePersonByName).Methods("PATCH")

	// soft deleletes a person ==> isDelete flag is set to true but record still exists
	router.HandleFunc("/api/soft_delete/{user_id:[0-9]+}", handlers.SoftDeletePerson).Methods("PATCH")
	router.HandleFunc("/api/soft_delete/{user_name}", handlers.SoftDeletePersonByName).Methods("PATCH")

	// wipes of a person's record
	router.HandleFunc("/api/{user_id:[0-9]+}", handlers.HardDeletePerson).Methods("DELETE")
	router.HandleFunc("/api/{user_name}", handlers.HardDeletePersonByName).Methods("DELETE")

	// renders generated postman (with newmann) test report and documentation
	router.PathPrefix("/api/static/").Handler(http.StripPrefix("/api/static/", http.FileServer(http.Dir("."))))

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Printf("app running on port %d\n", port)
	log.Fatal(srv.ListenAndServe())
}
