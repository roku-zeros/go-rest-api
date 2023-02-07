package store

import "github.com/roku-zeros/go-rest-api/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
