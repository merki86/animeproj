package main

import (
	"fmt"
	"log"
	"net/http"
	"serwennn/studyproject/controllers"
)

func main() {
	fmt.Println("Starting")

	static := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static", static))

	http.HandleFunc("/", controllers.GetAnimes)
	http.HandleFunc("/add-anime", controllers.AddAnime)
	http.HandleFunc("/delete-anime", controllers.DeleteAnime)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
