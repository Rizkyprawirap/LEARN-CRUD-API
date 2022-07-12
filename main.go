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

// Creating a new type of "Book" components
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Creating a new type of "Director" components
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// slice if movie
var movies []Movie

// Create Get Movie Function
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the type of data format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Search on the slice
	for index, item := range movies {
		// if slice match with the params
		if item.ID == params["id"] {
			// if the movie found, it will deleted, then the rest of it will append back
			movies = append(movies[:index], movies[index+1:]...)
			break

		} else {
			fmt.Println("Movie not found!")
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Creating movie id by random int then convert to string
	movie.ID = strconv.Itoa((rand.Intn(100000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Set json content type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	// loop over the movie range
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete movie by id
			movies = append(movies[:index], movies[index+1:]...)

			// add a new movie - added
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	// Declare a new router from the Gorilla mux package
	router := mux.NewRouter()

	// Insert Movies
	movies = append(movies,
		Movie{
			ID:       "1",
			ISBN:     "123456",
			Title:    "First Movie",
			Director: &Director{FirstName: "Iky", LastName: "Cat"},
		})
	// Insert Movies
	movies = append(movies,
		Movie{
			ID:       "2",
			ISBN:     "654321",
			Title:    "Second Movie",
			Director: &Director{FirstName: "Putra", LastName: "Mirea"},
		})

	// Declare the endpoints
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
