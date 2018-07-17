package service

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type (
	couponDAO interface {
		Get(rs app.RequestScope, id int) (*model.Coupon, error)
		GetByToken(rs app.RequestScope, token string) (*model.Coupon, error)
		Update(rs app.RequestScope, id int, coupon *model.Coupon) error
	}

	CouponService struct {
		dao couponDAO
	}
)

func NewCouponService(dao couponDAO) *CouponService {
	return &CouponService{dao}
}

func (s *CouponService) Get(rs app.RequestScope, id int) (*model.Coupon, error) {
	couponDB, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}

	return &model.Coupon{Id: couponDB.Id, Token: couponDB.Token}, nil
}

func (s *CouponService) GetByToken(rs app.RequestScope, token string) (*model.Coupon, error) {
	coupon, err := s.dao.GetByToken(rs, token)
	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (s *CouponService) Update(rs app.RequestScope, id int, model *model.Coupon) (*model.Coupon, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Update(rs, id, model); err != nil {
		return nil, err
	}
	return s.Get(rs, id)
}
