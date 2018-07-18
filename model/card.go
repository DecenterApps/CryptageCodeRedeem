package model

import (
	"fmt"
	"strconv"
)

type (
	Card struct {
		Id        uint   `json:"id" db:"id"`
		Flavor    string `json:"flavor" db:"flavor"`
		Image     string `json:"image" db:"image"`
		Mechanics string `json:"mechanics" db:"mechanics"`
		Rarity    string `json:"rarity" db:"rarity"`
		Title     string `json:"title" db:"title"`
		Type      string `json:"type" db:"type"`
		Level     uint   `json:"level" db:"level"`
		Cost      Cost   `json:"cost" db:"cost"`
		Values    Values `json:"values" db:"values"`
	}

	Cost struct {
		Development    uint `json:"development" db:"development"`
		Funds          uint `json:"funds" db:"funds"`
		Level          uint `json:"level" db:"level"`
		Power          uint `json:"power" db:"power"`
		Space          uint `json:"space" db:"space"`
		ContainerSpace uint `json:"containerSpace" db:"containerSpace"`
		Time           uint `json:"time" db:"time"`
	}

	Values struct {
		Development     uint `json:"development" db:"development"`
		Funds           uint `json:"funds" db:"funds"`
		FundsPerBlock   uint `json:"fundsPerBlock" db:"fundsPerBlock"`
		ContainerSpace  uint `json:"containerSpace" db:"containerSpace"`
		MultiplierDev   uint `json:"multiplierDev" db:"multiplierDev"`
		MultiplierFunds uint `json:"multiplierFunds" db:"multiplierFunds"`
		MultiplierTime  uint `json:"multiplierTime" db:"multiplierTime"`
		Space           uint `json:"space" db:"space"`
		Power           uint `json:"power" db:"power"`
		Xp              uint `json:"xp" db:"experience"`
	}

	CardInterface interface {
		GetId() int
		GetImage() string
		GetFlavor() string
		GetMechanics() string
		GetRarityColor() string
		GetRarity() string
		GetTitle() string
		GetType() string
		formatNumber(n uint) string
	}
)

var cardColor = map[uint]string{
	42: "#9F00C7",
	16: "#9F00C7",
	17: "#9F00C7",
	38: "#3CC8CC",
	43: "#3215E6",
	32: "#878787",
}

var cardColorByType = map[string]string{
	"Power":     "#CE060D",
	"Misc":      "#3215E5",
	"Location":  "#3CC8CC",
	"Person":    "#9F00C7",
	"Project":   "#878787",
	"Mining":    "#75341F",
	"Container": "#4A7420",
}

func (c Card) GetId() int {
	return int(c.Id)
}

func (c Card) GetImage() string {
	return c.Image
}

func (c Card) GetFlavor() string {
	return c.Flavor
}

func (c Card) GetMechanics() string {
	return c.Mechanics
}

func (c Card) GetRarityColor() string {
	rarity, _ := strconv.ParseInt(c.Rarity, 10, 0)
	if rarity >= 850 {
		return "normal"
	}
	if rarity >= 600 {
		return "blue"
	}
	if rarity >= 325 {
		return "purple"
	}
	return "gold"
}

func (c Card) GetRarityLevel() string {
	rarity, _ := strconv.ParseInt(c.Rarity, 10, 0)
	if rarity >= 850 {
		return "Regular"
	}
	if rarity >= 600 {
		return "Scarce"
	}
	if rarity >= 325 {
		return "Rare"
	}
	return "Elite"
}

func (c Card) GetTitle() string {
	return c.Title
}

func (c Card) GetType() string {
	return c.Type
}

func (c Card) GetColor() string {
	return cardColorByType[c.Type]
}

func (c Card) Validate() error {
	return nil
}

func (c Card) FormatNumber(n uint) string {
	if n >= 1000000 {
		return fmt.Sprintf("%vm", n/1000000)
	}
	if n >= 1000 {
		return fmt.Sprintf("%vk", n/1000)
	}
	return fmt.Sprintf("%v", n)
}
