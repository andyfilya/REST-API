package restapi

type Actor struct {
	ActorId   int
	FirstName string `json:"name" binding:"required"`
	LastName  string `json:"surname" binding:"required"`
	DateBirth string `json:"date_birth" binding:"required"'`
	Films     []Film
}
