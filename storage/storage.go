package storage

type IStorage interface {
	FindACL(email string) (string, error)
	ChangeACL(email, access string) error
}
