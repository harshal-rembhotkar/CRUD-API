package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//"google.golang.org/api/indexing/v3"
	//"google.golang.org/genproto/googleapis/cloud/aiplatform/v1beta1/schema/predict/params"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	// Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json: "firstname"`
	Lastname string `json : "lastname"`
}

var movies []Movie

func main(){
r:= mux.NewRouter()

movies = append(movies, Movie{ID:"1",Isbn: "438",Director: &Director{Firstname: "B", Lastname: "Rembhotkar"}})
movies = append(movies, Movie{ID:"2",Isbn: "4390",Director: &Director{Firstname: "H", Lastname: "Rembhotkar"}})
r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
r.HandleFunc("/movies",createMovie).Methods("POST")

fmt.Println("Starting server at port 8000/n")
fmt.Println("server is started")


}

//Functions
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID== params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index , item := range movies{

		if item.ID == params["id"]{
			movies= append(movies[:index], movies[index+1:]...)
			var movie Movie
			_= json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}

	}

}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewDecoder(r.Body).Decode(movie)

}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID ==params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}

	}
}