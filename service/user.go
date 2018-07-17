package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type userDAO interface {
	Get(rs app.RequestScope, id int) (*model.User, error)
	Create(rs app.RequestScope, user *model.User) error
}

type UserService struct {
	dao userDAO
}

func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

func (s *UserService) Get(rs app.RequestScope, id int) (*model.User, error) {
	return s.dao.Get(rs, id)
}

func (s *UserService) Create(rs app.RequestScope, model *model.User) (*model.User, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}