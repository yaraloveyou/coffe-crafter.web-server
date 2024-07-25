package store

import "github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
}
