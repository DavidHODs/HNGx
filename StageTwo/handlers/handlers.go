package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DavidHODs/HNGx/StageTwo/db"
)

// Person data structure with assigned struct tags
type Person struct {
	ID        int64     `json:"id" db:"id"`
	FullName  string    `json:"full_name"  db:"full_name"`
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at" db:"created_at" sql:"type:date"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" sql:"type:date"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at" sql:"type:date"`
}

// returned json response struct
type Response struct {
	Msg    string `json:"msg"`
	Data   string `json:"data"`
	Status int    `json:"status"`
}

// CreatePerson creates a new person and add them to the database
func CreatePerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var person Person

	// decodes json data from request body, parses it and stores it in person variable
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		http.Error(res, "error: bad request", http.StatusBadRequest)
		return
	}

	fmt.Println(person.FullName)

	// SQL query to insert a record; avouds SQL injection
	query := `INSERT INTO person (fullName, isDeleted) 
			  VALUES ($1, $2)
			  RETURNING id`

	// executes the query and returns the id of the created person
	err = db.DB.QueryRow(query, person.FullName, false).Scan(&person.ID)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: failed to insert record into database ==> %v", err), http.StatusInternalServerError)
		return
	}

	// Return the created person as JSON
	json.NewEncoder(res).Encode(Response{
		Msg:    "record succesfully created",
		Data:   person.FullName,
		Status: http.StatusCreated,
	})
}