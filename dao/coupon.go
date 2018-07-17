package dao

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
	"github.com/go-ozzo/ozzo-dbx"
)

type CouponDAO struct{}

type CouponDB struct {
	Id    int
	Token string
	CardId  int
	UserId  int
}

func NewCouponDAO() *CouponDAO {
	return &CouponDAO{}
}

func (dao *CouponDAO) Get(rs app.RequestScope, id int) (*model.Coupon, error) {
	var coupon model.Coupon
	err := rs.Tx().Select().Model(id, &coupon)
	return &coupon, err
}

func (dao *CouponDAO) GetByToken(rs app.RequestScope, token string) (*CouponDB, error) {
	var coupon CouponDB
	err := rs.Tx().Select().Where(dbx.HashExp{"token": token}).One(&coupon)
	return &coupon, err
}

func (dao *CouponDAO) Create(rs app.RequestScope, coupon *model.Coupon) error {
	coupon.Id = 0
	return rs.Tx().Model(coupon).Insert()
}

func (dao *CouponDAO) Update(rs app.RequestScope, id int, coupon *model.Coupon) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	coupon.Id = id
	return rs.Tx().Model(coupon).Exclude("Id").Update()
}

func (dao *CouponDAO) Delete(rs app.RequestScope, id int) error {
	coupon, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(coupon).Delete()
}

func (dao *CouponDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("coupon").Row(&count)
	return count, err
}

func (dao *CouponDAO) Query(rs app.RequestScope, offset, limit int) ([]model.Coupon, error) {
	coupons := []model.Coupon{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&coupons)
	return coupons, err
}
