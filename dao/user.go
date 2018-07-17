package dao

import (
	"github.com/DecenterApps/CryptageCodeRedeem/app"
	"github.com/DecenterApps/CryptageCodeRedeem/model"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(rs app.RequestScope, id int) (*model.User, error) {
	var user model.User
	err := rs.Tx().Select().Model(id, &user)
	return &user, err
}

func (dao *UserDAO) Create(rs app.RequestScope, user *model.User) error {
	user.Id = 0
	return rs.Tx().Model(user).Insert()
}