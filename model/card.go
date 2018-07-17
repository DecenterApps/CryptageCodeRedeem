package model

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
		GetRarityScore() string
		GetTitle() string
		GetType() string
	}
)

var cardColor = map[uint]string{
	42: "#9F00C7",
	16: "#9F00C7",
	17: "9F00C7",
	38: "#3CC8CC",
	43: "#3215E6",
	32: "#878787",
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

func (c Card) GetRarityScore() string {
	return c.RarityScore
}

func (c Card) GetTitle() string {
	return c.Title
}

func (c Card) GetType() string {
	return c.Type
}

func (c Card) GetColor() string {
	return cardColor[c.Id]
}

func (c Card) Validate() error {
	return nil
}
