package items

type Shield struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Level        int              `json:"level"`
	Rarity       Rarity           `json:"rarity"`
	Defend       int              `json:"defend"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

var ()
