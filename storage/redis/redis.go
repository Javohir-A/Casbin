package repository

import (
	"github.com/redis/go-redis/v9"
)

type UserRepo interface {
	SetKeyValue(key string, value string, expiration int) error
	GetValue(key string) (string, error)
	DeleteKey(key string) error
	GetKeys() ([]string, error)
}

type userRepoImpl struct {
	client *redis.Client
}

func NewUserRepo(client *redis.Client) *userRepoImpl {
	return &userRepoImpl{client: client}
}

func (u *userRepoImpl) SetKeyValue(key string, value string, expiration int) error {
	return nil
}
func (u *userRepoImpl) DeleteKey(key string) error {
	return nil
}

func (u *userRepoImpl) GetKey(key string) (string, error) {
	return "", nil
}

func (u *userRepoImpl) GetKeys() ([]string, error) {
	return nil, nil
}
