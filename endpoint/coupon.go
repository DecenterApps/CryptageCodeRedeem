package endpoint

import (
	"html/template"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type (
	// couponService specifies the interface for the coupon service needed by couponResource.
	couponService interface {
		Get(rs app.RequestScope, id int) (*model.Coupon, error)
		GetByToken(rs app.RequestScope, token string) (*model.Coupon, error)
		Update(rs app.RequestScope, id int, model *model.Coupon) (*model.Coupon, error)
	}

	userService interface {
		Get(rs app.RequestScope, id int) (*model.User, error)
		Create(rs app.RequestScope, model *model.User) (*model.User, error)
	}

	couponResource struct {
		service couponService
		userService userService
	}

	response struct {
		Coupon *model.Coupon
		Error string
	}
)

func ServeCouponResource(rg *routing.RouteGroup, service couponService, userService userService) {
	r := &couponResource{service, userService}
	rg.Get("/freeCards/<token>", r.get)
	rg.Post("/freeCards/<token>", r.update)
}

func (r *couponResource) get(c *routing.Context) error {
	tmpl := template.Must(template.ParseFiles("./front/src/index.html"))
	coupon, err := r.service.GetByToken(app.GetRequestScope(c), c.Param("token"))
	if err != nil {
		return tmpl.Execute(c.Response, response{Error: "invalid"})
	}

	if coupon.User != nil {
		return tmpl.Execute(c.Response, response{Error: "already-used"})
	}

	return tmpl.Execute(c.Response, response{Coupon: coupon})
}

func (r *couponResource) update(c *routing.Context) error {
	rs := app.GetRequestScope(c)

	coupon, err := r.service.GetByToken(rs, c.Param("token"))

	if coupon.User != nil {
		return c.Write("Coupon already used")
	}

	if err != nil {
		return err
	}

	user := model.User{Email: c.Param("email"), Address: c.Param("address")}
	r.userService.Create(rs, &user)
	coupon.User = &user

	response, err := r.service.Update(rs, coupon.Id, coupon)
	if err != nil {
		return err
	}

	return c.Write(response)
}