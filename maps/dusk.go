package maps

import . "github.com/ughoavgfhw/libkq/common"

var Dusk = &MapMetadata{
	Hives: []HiveMetadata{
		{Holes: make([]Position, 12)},
		{Holes: make([]Position, 12)},
	},
	WarriorGates: []GateMetadata{
		{Pos: Position{310, 620}},
		{Pos: Position{960, 140}},
		{Pos: Position{1610, 620}},
	},
	SpeedGates: []GateMetadata{
		{Pos: Position{340, 140}},
		{Pos: Position{1580, 140}},
	},
	Snails: []SnailMetadata{
		{
			Nets: [2]Position{
				{60, 11},
				{1860, 11},
			},
		},
	},
	BerriesAvailable:      66,
	QueenLives:            3,
	FirstLifeSpeedQueen:   false,
	FirstLifeSpeedWarrior: false,
}
