package routinggorillamux

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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Conntent-Type", "application/json")
	json.NewEncoder(q).Encode(movies)

}

func deleteMovie(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(q).Encode(movies)

}

func createMovie(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(q).Encode(movie)

}

func updateMovies(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// movies = append(movies[:index],movies[index+1]...)
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(q).Encode(movie)
			return
		}

	}

}

func getMovie(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(q).Encode(item)
			return
		}
	}

}

func GorillaMuxApiImplementation() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "123",
		Isbn:  "123435po",
		Title: "Malika tome",
		Director: &Director{
			FirstName: "Elizabeth ",
			LastName:  "Alison",
		},
	})

	movies = append(movies, Movie{
		ID:    "10",
		Isbn:  "12we235",
		Title: "Movie 7",
		Director: &Director{
			FirstName: "Maina",
			LastName:  "King'ong'i",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000...")

	log.Fatal(http.ListenAndServe("localhost:8000", r))

}
