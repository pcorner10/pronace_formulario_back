package domain

type ReponseLogin struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
	Id       uint   `json:"id"`
}

type AuthDB interface {
	Login(email, password string) (*ReponseLogin, error)
	Register(user *User) error
}

type AuthService interface {
	Login(email, password string) (*ReponseLogin, error)
	Register(user *User) error
}
