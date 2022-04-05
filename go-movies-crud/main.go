package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	setMoviesData(&movies)

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at localhost:8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func getMovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(movies)

}

func getMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(response).Encode(movie)
			return
		}
	}
}

func createMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rangeIn(100000, 999999))
	movies = append(movies, movie)
	json.NewEncoder(response).Encode(movie)
}

func updateMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			deleteMovieAtIndex(&movies, index)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(response).Encode(movie)
			return
		}
	}

}

func deleteMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			deleteMovieAtIndex(&movies, index)
			break
		}
	}
	json.NewEncoder(response).Encode(movies)

}

func deleteMovieAtIndex(movies *[]Movie, index int) {
	(*movies)[index] = (*movies)[len((*movies))-1]
	*movies = (*movies)[:len(*movies)-1]
}

func setMoviesData(movies *[]Movie) {
	filename := "data/movies.csv"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	if len(lines[0]) < 5 {
		panic("Malformed input.")
	}

	*movies = make([]Movie, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			continue
		}

		(*movies)[i-1] =
			Movie{
				ID:    strconv.Itoa(rangeIn(100000, 999999)),
				ISBN:  line[1],
				Title: line[2],
				Director: &Director{
					FirstName: line[3],
					LastName:  line[4],
				},
			}
	}
}

func rangeIn(lowerbound, upperbound int) int {
	return lowerbound + rand.Intn(upperbound-lowerbound)
}
