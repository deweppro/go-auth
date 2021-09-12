package storage

type IStorage interface {
	FindACL(string) (string, bool)
}
