package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // http router
)

// set up movie and director data as structs
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// define a slice of movies
var movies []Movie

// endpoints
func getMovies(w http.ResponseWriter, r *http.Request) {
	// set the header tag "content-type" to be a JSON object
	w.Header().Set("Content-Type", "application/json")

	// encode all movies as json object, write to w
	json.NewEncoder(w).Encode(movies)

	return
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// parse out ID from the request
	params := mux.Vars(r)

	// delete movie w/ ID from slice
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...) // append second half of list to prior half, skip over idx
			break
		}
	}

	// return remaining movies
	json.NewEncoder(w).Encode(movies)

	return
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// set the header
	w.Header().Set("Content-Type", "application/json")

	// parse out ID from the request (get all route variables)
	params := mux.Vars(r)

	// find movie w/ ID
	for _, item := range movies {
		if item.ID == params["id"] {
			// encode movie info in JSON and send
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func main() {
	// initialize router (handler)
	r := mux.NewRouter()

	// initialize the slice of movies
	movies = append(movies, Movie{ID: "0", ISBN: "438227", Title: "Parasite", Director: &Director{Firstname: "Joon-Ho", Lastname: "Bong"}})
	movies = append(movies, Movie{ID: "1", ISBN: "45455", Title: "Moonlight", Director: &Director{Firstname: "Barry", Lastname: "Jenkins"}})

	// create endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies", createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	// start server
	fmt.Printf("Starting server at port 8000...\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
