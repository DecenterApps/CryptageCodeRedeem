package model

type CardInterface interface {
	GetId() int
	GetName() string
}

type Card struct {
	Id      int
	Name   string
}

func (c Card) GetId() int {
	return c.Id
}

func (c Card) GetName() string {
	return c.Name
}

func (c Card) Validate() error {
	return nil
}