package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"` // Json handle name
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (user *User) ToJson() ([]byte, error) {
	return json.Marshal(user) // Return a Json encoding
}

type MetaData interface{}
