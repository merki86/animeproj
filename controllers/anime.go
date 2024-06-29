package controllers

import (
	"io"
	"net/http"
	"os"
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

	// Parsing the form
	err := req.ParseMultipartForm(10 << 20) // Max file size 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the uploaded image
	file, header, err := req.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Creating the file
	targetFile := "static/images/" + header.Filename
	if _, err := os.Stat(targetFile); !os.IsNotExist(err) {
		http.Error(w, "File already exists.", http.StatusConflict)
		return
	}

	// Saving the file
	dst, err := os.Create(targetFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copying the image into the file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(req.FormValue("id"))
	if err != nil {
		views.Template.ExecuteTemplate(w, "error.gohtml", struct{ Error error }{Error: err})
		return
	}
	title := req.FormValue("title")
	translatedTitle := req.FormValue("translatedTitle")

	domain.AddAnime(id, title, translatedTitle, header.Filename)

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
