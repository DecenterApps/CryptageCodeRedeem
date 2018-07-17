package dao

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type CouponDAO struct{}

func NewCouponDAO() *CouponDAO {
	return &CouponDAO{}
}

func (dao *CouponDAO) Get(rs app.RequestScope, id int) (*model.Coupon, error) {
	var coupon model.Coupon
	err := rs.Tx().Select().Model(id, &coupon)
	return &coupon, err
}

func (dao *CouponDAO) GetByToken(rs app.RequestScope, token string) (*model.Coupon, error) {
	var coupon model.Coupon
	err := rs.Tx().Select().Where(dbx.HashExp{"token": token}).One(&coupon)
	return &coupon, err
}

func (dao *CouponDAO) Update(rs app.RequestScope, id int, coupon *model.Coupon) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	coupon.Id = id
	return rs.Tx().Model(coupon).Exclude("Id").Update()
}