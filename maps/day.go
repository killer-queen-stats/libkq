package maps

import . "github.com/ughoavgfhw/libkq/common"

var Day = &MapMetadata{
	Hives: []HiveMetadata{
		{Holes: make([]Position, 12)},
		{Holes: make([]Position, 12)},
	},
	WarriorGates: []GateMetadata{
		{Pos: Position{560, 260}},
		{Pos: Position{960, 500}},
		{Pos: Position{1360, 260}},
	},
	SpeedGates: []GateMetadata{
		{Pos: Position{410, 860}},
		{Pos: Position{1510, 860}},
	},
	Snails: []SnailMetadata{
		{
			Nets: [2]Position{
				{60, 11},
				{1860, 11},
			},
		},
	},
	BerriesAvailable:      60,
	QueenLives:            3,
	FirstLifeSpeedQueen:   false,
	FirstLifeSpeedWarrior: false,
}
