package storage

type IStorage interface {
	FindACL(email string) (string, bool)
	ChangeACL(email, access string) error
}
