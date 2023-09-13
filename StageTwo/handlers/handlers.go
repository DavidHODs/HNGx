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

// returned json response struct for a single person
type Response struct {
	Msg       string    `json:"msg"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    int       `json:"status"`
}

// returned json response struct for a list of persons
type ResponseRecord struct {
	Msg    string   `json:"msg"`
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
	query := `INSERT INTO person (fullName) 
			  VALUES ($1)
			  RETURNING id`

	// executes the query and returns the id of the created person
	err = db.DB.QueryRow(query, person.FullName).Scan(&person.ID)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: failed to insert record into database ==> %v", err), http.StatusInternalServerError)
		return
	}

	// Return the created person as JSON
	json.NewEncoder(res).Encode(Response{
		Msg:    "record succesfully created",
		Name:   person.FullName,
		Status: http.StatusCreated,
	})
}

// GetPerson retrieves the details of a person by their ID
func GetPerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Extracts the person's ID from the URL path or request parameters before converting to integer
	personID := mux.Vars(req)["id"]
	ID, _ := strconv.Atoi(personID)

	// SQL query to retrieve a person's details by their ID
	query := `SELECT id, fullName, createdAt, updatedAt FROM person WHERE id = $1 and isDeleted = false`

	var person Person

	// Execute the query and scan the result into the person variable
	err := db.DB.QueryRow(query, ID).Scan(&person.ID, &person.FullName, &person.CreatedAt, &person.UpdatedAt)
	if err != nil {
		http.Error(res, fmt.Sprintf("error: person with ID %d not found", ID), http.StatusNotFound)
		return
	}

	// Returns the person's details as JSON
	json.NewEncoder(res).Encode(Response{
		Msg:       "record succesfully retrieved",
		Name:      person.FullName,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
		Status:    http.StatusOK,
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

	// Check if no records were found
	if len(persons) == 0 {
		http.Error(res, "no record found", http.StatusNotFound)
		return
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(res, fmt.Sprintf("error: row iteration error ==> %v", err), http.StatusInternalServerError)
		return
	}

	// Return the list of persons as JSON
	json.NewEncoder(res).Encode(ResponseRecord{
		Msg:    "records successfully retrieved",
		Status: http.StatusOK,
		Record: persons,
	})
}

// UpdatePerson updates a person's record in the database
func UpdatePerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Extracts the person's ID from the URL path or request parameters before converting to integer
	personID := mux.Vars(req)["id"]
	ID, _ := strconv.Atoi(personID)

	// Creates a struct to hold the updated person data
	var updatedPerson Person
	if err := json.NewDecoder(req.Body).Decode(&updatedPerson); err != nil {
		http.Error(res, "error: bad request", http.StatusBadRequest)
		return
	}

	// SQL query to update a person's record
	query := `UPDATE person SET fullName = $2, createdAt = $3 WHERE id = $1 and isDeleted = false`

	// Execute the query to update the person's information
	_, err := db.DB.Exec(query, ID, updatedPerson.FullName, time.Now())
	if err != nil {
		http.Error(res, fmt.Sprintf("error: failed to update record ==> %v", err), http.StatusInternalServerError)
		return
	}

	// Return a success response
	json.NewEncoder(res).Encode(Response{
		Msg:       "record successfully updated",
		Name:      updatedPerson.FullName,
		UpdatedAt: time.Now(),
		Status:    http.StatusOK,
	})
}
