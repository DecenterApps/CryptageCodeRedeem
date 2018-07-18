package model

import (
	"fmt"
	"strconv"
)

type (
	Card struct {
		Id            uint   `json:"id"`
		FlavorText    string `json:"flavorText"`
		Image         string `json:"image"`
		MechanicsText string `json:"mechanicsText"`
		RarityScore   string `json:"rarityScore"`
		Title         string `json:"title"`
		Type          string `json:"type"`
		Subtype       string `json:"subtype"`
		Level         uint   `json:"level"`
		Costs         Costs  `json:"cost"`
		Gains         Gains  `json:"values"`
	}

	Costs struct {
		Development    uint `json:"development"`
		Funds          uint `json:"funds"`
		Level          uint `json:"level"`
		Power          uint `json:"power"`
		Space          uint `json:"space"`
		ContainerSpace uint `json:"containerSpace"`
		Time           uint `json:"time"`
	}

	Gains struct {
		Development     uint `json:"development"`
		Funds           uint `json:"funds"`
		FundsPerBlock   uint `json:"fundsPerBlock"`
		MultiplierDev   uint `json:"multiplierDev"`
		MultiplierFunds uint `json:"multiplierFunds"`
		MultiplierTime  uint `json:"multiplierTime"`
		Power           uint `json:"power"`
		Space           uint `json:"space"`
		ContainerSpace  uint `json:"containerSpace"`
		Xp              uint `json:"xp"`
	}

	CardInterface interface {
		GetId() int
		GetImage() string
		GetFlavorText() string
		GetMechanicsText() string
		GetRarityColor() string
		GetRarityLevel() string
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
	"Power": "#CE060D",
	"Misc": "#3215E5",
	"Location": "#3CC8CC",
	"Person": "#9F00C7",
	"Project": "#878787",
	"Mining": "#75341F",
	"Container": "#4A7420",
}

func (c Card) GetId() int {
	return int(c.Id)
}

func (c Card) GetImage() string {
	return c.Image
}

func (c Card) GetFlavorText() string {
	return c.FlavorText
}

func (c Card) GetMechanicsText() string {
	return c.MechanicsText
}

func (c Card) GetRarityColor() string {
	rarity, _ := strconv.ParseInt(c.RarityScore, 10, 0)
	if rarity >= 850 { return "normal" }
	if rarity >= 600 { return "blue"}
	if rarity >= 325 { return "purple" }
	return "gold"
}

func (c Card) GetRarityLevel() string {
	rarity, _ := strconv.ParseInt(c.RarityScore, 10, 0)
	if rarity >= 850 { return "Regular" }
	if rarity >= 600 { return "Scarce" }
	if rarity >= 325 { return "Rare" }
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
	if n >= 1000000 { return fmt.Sprintf("%vm", n/1000000) }
	if n >= 1000 { return fmt.Sprintf("%vk", n/1000) }
	return fmt.Sprintf("%v", n)
}
