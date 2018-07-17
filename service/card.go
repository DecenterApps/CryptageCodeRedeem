package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type cardDAO interface {
	// Get returns the card with the specified card ID.
	Get(rs app.RequestScope, id int) (*model.Card, error)
	// Count returns the number of cards.
	Count(rs app.RequestScope) (int, error)
	// Query returns the list of cards with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]model.Card, error)
	// Create saves a new card in the storage.
	Create(rs app.RequestScope, card *model.Card) error
	// Update updates the card with given ID in the storage.
	Update(rs app.RequestScope, id int, card *model.Card) error
	// Delete removes the card with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
}

type CardService struct {
	dao cardDAO
}

func NewCardService(dao cardDAO) *CardService {
	return &CardService{dao}
}

func (s *CardService) Get(rs app.RequestScope, id int) (*model.Card, error) {
	return s.dao.Get(rs, id)
}

func (s *CardService) Create(rs app.RequestScope, model *model.Card) (*model.Card, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

func (s *CardService) Update(rs app.RequestScope, id int, model *model.Card) (*model.Card, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}

func (s *CardService) Delete(rs app.RequestScope, id int) (*model.Card, error) {
	card, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return card, err
}

func (s *CardService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}

func (s *CardService) Query(rs app.RequestScope, offset, limit int) ([]model.Card, error) {
	return s.dao.Query(rs, offset, limit)
}