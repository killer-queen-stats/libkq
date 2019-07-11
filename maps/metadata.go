package maps

import . "github.com/ughoavgfhw/libkq/common"

type MapMetadata struct {
	Hives                 []HiveMetadata
	WarriorGates          []GateMetadata
	SpeedGates            []GateMetadata
	Snails                []SnailMetadata
	BerriesAvailable      int
	QueenLives            int
	FirstLifeSpeedQueen   bool
	FirstLifeSpeedWarrior bool
}

type HiveMetadata struct {
	Holes []Position
}

type GateMetadata struct {
	Pos Position
}

type SnailMetadata struct {
	Nets [2]Position
}
