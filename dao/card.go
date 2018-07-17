package dao

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type CardDAO struct{}

func NewCardDAO() *CardDAO {
	return &CardDAO{}
}

func (dao *CardDAO) Get(rs app.RequestScope, id int) (*model.Card, error) {
	var card model.Card
	err := rs.Tx().Select().Model(id, &card)
	return &card, err
}

func (dao *CardDAO) Create(rs app.RequestScope, card *model.Card) error {
	card.Id = 0
	return rs.Tx().Model(card).Insert()
}

func (dao *CardDAO) Update(rs app.RequestScope, id int, card *model.Card) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	card.Id = id
	return rs.Tx().Model(card).Exclude("Id").Update()
}

func (dao *CardDAO) Delete(rs app.RequestScope, id int) error {
	card, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(card).Delete()
}

func (dao *CardDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("card").Row(&count)
	return count, err
}

func (dao *CardDAO) Query(rs app.RequestScope, offset, limit int) ([]model.Card, error) {
	cards := []model.Card{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&cards)
	return cards, err
}
