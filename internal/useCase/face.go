package useCase

type Info interface {
	GetInfo(userID string) string
}
