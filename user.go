package auth

type IUser interface {
	GetName() string
	GetEmail() string
	GetIcon() string
	GetACL() string
	SetACL(string)
}
