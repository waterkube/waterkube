package artifact

// Combinable type.
type Combinable struct {
	Price  int
	Pair   string
	Result string
}

// ShallowCombinable list.
var ShallowCombinable = map[string]*Combinable{
	"Key": {
		Price:  2,
		Pair:   "Chest",
		Result: "Treasure",
	},
	"Chest": {
		Price:  2,
		Pair:   "Key",
		Result: "Treasure",
	},
}

// DeepCombinable list.
var DeepCombinable = map[string]*Combinable{
	"Code": {
		Price:  3,
		Pair:   "Safe",
		Result: "Gold",
	},
	"Safe": {
		Price:  3,
		Pair:   "Code",
		Result: "Gold",
	},
	"Crowbar": {
		Price:  4,
		Pair:   "Container",
		Result: "Automobile",
	},
	"Container": {
		Price:  4,
		Pair:   "Crowbar",
		Result: "Automobile",
	},
}
