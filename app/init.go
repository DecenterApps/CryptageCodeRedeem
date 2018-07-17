package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/access"
	"github.com/go-ozzo/ozzo-routing/fault"
	"github.com/go-ozzo/ozzo-validation"
	cryptageError "github.com/DecenterApps/CryptageCodeRedeem/error"
)

func Init(logger *logrus.Logger) routing.Handler {
	return func(rc *routing.Context) error {
		now := time.Now()

		rc.Response = &access.LogResponseWriter{rc.Response, http.StatusOK, 0}

		ac := newRequestScope(now, logger, rc.Request)
		rc.Set("Context", ac)

		fault.Recovery(ac.Errorf, convertError)(rc)
		logAccess(rc, ac.Infof, ac.Now())

		return nil
	}
}

func GetRequestScope(c *routing.Context) RequestScope {
	return c.Get("Context").(RequestScope)
}

func logAccess(c *routing.Context, logFunc access.LogFunc, start time.Time) {
	rw := c.Response.(*access.LogResponseWriter)
	elapsed := float64(time.Now().Sub(start).Nanoseconds()) / 1e6
	requestLine := fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto)
	logFunc(`[%.3fms] %s %d %d`, elapsed, requestLine, rw.Status, rw.BytesWritten)
}

func convertError(c *routing.Context, err error) error {
	if err == sql.ErrNoRows {
		return cryptageError.NotFound("the requested resource")
	}
	switch err.(type) {
	case *cryptageError.APIError:
		return err
	case validation.Errors:
		return cryptageError.InvalidData(err.(validation.Errors))
	case routing.HTTPError:
		switch err.(routing.HTTPError).StatusCode() {
		case http.StatusUnauthorized:
			return cryptageError.Unauthorized(err.Error())
		case http.StatusNotFound:
			return cryptageError.NotFound("the requested resource")
		}
	}
	return cryptageError.InternalServerError(err)
}
