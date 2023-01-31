package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie;


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode response as json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// define route paths
	r := mux.NewRouter()

	// add some default movies to the slice
	movies = append(movies, Movie{ID: "1", Isbn: "435890", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	movies = append(movies, Movie{ID: "2", Isbn: "475823", Title: "Movie Two", Director: &Director{Firstname: "George", Lastname: "Green"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) //throw error if server can't start
}