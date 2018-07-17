package model

type CardInterface interface {
	GetId() int
	GetImage() string
	GetFlavorText() string
	GetMechanicsText() string
	GetRarityScore() string
	GetTitle() string
	GetType() string
}

type Card struct {
	Id            int    `json:"id"`
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

type Costs struct {
	Development    uint `json:"development"`
	Funds          uint `json:"funds"`
	Level          uint `json:"level"`
	Power          uint `json:"power"`
	Space          uint `json:"space"`
	ContainerSpace uint `json:"containerSpace"`
	Time           uint `json:"time"`
}

type Gains struct {
	Development     uint `json:"development"`
	Funds           uint `json:"funds"`
	MultiplierDev   uint `json:"multiplierDev"`
	MultiplierFunds uint `json:"multiplierFunds"`
	MultiplierTime  uint `json:"multiplierTime"`
	Power           uint `json:"power"`
	Space           uint `json:"space"`
	ContainerSpace  uint `json:"containerSpace"`
	Xp              uint `json:"xp"`
}

func (c Card) GetId() int {
	return c.Id
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

func (c Card) GetRarityScore() string {
	return c.RarityScore
}

func (c Card) GetTitle() string {
	return c.Title
}

func (c Card) Validate() error {
	return nil
}
