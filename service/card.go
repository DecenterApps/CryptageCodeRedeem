package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type cardDAO interface {
	Get(rs app.RequestScope, id int) (model.Card, error)
}

type CardService struct {
	dao cardDAO
}

func (s *CardService) Get(rs app.RequestScope, id int) (model.Card, error) {
	return s.dao.Get(rs, id)
}