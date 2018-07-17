package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type userDAO interface {
	// Get returns the user with the specified user ID.
	Get(rs app.RequestScope, id int) (*model.User, error)
	// Count returns the number of users.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of users with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]model.User, error)
	// Create saves a new user in the storage.
	Create(rs app.RequestScope, user *model.User) error
	// Update updates the user with given ID in the storage.
	Update(rs app.RequestScope, id int, user *model.User) error
	// Delete removes the user with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
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

func (s *UserService) Update(rs app.RequestScope, id int, model *model.User) (*model.User, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *UserService) Delete(rs app.RequestScope, id int) (*model.User, error) {
	user, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return user, err
}

func (s *UserService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *UserService) Query(rs app.RequestScope, offset, limit int) ([]model.User, error) {
	return s.dao.Query(rs, offset, limit)
}