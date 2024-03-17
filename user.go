package restapi

type User struct {
	Id       int    `json:"-" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"user_role"`
}
