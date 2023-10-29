package adapters

type Info interface {
	GetInfoByUserID(userID string) string
}
