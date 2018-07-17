package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
	"github.com/DecenterApps/CryptageCodeRedeem/dao"
)

type couponDAO interface {
	Get(rs app.RequestScope, id int) (*model.Coupon, error)
	GetByToken(rs app.RequestScope, token string) (*dao.Coupon, error)
	Update(rs app.RequestScope, id int, coupon *model.Coupon) error
}

type CouponService struct {
	dao couponDAO
	cardDao cardDAO
}

func NewCouponService(dao couponDAO, userDao userDAO, cardDao cardDAO) *CouponService {
	return &CouponService{dao, cardDao}
}

func (s *CouponService) Get(rs app.RequestScope, id int) (*model.Coupon, error) {
	return s.dao.Get(rs, id)
}

func (s *CouponService) GetByToken(rs app.RequestScope, token string) (*model.Coupon, error) {
	couponDB, err := s.dao.GetByToken(rs, token)
	if err != nil {
		return nil, err
	}
	coupon := model.Coupon{Id: couponDB.Id, Token: couponDB.Token}

	coupon.Card, err = s.cardDao.Get(rs, couponDB.CardId)
	if err != nil {
		return nil, err
	}

	if couponDB.UserId != nil {
		return nil, err
	}

	return &coupon, nil
}

func (s *CouponService) Update(rs app.RequestScope, id int, model *model.Coupon) (*model.Coupon, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, id)
}