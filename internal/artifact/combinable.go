package artifact

// ShallowCombinable list.
var ShallowCombinable = []*Combinable{
	Chest,
	Key,
}

// DeepCombinable list.
var DeepCombinable = []*Combinable{
	// TODO
}

// Combinable type.
type Combinable struct {
	Name  string
	Price int
	Pair  string
}

// Chest artifact.
var Chest = &Combinable{
	Name:  "Chest",
	Price: 400,
	Pair:  "Key",
}

// Key artifact.
var Key = &Combinable{
	Name:  "Key",
	Price: 50,
	Pair:  "Chest",
}
