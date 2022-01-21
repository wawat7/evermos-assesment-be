package user

import (
	"errors"
)

type Service interface {
	Save(user User) User
	FindAll() (users []User)
	FindById(Id int) (user User, err error)
	Delete(user User)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Save(user User) User {
	user = s.repository.Create(user)
	return user
}

func (s *service) FindAll() (users []User) {
	users = s.repository.FindAll()
	return
}

func (s *service) FindById(Id int) (user User, err error) {
	user = s.repository.FindById(Id)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *service) Delete(user User) {
	s.repository.Delete(user)

	return
}
