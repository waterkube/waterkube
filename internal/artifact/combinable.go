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
		Price:  30,
		Pair:   "Chest",
		Result: "Treasure",
	},
	"Chest": {
		Price:  30,
		Pair:   "Key",
		Result: "Treasure",
	},
}

// ShallowResult list.
var ShallowResult = map[string]*Unique{
	"Treasure": {
		Price: 100,
	},
}

// DeepCombinable list.
var DeepCombinable = map[string]*Combinable{
	"Code": {
		Price:  40,
		Pair:   "Safe",
		Result: "Gold",
	},
	"Safe": {
		Price:  40,
		Pair:   "Code",
		Result: "Gold",
	},
	"Crowbar": {
		Price:  50,
		Pair:   "Container",
		Result: "Automobile",
	},
	"Container": {
		Price:  50,
		Pair:   "Crowbar",
		Result: "Automobile",
	},
}

// DeepResult list.
var DeepResult = map[string]*Unique{
	"Gold": {
		Price: 200,
	},
	"Automobile": {
		Price: 200,
	},
}
