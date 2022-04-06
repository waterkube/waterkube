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
		Price:  3,
		Pair:   "Chest",
		Result: "Treasure",
	},
	"Chest": {
		Price:  3,
		Pair:   "Key",
		Result: "Treasure",
	},
}

// DeepCombinable list.
var DeepCombinable = map[string]*Combinable{
	"Code": {
		Price:  4,
		Pair:   "Safe",
		Result: "Gold",
	},
	"Safe": {
		Price:  4,
		Pair:   "Code",
		Result: "Gold",
	},
	"Crowbar": {
		Price:  5,
		Pair:   "Container",
		Result: "Automobile",
	},
	"Container": {
		Price:  5,
		Pair:   "Crowbar",
		Result: "Automobile",
	},
}
