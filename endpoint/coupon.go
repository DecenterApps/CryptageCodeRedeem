package endpoint

import (
	"html/template"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type (
	couponService interface {
		Get(rs app.RequestScope, id int) (*model.Coupon, error)
		GetByToken(rs app.RequestScope, token string) (*model.Coupon, error)
		Update(rs app.RequestScope, id int, model *model.Coupon) (*model.Coupon, error)
	}

	cardService interface {
		Get(rs app.RequestScope, id int) (*model.Card, error)
	}

	couponResource struct {
		couponService couponService
		cardService   cardService
	}

	response struct {
		Coupon *model.Coupon
		Card   *model.Card
		Error  string
		Done   bool
	}
)

func ServeCouponResource(rg *routing.RouteGroup, couponService couponService, cardService cardService) {
	r := &couponResource{couponService, cardService}
	rg.Get("/freeCards/<token>", r.get)
	rg.Post("/freeCards/<token>", r.update)
}

func (r *couponResource) get(c *routing.Context) error {
	tmpl := template.Must(template.ParseFiles("./front/src/index.html"))
	coupon, err := r.couponService.GetByToken(app.GetRequestScope(c), c.Param("token"))
	if err != nil {
		return tmpl.Execute(c.Response, response{Error: "invalid"})
	}

	if coupon.Email != nil && coupon.Address != nil {
		return tmpl.Execute(c.Response, response{Error: "used"})
	}

	card, _ := r.cardService.Get(app.GetRequestScope(c), coupon.Id)
	return tmpl.Execute(c.Response, response{Coupon: coupon, Card: card})
}

func (r *couponResource) update(c *routing.Context) error {
	rs := app.GetRequestScope(c)

	tmpl := template.Must(template.ParseFiles("./front/src/confirmation.html"))
	coupon, err := r.couponService.GetByToken(rs, c.Param("token"))
	if err != nil {
		return tmpl.Execute(c.Response, response{Error: "invalid"})
	}

	if coupon.Email != nil && coupon.Address != nil {
		return tmpl.Execute(c.Response, response{Error: "used"})
	}

	coupon.Email = new(string)
	coupon.Address = new(string)
	*coupon.Email = c.PostForm("email")
	*coupon.Address = c.PostForm("address")

	_, err = r.couponService.Update(rs, coupon.Id, coupon)
	if err != nil {
		return tmpl.Execute(c.Response, response{Error: "error"})
	}

	return tmpl.Execute(c.Response, response{Done: true})
}
