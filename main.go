package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
)

type Language struct {
	Id string `json:"Id,omitempty"`
	Name string `json:"Name, omitempty"`
}

var languages []Language

func main() {
	router := mux.NewRouter()

	languages = append(languages, Language{Id: "1", Name: "Javascript"})
	languages = append(languages, Language{Id: "2", Name: "PHP"})
	languages = append(languages, Language{Id: "3", Name: "Java"})
	languages = append(languages, Language{Id: "4", Name: "C#"})

	router.HandleFunc("/languages", GetLanguage).Methods("GET")
	router.HandleFunc("/languages/{id}", GetLanguage).Methods("GET")
	router.HandleFunc("/languages/{id}", PostLanguage).Methods("POST")
	router.HandleFunc("/languages/{id}", DeleteLanguage).Methods("DELETE")

	fmt.Println("Server run")
	log.Fatal(http.ListenAndServe(":4000", router))

}

func GetLanguages(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(languages)
}

func GetLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range languages {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Language{})
}

func PostLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var language Language
	_ = json.NewDecoder(req.Body).Decode(&language)
	language.Id = params["id"]
	languages = append(languages, language)
	json.NewEncoder(w).Encode(languages)
}

func DeleteLanguage(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range languages {
		if item.Id == params["id"] {
			languages = append(languages[:index], languages[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(languages)
}

