package items

type twoHandSword struct {
	SwordName    string             `json:"sword_name"`
	Rarity       string             `json:"rarity"`
	Damage       int                `json:"damage"`
	Weight       int                `json:"weight"`
	Durability   int                `json:"durability "`
	MarketPrice  int                `json:"market_price"`
	Requirements []ItemRequirements `json:"requirements"`
}

var ()
