package entity

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Status struct {
	Status string `json:"status" db:"status" binding:"required"`
}
