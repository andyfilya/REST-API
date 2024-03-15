package restapi

type Film struct {
	FilmId      int
	Title       string
	Description string
	URLfilm     string
	Actors      []Actor
}
