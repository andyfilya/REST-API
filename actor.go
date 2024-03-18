package restapi

type Actor struct {
	ActorId   int    `db:"actor_id"`
	FirstName string `json:"name" db:"actor_name" binding:"required"`
	LastName  string `json:"surname" db:"actor_surname" binding:"required"`
	DateBirth string `json:"date_birth" db:"actor_birth_date" binding:"required"'`
	Films     []Film
}

// ToChange struct use for update old fields of struct for new (TO_CHANGE_USERNAME,...)
type ToChange struct {
	Actor
	ToChangeUsername string `json:"to_change_name"`
	ToChangeSurname  string `json:"to_change_surname"`
	ToChangeBirth    string `json:"to_change_birth"'`
}

type ActorFragment struct {
	ActorNameFragment    string `json:"name"`
	ActorSurnameFragment string `json:"surname"`
}
