package maps

import . "github.com/ughoavgfhw/libkq/common"

var WarriorBonus = &MapMetadata{
	Hives: []HiveMetadata{},
	WarriorGates: []GateMetadata{
		{}, // ?
		{}, // ?
		{}, // ?
	},
	SpeedGates: []GateMetadata{
		{}, // ?
	},
	Snails:                []SnailMetadata{},
	BerriesAvailable:      36, // ?
	QueenLives:            2,
	FirstLifeSpeedQueen:   true,
	FirstLifeSpeedWarrior: true,
}

var SnailBonus = &MapMetadata{
	Hives: []HiveMetadata{},
	WarriorGates: []GateMetadata{
		{}, // ?
		{}, // ?
		{}, // ?
	},
	SpeedGates: []GateMetadata{
		{}, // ?
		{}, // ?
		{}, // ?
	},
	Snails: []SnailMetadata{
		{
			Nets: [2]Position{
				{}, // ?
				{}, // ?
			},
		},
		{
			Nets: [2]Position{
				{}, // ?
				{}, // ?
			},
		},
		{
			Nets: [2]Position{
				{}, // ?
				{}, // ?
			},
		},
	},
	BerriesAvailable:      36, // ?
	QueenLives:            5,
	FirstLifeSpeedQueen:   true,
	FirstLifeSpeedWarrior: false,
}
