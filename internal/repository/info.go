package repository

import (
	"fmt"
	"time"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetInfoByUserID(userID string) string {
	return fmt.Sprintf("userID: %s, called at: %s", userID, time.Now().String())
}
