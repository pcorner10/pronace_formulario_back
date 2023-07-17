package domain

type ReponseLogin struct {
	Token string `json:"token"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Id    uint   `json:"id"`
}
