package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

type personInfo struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	DOB             string `json:"dob"`
	AddressAndPhone string `json:"address_and_phone"`
}

var db1ByID = make(map[int]personInfo)

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var person personInfo
	json.Unmarshal(body, &person)

	db1ByID[person.ID] = person

	fmt.Println(db1ByID)
}

func getByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	person := db1ByID[id]
	JsonMessage, _ := json.Marshal(person)
	w.Write(JsonMessage)
	fmt.Println(string(JsonMessage))
}

func updateByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	requestBody, _ := io.ReadAll(r.Body)

	person := db1ByID[id]
	json.Unmarshal(requestBody, &person)

	db1ByID[id] = person
	fmt.Println(db1ByID)
}

func deleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(path.Base(r.URL.Path))
	delete(db1ByID, id)
	fmt.Println(db1ByID)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/create", http.HandlerFunc(createHandler))
	mux.Handle("/get/", http.HandlerFunc(getByIDHandler))
	mux.Handle("/update/", http.HandlerFunc(updateByIDHandler))
	mux.Handle("/delete/", http.HandlerFunc(deleteByIDHandler))

	fmt.Println(mux) //map[/create:{0x115e8a0 /create} /delete/:{0x115ef00 /delete/} /get/:{0x115ea00 /get/} /update/:{0x115ec60 /update/}]

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
