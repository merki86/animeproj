package controllers

import (
	"fmt"
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

	// Разбор формы
	err := req.ParseMultipartForm(10 << 20) // Максимальный размер файла 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получение файла из формы
	file, header, err := req.FormFile("image") // Имя поля формы должно соответствовать имени input в форме
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Создание целевого файла для сохранения
	targetFile := "uploads/" + header.Filename
	if _, err := os.Stat(targetFile); !os.IsNotExist(err) {
		http.Error(w, "File already exists.", http.StatusConflict)
		return
	}

	// Сохранение файла
	dst, err := os.Create(targetFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копирование содержимого файла в целевой файл
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File '%s' uploaded successfully.\n", header.Filename)

	// id, err := strconv.Atoi(req.FormValue("id"))
	// if err != nil {
	// 	views.Template.ExecuteTemplate(w, "error.gohtml", struct{ Error error }{Error: err})
	// 	return
	// }
	// title := req.FormValue("title")
	// translatedTitle := req.FormValue("translatedTitle")

	domain.AddAnime(1, "title", "translatedTitle", header.Filename)

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
