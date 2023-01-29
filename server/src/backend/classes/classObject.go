package classes

import (
	"OpenRPG/src/backend/items"
	r "OpenRPG/src/backend/races"
)

type ClassJson struct {
	Classes []Class `json:"classes"`
}
type Class struct {
	ClassName      string    `json:"class_name"`
	Race           string    `json:"race"`
	Level          int       `json:"level"`
	Strength       int       `json:"strength"`
	Dexterity      int       `json:"dexterity"`
	BodyDifficulty int       `json:"body_difficulty"`
	Intelligence   int       `json:"intelligence"`
	Wisdom         int       `json:"wisdom"`
	Charisma       int       `json:"charisma"`
	BonusPoints    int       `json:"bonus_points"`
	Points         points    `json:"points"`
	Inventory      Inventory `json:"inventory"`
	Equip          Equip     `json:"eqip"`
}

type Inventory struct {
	Items []items.Item
}

type Equip struct {
	Items []items.Item
}

type points struct {
	HP  int `json:"hp"`
	MP  int `json:"mp"`
	SP  int `json:"sp"`
	EXP int `json:"exp"`
}

//start point for classes
var (
	Mage = Class{ClassName: "Mage", Race: r.Human.Name, Level: 4,
		Strength:       10,
		Dexterity:      8,
		BodyDifficulty: 13,
		Intelligence:   15,
		Wisdom:         14,
		Charisma:       12,
	}

	Knight  = Class{ClassName: "Knight"}
	Thief   = Class{ClassName: "Thief"}
	Warrior = Class{ClassName: "Warrior"}
	Custom  = Class{ClassName: "Custom", BonusPoints: 24}
	Lucky   = Class{ClassName: "Lucky", Race: r.RandomRace().Name, Level: 1,
		Strength:       RandomRollPoints(),
		Dexterity:      RandomRollPoints(),
		BodyDifficulty: RandomRollPoints(),
		Intelligence:   RandomRollPoints(),
		Wisdom:         RandomRollPoints(),
		Charisma:       RandomRollPoints(),
	}

	ClassArray = []Class{Mage, Knight, Thief, Warrior, Custom, Lucky}
)

func (c *Class) ModifyRace(value string) {
	c.Race = value
}
func (c *Class) ModifyDexterity(value int) {
	c.Dexterity = value
}
func (c *Class) ModifyBodyDifficulty(value int) {
	c.BodyDifficulty = value
}
func (c *Class) ModifyIntelligence(value int) {
	c.Intelligence = value
}
func (c *Class) ModifyWisdom(value int) {
	c.Wisdom = value
}
func (c *Class) ModifyCharisma(value int) {
	c.Charisma = value
}
func (c *Class) ModifyStrength(value int) {
	c.Strength = value
}

func (c *Class) CalcLifePoints() {
	c.Points.HP = (c.BodyDifficulty * c.Strength) * c.Level
	c.Points.MP = (c.Intelligence * c.Wisdom) * c.Level
	c.Points.SP = (c.Dexterity * 5) * c.Level
}

func (c Class) lvlUp() {
	c.Level = c.Level + 1
}

type Race struct {
	Name string
}

func (c *Class) RaceAbilities() {
	switch c.Race {
	case "Human":
		c.ModifyStrength(c.Strength + 1)
		c.ModifyDexterity(c.Dexterity + 1)
		c.ModifyIntelligence(c.Intelligence + 1)
	case "Elf":
		c.ModifyDexterity(c.Dexterity + 2)
		c.ModifyIntelligence(c.Intelligence + 1)
	case "Dwarf":
		c.ModifyStrength(c.Strength + 2)
		c.ModifyCharisma(c.Charisma + 1)
	}
}

func RerollLuckyClassStats() Class {
	Lucky = Class{ClassName: "Lucky", Race: r.RandomRace().Name, Level: 1,
		Strength:       RandomRollPoints(),
		Dexterity:      RandomRollPoints(),
		BodyDifficulty: RandomRollPoints(),
		Intelligence:   RandomRollPoints(),
		Wisdom:         RandomRollPoints(),
		Charisma:       RandomRollPoints(),
	}
	return Lucky
}
