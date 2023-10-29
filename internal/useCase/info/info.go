package info

import (
	"requestLimiter/adapters"
)

type UseCase struct {
	repo adapters.Info
}

func NewUseCase(infoRepo adapters.Info) *UseCase {
	return &UseCase{repo: infoRepo}
}

func (u *UseCase) GetInfo(userID string) string {
	return u.repo.GetInfoByUserID(userID)
}
