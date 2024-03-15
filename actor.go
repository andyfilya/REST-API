package restapi

type Actor struct {
	ActorId   int
	FirstName string
	LastName  string
	DateBirth string
	Films     []Film
}
