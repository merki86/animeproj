package domain

import "fmt"

type Anime struct {
	ID              int
	Title           string
	TranslatedTitle string
	Image           string
}

var Animes []Anime

func GetAnimes() map[string][]Anime {
	animesMap := make(map[string][]Anime)
	animesMap["Animes"] = Animes

	return animesMap
}

func AddAnime(id int, title string, translatedTitle string, image string) {
	anime := Anime{
		ID:              id,
		Title:           title,
		TranslatedTitle: translatedTitle,
		Image:           image,
	}

	fmt.Println(anime.Image)

	Animes = append(Animes, anime)
}

func DeleteAnimeById(id int) {
	for index, anime := range Animes {
		if anime.ID == id {
			Animes = append(Animes[:index], Animes[index+1:]...)
		}
	}
}
