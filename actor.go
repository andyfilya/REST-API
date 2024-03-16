package restapi

type Actor struct {
	ActorId   int
	FirstName string `json:"name" binding:"required"`
	LastName  string `json:"surname" binding:"required"`
	DateBirth string `json:"date_birth" binding:"required"'`
	Films     []Film
}

// ToChange struct use for update old fields of struct for new (TO_CHANGE_USERNAME,...)
type ToChange struct {
	Actor
	ToChangeUsername string `json:"to_change_name"`
	ToChangeSurname  string `json:"to_change_surname"`
	ToChangeBirth    string `json:"to_change_birth"'`
}
