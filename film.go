package restapi

type Film struct {
	FilmId      int    `json:"film_id"`
	Title       string `json:"title" binding:"required"`
	Date        string `json:"date" binding:"required"`
	Description string `json:"description" binding:"required"`
	Rate        string `json:"rate" binding:"required"`
	Actors      []Actor
}

type ToChangeFilm struct {
	Film
	ToChangeTitle       string `json:"to_change_title"`
	ToChangeDate        string `json:"to_change_date"`
	ToChangeDescription string `json:"to_change_description""`
	ToChangeRate        string `json:"to_change_rate"`
}
