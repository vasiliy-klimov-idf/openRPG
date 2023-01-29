package items

type Sword struct {
	OneHandSword OneHandSword
	TwoHandSword TwoHandSword
	Dagger       Dagger
}

type OneHandSword struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Level        int              `json:"level"`
	Rarity       Rarity           `json:"rarity"`
	Damage       int              `json:"damage"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

type TwoHandSword struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Level        int              `json:"level"`
	Rarity       Rarity           `json:"rarity"`
	Damage       int              `json:"damage"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

type Dagger struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Level        int              `json:"level"`
	Rarity       Rarity           `json:"rarity"`
	Damage       int              `json:"damage"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

var (
	beginnersSword = OneHandSword{Name: "Beginner's Sword",
		Rarity:       Common,
		Damage:       10,
		Weight:       5,
		Durability:   30,
		Price:        10,
		Requirements: ItemRequirements{Strength: 10, Dexterity: 5},
	}
	oldKnightSword = OneHandSword{Name: "Old knight sword"}
	generalSword   = OneHandSword{Name: "Sword of the general"}
	recruitsSword  = OneHandSword{Name: "Recruit's Sword"}
	officersSword  = OneHandSword{Name: "Officer's sword"}
)
