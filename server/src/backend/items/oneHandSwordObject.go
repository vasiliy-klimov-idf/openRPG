package items

type oneHandSword struct {
	SwordName    string             `json:"sword_name"`
	Rarity       string             `json:"rarity"`
	Damage       int                `json:"damage"`
	Weight       int                `json:"weight"`
	Durability   int                `json:"durability "`
	MarketPrice  int                `json:"market_price"`
	Requirements []ItemRequirements `json:"requirements"`
}

var (
	beginnersSword = oneHandSword{SwordName: "Beginner's Sword",
		Rarity:       "Common",
		Damage:       10,
		Weight:       5,
		Durability:   30,
		MarketPrice:  10,
		Requirements: []ItemRequirements{{Strength: 10, Dexterity: 5}},
	}
	oldKnightSword = oneHandSword{SwordName: "Old knight sword"}
	generalSword   = oneHandSword{SwordName: "Sword of the general"}
	recruitsSword  = oneHandSword{SwordName: "Recruit's Sword"}
	officersSword  = oneHandSword{SwordName: "Officer's sword"}
)
