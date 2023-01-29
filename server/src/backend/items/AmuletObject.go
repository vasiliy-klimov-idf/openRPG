package items

type Amulet struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Level        int              `json:"level"`
	Rarity       Rarity           `json:"rarity"`
	Buff         string           `json:"buff"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

var ()
