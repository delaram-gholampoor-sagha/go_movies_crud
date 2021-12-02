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
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

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
	json.NewEncoder(w).Encode(r)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// then we look through the movies one by one then find that movie then encode it into json and then send it into front end or post man
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(r)
		}
	}
}

func createMovie(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-type", "application/json")
	// the new movie that has come from the json has been stored into this variable
	var movie Movie
	// we decoded this from json into a fromat that is readable for golang
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// to create a new movie we haave to create a new id
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(r)
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "43256",
		Title: "Movie one",
		Director: &Director{
			Firstname: "delaram",
			LastName:  "gholampoor sagha",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("Get")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
