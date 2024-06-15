package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"serwennn/studyproject/domain"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*.gohtml"))
}

func main() {
	fmt.Println("Starting")

	http.HandleFunc("/", index)
	http.HandleFunc("/add-anime", addAnime)
	// http.HandleFunc("/anime/{id}", anime)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	animesMap := make(map[string][]domain.Anime)

	animesMap["Animes"] = domain.Animes

	tpl.ExecuteTemplate(w, "index.gohtml", animesMap)
}

func addAnime(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	title := req.FormValue("title")
	translatedTitle := req.FormValue("translatedTitle")
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		tpl.ExecuteTemplate(w, "error.gohtml", struct{ Error error }{Error: err})
		return
	}

	anime := domain.Anime{
		Title:           title,
		TranslatedTitle: translatedTitle,
		ID:              id,
	}

	domain.Animes = append(domain.Animes, anime)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// func anime(w http.ResponseWriter, req *http.Request) {
// 	idString := req.PathValue("id")
// 	fmt.Fprintf(w, "Anime ID: %v", idString)
// }
