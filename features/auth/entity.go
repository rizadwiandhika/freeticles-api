package auth

type UserCore struct {
	ID       uint
	Username string
	Name     string
	Email    string
	Password string
}

type IBusiness interface {
	Authenticate(user UserCore) (string, error, int)
	Register(user UserCore) (UserCore, error, int)
}
