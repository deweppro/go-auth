package auth

type IUser interface {
	GetEmail() string
	SetACL(string)
}
