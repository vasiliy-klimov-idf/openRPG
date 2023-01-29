package races

import rand "github.com/mazen160/go-random"

type Races struct {
	Name  string
	Buffs []buffs
}

type buffs struct {
	Name        string
	Description string
	UsesInDay   int
}

var (
	Human = Races{Name: "Human", Buffs: []buffs{
		{Name: "God's Gift", Description: "Can up random point on +1 once in the battle", UsesInDay: 1},
		{Name: "Humanity", Description: "Can up Charisma point on +2 one time in day", UsesInDay: 1},
	}}
	Elf = Races{Name: "Elf", Buffs: []buffs{
		{}, {},
	}}
	Dwarf = Races{Name: "Dwarf", Buffs: []buffs{
		{}, {},
	}}

	RaceArray = []Races{Human, Elf, Dwarf}
)

func RandomRace() Races {
	randomIndex := rand.GetIntInsecure(len(RaceArray))
	return RaceArray[randomIndex]
}
