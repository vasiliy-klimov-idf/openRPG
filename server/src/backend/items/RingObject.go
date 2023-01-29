package items

type Ring struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Rarity       Rarity           `json:"rarity"`
	Buff         string           `json:"buff"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

var ()
