package model

type Movie struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year"`
	Genre       string `json:"genre"`
	Plot        string `json:"plot"`
	Actors      string `json:"actors"`
}

type MovieResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseYear int    `json:"release_year"`
	Genre       string `json:"genre"`
	Plot        string `json:"plot"`
	Actors      string `json:"actors"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
