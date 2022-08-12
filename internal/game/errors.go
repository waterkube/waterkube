package game

import "errors"

var (
	// ErrNoMoney error.
	ErrNoMoney = errors.New("game: you have no money")

	// ErrNoBoat error.
	ErrNoBoat = errors.New("game: you have no boat")

	// ErrNoDiver error.
	ErrNoDiver = errors.New("game: you have no diver")

	// ErrNoSubmarine error.
	ErrNoSubmarine = errors.New("game: you have no submarine")

	// ErrNoArtifact error.
	ErrNoArtifact = errors.New("game: you have no artifact")

	// ErrInvalidGridName error.
	ErrInvalidGridName = errors.New("game: invalid grid name")

	// ErrInvalidGridType error.
	ErrInvalidGridType = errors.New("game: invalid grid type")

	// ErrInvalidGridStatus error.
	ErrInvalidGridStatus = errors.New("game: invalid grid status")
)
