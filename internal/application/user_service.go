package application

import "github.com/Agu-GC/Bitso/internal/domain"

type UserServiceInterface interface {
	GetUserInfo(id, surname string) domain.User
}

type UserService struct {
}

func (u *UserService) GetUserInfo(id, surname string) domain.User {

	if id == "" {
		return domain.User{}
	}

	return domain.User{
		Id:   id,
		Name: surname,
	}
}
