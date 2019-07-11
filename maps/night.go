package maps

import . "github.com/ughoavgfhw/libkq/common"

var Night = &MapMetadata{
	Hives: []HiveMetadata{
		{Holes: make([]Position, 12)},
		{Holes: make([]Position, 12)},
	},
	WarriorGates: []GateMetadata{
		{Pos: Position{700, 260}},
		{Pos: Position{960, 700}},
		{Pos: Position{1220, 260}},
	},
	SpeedGates: []GateMetadata{
		{Pos: Position{170, 740}},
		{Pos: Position{1750, 740}},
	},
	Snails: []SnailMetadata{
		{
			Nets: [2]Position{
				{270, 491},  // ?
				{1650, 491}, // ?
			},
		},
	},
	BerriesAvailable:      60,
	QueenLives:            3,
	FirstLifeSpeedQueen:   false,
	FirstLifeSpeedWarrior: false,
}
