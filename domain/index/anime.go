package domain

type Anime struct {
	Title           string
	TranslatedTitle string
	ID              int
}

var Animes []Anime

func GetAnimeList() map[string][]Anime {
	animesMap := make(map[string][]Anime)
	animesMap["Animes"] = Animes

	return animesMap
}
