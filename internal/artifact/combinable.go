package artifact

// Combinable type.
type Combinable struct {
	Price int
	Pair  string
}

// ShallowCombinable list.
var ShallowCombinable = map[string]*Combinable{
	"Chest": {
		Price: 400,
		Pair:  "Key",
	},
	"Key": {
		Price: 50,
		Pair:  "Chest",
	},
}

// DeepCombinable list.
var DeepCombinable = map[string]*Combinable{
	// TODO
}
