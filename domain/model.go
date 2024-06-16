package domain

type Anime struct {
	Title           string
	TranslatedTitle string
	ID              int
}

var Animes []Anime

func GetAnimes() map[string][]Anime {
	animesMap := make(map[string][]Anime)
	animesMap["Animes"] = Animes

	return animesMap
}

func AddAnime(title string, translatedTitle string, id int) {
	anime := Anime{
		Title:           title,
		TranslatedTitle: translatedTitle,
		ID:              id,
	}

	Animes = append(Animes, anime)
}

func DeleteAnimeById(id int) {
	for index, anime := range Animes {
		if anime.ID == id {
			Animes = append(Animes[:index], Animes[index+1:]...)
		}
	}
}
