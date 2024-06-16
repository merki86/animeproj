package controllers

import (
	"net/http"
	"serwennn/studyproject/domain"
	"serwennn/studyproject/views"
	"strconv"
)

func GetAnimes(w http.ResponseWriter, req *http.Request) {
	views.Template.ExecuteTemplate(w, "index.gohtml", domain.GetAnimes())
}

func AddAnime(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	title := req.FormValue("title")
	translatedTitle := req.FormValue("translatedTitle")
	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		views.Template.ExecuteTemplate(w, "error.gohtml", struct{ Error error }{Error: err})
		return
	}

	domain.AddAnime(title, translatedTitle, id)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func DeleteAnime(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		views.Template.ExecuteTemplate(w, "error.gohtml", struct{ Error error }{Error: err})
		return
	}

	domain.DeleteAnimeById(id)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
