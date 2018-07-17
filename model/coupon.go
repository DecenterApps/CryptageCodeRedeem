package model

type (
	CouponInterface interface {
		GetId() int
		GetCardId() int
		GetToken() string
		GetEmail() *string
		GetAddress() *string
	}

	Coupon struct {
		Id      int     `json:"id" db:"id"`
		CardId  int     `json:"card" db:"card_id"`
		Token   string  `json:"token" db:"token"`
		Email   *string `json:"email" db:"email"`
		Address *string `json:"address" db:"address"`
	}
)

func (c Coupon) GetId() int {
	return c.Id
}

func (c Coupon) GetToken() string {
	return c.Token
}

func (c Coupon) GetCardId() int {
	return c.CardId
}

func (c Coupon) GetEmail() *string {
	return c.Email
}

func (c Coupon) GetAddress() *string {
	return c.Address
}

func (c Coupon) Validate() error {
	return nil
}
