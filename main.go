package main

//libraries
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// making crud for the movies data without using databases , using slice and struct
// struct movies
type movies struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *director `json:"director"` //associated struct director
}
type director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//slice variable
var movie []movies 

// Handler Function
//for all movies 
func GetAllMovie(w http.ResponseWriter, r *http.Request) {\
	//setting the header content type 
	w.Header().Set("Content-Type", "application/json")
	//encoder -> stream(slice or array ) data into json format 
	json.NewEncoder(w).Encode(movie)

}

//to get movie by id 
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//to get the input request 
	params := mux.Vars(r)
	for _, value := range movie {
		if value.ID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
}

//to create a new movie 
func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newmovie movies
	//decoder -> convert json data into stream 
	_ = json.NewDecoder(r.Body).Decode(&newmovie)
	//creating random id 
	newmovie.ID = strconv.Itoa(rand.Intn(100000))
	//append into the slice movie 
	movie = append(movie, newmovie)
	//converting stream into json 
	json.NewEncoder(w).Encode(newmovie)
}


//to update data by id 
func updateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//delete the data 
	//create the new data 
	parmas := mux.Vars(r)
	for index, value := range movie {
		if value.ID == parmas["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var newmovie movies
			_ = json.NewDecoder(r.Body).Decode(&newmovie)
			newmovie.ID = strconv.Itoa(rand.Intn(100000))
			movie = append(movie, newmovie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}

//to delete the data by id 
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, value := range movie {
		if value.ID == params["id"] {
			//this append function deleted the value at index and join index+1 at that position 
			movie = append(movie[:index], movie[index+1:]...)
			break

		}
	}
	fmt.Fprintf(w, "Id successfully deleted")
	json.NewEncoder(w).Encode(movie)

}

func main() {
	//append values in slice movie
	movie = append(movie, movies{ID: "1", ISBN: "4567899", Title: "Article 370", Director: &director{FirstName: "John", LastName: "Doe"}})
	movie = append(movie, movies{ID: "2", ISBN: "4567889", Title: "Bhootnath", Director: &director{FirstName: "Rohit", LastName: "Sheety"}})

	//instance of router
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetAllMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", GetHandler).Methods("GET")
	router.HandleFunc("/movies", createHandler).Methods("POST")
	router.HandleFunc("/movies/{id}", updateHandler).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteHandler).Methods("DELETE")

	fmt.Println("Starting the CRUD API")
	fmt.Println("Starting the server at 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
