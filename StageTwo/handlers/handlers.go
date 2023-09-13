package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DavidHODs/HNGx/StageTwo/db"
	"github.com/gorilla/mux"
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
	Msg    string   `json:"msg"`
	Data   Person   `json:"data"`
	Status int      `json:"status"`
	Record []Person `json:"record"`
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
		Msg: "record succesfully created",
		Data: Person{
			FullName: person.FullName,
		},
		Status: http.StatusCreated,
	})
}

// GetPerson retrieves the details of a person by their ID
func GetPerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	personID := mux.Vars(req)["id"]
	ID, _ := strconv.Atoi(personID)

	// SQL query to retrieve a person's details by their ID
	query := `SELECT id, fullName, createdAt, updatedAt FROM person WHERE id = $1 and isDeleted = false`

	var person Person

	// Execute the query and scan the result into the person variable
	err := db.DB.QueryRow(query, ID).Scan(&person.ID, &person.FullName, &person.CreatedAt)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: person with ID %d not found", ID), http.StatusNotFound)
		return
	}

	// Returns the person's details as JSON
	json.NewEncoder(res).Encode(Response{
		Msg: "record succesfully retrieved",
		Data: Person{
			ID:        person.ID,
			FullName:  person.FullName,
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
		},
		Status: http.StatusOK,
	})
}

// GetAllPersons retrieves all person records from the database
func GetAllPersons(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// SQL query to retrieve all person records
	query := `SELECT id, fullName, createdAt, updatedAt FROM person WHERE isDeleted = false`

	// Execute the query to fetch all records
	rows, err := db.DB.Query(query)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: failed to fetch records ==> %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the results
	var persons []Person

	// Iterate through the rows and scan each record into a Person struct
	for rows.Next() {
		var person Person
		err := rows.Scan(&person.ID, &person.FullName, &person.CreatedAt, &person.UpdatedAt)
		if err != nil {
			http.Error(res, fmt.Sprintf("error: failed to scan record ==> %v", err), http.StatusInternalServerError)
			return
		}
		persons = append(persons, person)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(res, fmt.Sprintf("error: row iteration error ==> %v", err), http.StatusInternalServerError)
		return
	}

	// Return the list of persons as JSON
	json.NewEncoder(res).Encode(Response{
		Msg:    "records successfully retrieved",
		Record: persons,
		Status: http.StatusOK,
	})
}
