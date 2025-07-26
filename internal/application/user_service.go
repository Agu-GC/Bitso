package application

type UserServiceInterface interface {
	GetUserInfo(id string) User
}

type UserService struct {
}

func (u *UserService) GetUserInfo(id string) User {

	if id == "" {
		return User{}
	}

	return User{
		Id:   id,
		Name: "Pablo",
	}
}
