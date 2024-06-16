package main

import (
	"fmt"
	"log"
	"net/http"
	"serwennn/studyproject/controllers"
)

func main() {
	fmt.Println("Starting")

	http.HandleFunc("/", controllers.GetAnimes)
	http.HandleFunc("/add-anime", controllers.AddAnime)
	http.HandleFunc("/delete-anime", controllers.DeleteAnime)
	// http.HandleFunc("/anime/{id}", anime)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func anime(w http.ResponseWriter, req *http.Request) {
// 	idString := req.PathValue("id")
// 	fmt.Fprintf(w, "Anime ID: %v", idString)
// }
