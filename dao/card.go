package dao

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type CardDAO struct {
	cards map[uint]map[uint]model.Card
}

func NewCardDAO(cryptage app.Cryptage) *CardDAO {
	return &CardDAO{cryptage.Cards}
}

func (dao *CardDAO) Get(rs app.RequestScope, id int) (model.Card, error) {
	return dao.cards[uint(id)][1], nil
}