package items

type Armor struct {
	Helmet    Helmet
	Chest     Chest
	Pants     Pants
	Gauntlets Gauntlets
}

type Helmet struct {
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

type Chest struct {
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

type Pants struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Rarity       Rarity           `json:"rarity"`
	Defend       int              `json:"defend"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

type Gauntlets struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Rarity       Rarity           `json:"rarity"`
	Defend       int              `json:"defend"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}
type Boots struct {
	Name         string           `json:"sword_name"`
	Type         string           `json:"type"`
	Rarity       Rarity           `json:"rarity"`
	Defend       int              `json:"defend"`
	Weight       int              `json:"weight"`
	Durability   int              `json:"durability"`
	Price        int              `json:"price"`
	Requirements ItemRequirements `json:"requirements"`
}

var ()
