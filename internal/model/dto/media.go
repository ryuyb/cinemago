package dto

type BasicMediaInfo struct {
	TmdbId           int64   `json:"tmdb_id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Overview         string  `json:"overview"`
	Adult            bool    `json:"adult"`
	ReleaseDate      string  `json:"release_date"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int64   `json:"vote_count"`
	Popularity       float32 `json:"popularity"`
	GenreIds         []int64 `json:"genre_ids"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	BackdropPath     string  `json:"backdrop_path"`
}

type MovieDetail struct {
	TmdbId              int64   `json:"tmdb_id"`
	ImdbId              string  `json:"imdb_id"`
	Title               string  `json:"title"`
	VoteAverage         float32 `json:"vote_average"`
	VoteCount           int64   `json:"vote_count"`
	Adult               bool    `json:"adult"`
	BackdropPath        string  `json:"backdrop_path"`
	BelongsToCollection struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		PosterPath   string `json:"poster_path"`
		BackdropPath string `json:"backdrop_path"`
	} `json:"belongs_to_collection"`
	Budget int64 `json:"budget"`
	Genres []struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalTitle       string   `json:"original_title"`
	Overview            string   `json:"overview"`
	Popularity          float32  `json:"popularity"`
	PosterPath          string   `json:"poster_path"`
	ProductionCompanies []struct {
		Id            int64  `json:"id"`
		Name          string `json:"name"`
		LogoPath      string `json:"logo_path"`
		OriginCountry string `json:"origin_country"`
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string `json:"name"`
	} `json:"production_countries"`
	ReleaseDate     string  `json:"release_date"`
	Revenue         float32 `json:"revenue"`
	Runtime         int     `json:"runtime"`
	SpokenLanguages []struct {
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Iso639_1    string `json:"iso_639_1"`
	} `json:"spoken_language"`
	Status string `json:"status"`
	Video  bool   `json:"video"`
}
