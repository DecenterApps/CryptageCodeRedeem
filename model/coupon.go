package model

type CouponInterface interface {
	GetId() int
	GetToken() string
	GetCard() Card
	GetUser() *User
}

type Coupon struct {
	Id    int
	Token string
	Card  Card
	User  *User
}

func (c Coupon) GetId() int {
	return c.Id
}

func (c Coupon) GetToken() string {
	return c.Token
}

func (c Coupon) GetCard() Card {
	return c.Card
}

func (c Coupon) GetUser() *User {
	return c.User
}

func (c Coupon) Validate() error {
	return nil
}