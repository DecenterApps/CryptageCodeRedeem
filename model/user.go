package model

type UserInterface interface {
	GetId() int
	GetEmail() string
	GetAddress() string
}

type User struct {
	Id      int
	Email   string
	Address string
}

func (u User) GetId() int {
	return u.Id
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetAddress() string {
	return u.Address
}

func (u User) Validate() error {
	return nil
}