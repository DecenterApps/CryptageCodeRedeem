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

func (dao *UserDAO) Update(rs app.RequestScope, id int, user *model.User) error {
	if _, err := dao.Get(rs, id); err != nil {
		return err
	}
	user.Id = id
	return rs.Tx().Model(user).Exclude("Id").Update()
}

func (dao *UserDAO) Delete(rs app.RequestScope, id int) error {
	user, err := dao.Get(rs, id)
	if err != nil {
		return err
	}
	return rs.Tx().Model(user).Delete()
}

func (dao *UserDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("user").Row(&count)
	return count, err
}

func (dao *UserDAO) Query(rs app.RequestScope, offset, limit int) ([]model.User, error) {
	users := []model.User{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&users)
	return users, err
}
