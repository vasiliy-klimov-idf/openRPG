package items

import (
	c "OpenRPG/src/backend/colors"
)

type Item struct {
	Name             string
	Have             bool
	Outfitted        bool
	Rarity           string
	ItemRequirements ItemRequirements
}

type Rarity struct {
	Name        string
	Color       string
	Probability float64
}

type ItemRequirements struct {
	Strength       int `json:"strength"`
	Dexterity      int `json:"dexterity"`
	BodyDifficulty int `json:"body_difficulty"`
	Intelligence   int `json:"intelligence"`
	Wisdom         int `json:"wisdom"`
	Charisma       int `json:"charisma"`
}

var (
	Common    = Rarity{Name: "Common", Color: c.White, Probability: 1.00}
	Uncommon  = Rarity{Name: "Uncommon", Color: c.Green, Probability: 0.80}
	Rare      = Rarity{Name: "Rare", Color: c.Blue, Probability: 0.60}
	Epic      = Rarity{Name: "Epic", Color: c.Purple, Probability: 0.40}
	Legendary = Rarity{Name: "Legendary", Color: c.Yellow, Probability: 0.20}
	Mythic    = Rarity{Name: "Mythic", Color: c.Red, Probability: 0.10}
)
