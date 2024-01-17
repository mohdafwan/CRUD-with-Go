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
var director []Director

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// GET THE PERTICULAR MOVIES
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //get access all the request for the param
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) //get access all the request for the parem knower
	for _, v := range movies {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
}

func createMovies(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-type", "application/json")
	q.WriteHeader(http.StatusOK)
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)
	json.NewEncoder(q).Encode(movies)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "43522",
		Title: "Movie One",
		Director: &Director{
			Firstname: "TheName",
			Lastname:  "TheLastName",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "43523",
		Title: "Movie Two",
		Director: &Director{
			Firstname: "someName",
			Lastname:  "NP.m",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting Server at PORT 8000")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
