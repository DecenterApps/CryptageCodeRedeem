package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type (
	cardDAO interface {
		Get(rs app.RequestScope, id int) (*model.Card, error)
	}

	CardService struct {
		dao cardDAO
	}
)

func NewCardService(dao cardDAO) *CardService {
	return &CardService{dao}
}

func (s *CardService) Get(rs app.RequestScope, id int) (*model.Card, error) {
	return s.dao.Get(rs, id)
}
