package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article about Programming Language
type Language struct {
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// simulate a database of languages
var Languages []Language

// POST /language
func createNewLanguage(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var language Language
	err := json.Unmarshal(reqBody, &language)
	if err != nil { // cannot create a new Language
		fmt.Println(err)
		return
	}
	fmt.Println("createlanguage: " + language.Id)
	Languages = append(Languages, language)
}

// GET /language/{id}
func getLanguage(w http.ResponseWriter, r *http.Request) {
	// get the ID from the URL
	vars := mux.Vars(r)
	key := vars["id"]
	for _, language := range Languages {
		if language.Id == key {
			fmt.Println("getLanguage: " + key)
			json.NewEncoder(w).Encode(language)
		}
	}
}

// PUT /language/{id}
func updateLanguage(w http.ResponseWriter, r *http.Request) {
	// get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]
	// get the body of our PUT request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var language Language
	err := json.Unmarshal(reqBody, &language)
	if err != nil {
		fmt.Println(err)
		return
	}
	for index, item := range Languages {
		if item.Id == id {
			fmt.Println("updateLanguage: " + id)
			Languages = append(Languages[:index], Languages[index+1:]...)
			Languages = append(Languages, language)
		}
	}
}

// DELETE /language/{id}
func deleteLanguage(w http.ResponseWriter, r *http.Request) {
	// get the ID from the URL
	vars := mux.Vars(r)
	id := vars["id"]
	for index, language := range Languages {
		if language.Id == id {
			fmt.Println("deleteLanguage: " + id)
			Languages = append(Languages[:index], Languages[index+1:]...)
		}
	}
}

// GET ALL languageS
func returnAllLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returnAllLanguage")
	json.NewEncoder(w).Encode(Languages)
}

// Homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Here !")
}

// Handle the requests
func handleRequests() {
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// attach the homePage route and handler to the router
	myRouter.HandleFunc("/", homePage)
	// attach the returnAllLanguages route and handler to the router
	myRouter.HandleFunc("/languages", returnAllLanguage)
	myRouter.HandleFunc("/language", createNewLanguage).Methods("POST")     // CREATE
	myRouter.HandleFunc("/language/{id}", getLanguage).Methods("GET")       // READ
	myRouter.HandleFunc("/language/{id}", updateLanguage).Methods("PUT")    // UPDATE
	myRouter.HandleFunc("/language/{id}", deleteLanguage).Methods("DELETE") // DELETE
	// start the server on port 8080 and log any errors
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	// init the language array
	Languages = []Language{
		{Id: "1", Title: "Go", Desc: "A nice language !", Content: "Some text about Go"},
		{Id: "2", Title: "Rust", Desc: "a cool language", Content: "Some text about Rust"},
	}
	println("Listening on http://localhost:8080")
	handleRequests()
}
