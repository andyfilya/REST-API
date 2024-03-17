package restapi

type Film struct {
	FilmId      int    `json:"film_id" db:"film_id"`
	Title       string `json:"title" db:"film_title" binding:"required"`
	Date        string `json:"date" db:"film_date" binding:"required"`
	Description string `json:"description" db:"film_description" binding:"required"`
	Rate        string `json:"rate" db:"film_rate" binding:"required"`
	Actors      []Actor
}

type ToChangeFilm struct {
	Film
	ToChangeTitle       string `json:"to_change_title"`
	ToChangeDate        string `json:"to_change_date"`
	ToChangeDescription string `json:"to_change_description""`
	ToChangeRate        string `json:"to_change_rate"`
}
