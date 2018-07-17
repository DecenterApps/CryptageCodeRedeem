package app

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
)

type RequestScope interface {
	Logger
	UserID() string
	SetUserID(id string)
	RequestID() string
	Tx() *dbx.Tx
	SetTx(tx *dbx.Tx)
	Rollback() bool
	SetRollback(bool)
	Now() time.Time
}

type requestScope struct {
	Logger              // the logger tagged with the current request information
	now       time.Time // the time when the request is being processed
	requestID string    // an ID identifying one or multiple correlated HTTP requests
	userID    string    // an ID identifying the current user
	rollback  bool      // whether to roll back the current transaction
	tx        *dbx.Tx   // the currently active transaction
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) Tx() *dbx.Tx {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *dbx.Tx) {
	rs.tx = tx
}

func (rs *requestScope) Rollback() bool {
	return rs.rollback
}

func (rs *requestScope) SetRollback(v bool) {
	rs.rollback = v
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
