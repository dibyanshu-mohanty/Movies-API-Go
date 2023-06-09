package main

import (
	"encoding/json"
	"fmt"
	"log" // Encoding data to JSON
	"math/rand"
	"strconv"

	// Makes a random Function
	"net/http" // Creates a Server

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json :"id"`
	Isbn string `json :"isbn"`
	Title string `json :"title"`
	Director *Director `json :"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie


func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	} 
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


// Currently Using Temporary Database
func updateMovie(w http.ResponseWriter, r* http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()


	movies = append(movies, Movie{ID: "1", Isbn: "438027", Title: "The Kerela Story", Director: &Director{FirstName: "Sudipto", LastName: "Sen"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438028", Title: "Kashmir Files", Director: &Director{FirstName: "Vivek", LastName: "Agnihotri"}})
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Startign Server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}