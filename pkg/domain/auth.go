package domain

type ReponseLogin struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	Id       uint   `json:"id"`
}
