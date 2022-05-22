package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	_ = json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newmovie Movie
	json.NewDecoder(r.Body).Decode(&newmovie)
	newmovie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, newmovie)
	json.NewEncoder(w).Encode(newmovie)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var updatedmovie Movie
			json.NewDecoder(r.Body).Decode(&updatedmovie)
			updatedmovie.ID = params["id"]
			movies = append(movies, updatedmovie)
			json.NewEncoder(w).Encode(updatedmovie)
			return
		}
	}
}

func initialize() {
	movies = append(movies, Movie{ID: "1", Isbn: "42341", Title: "Never backdown", Director: &Director{Firstname: "John", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "2", Isbn: "14514", Title: "Dune", Director: &Director{Firstname: "Mr X", Lastname: "Mr Y"}})
	movies = append(movies, Movie{ID: "5", Isbn: "12412511525", Title: "The Beach", Director: &Director{Firstname: "Yolo", Lastname: "yaw"}})
}
func main() {
	port := ":8000"
	r := mux.NewRouter()
	initialize()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server at port %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}
