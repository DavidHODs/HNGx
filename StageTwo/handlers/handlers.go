package handlers

import "time"

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
