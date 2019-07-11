package libkq

import "time"
import . "github.com/ughoavgfhw/libkq/common"

type GameState struct {
	Start, End         time.Time
	Map                Map

	Players            [NumPlayers]PlayerState
	BlueTeam, GoldTeam TeamState
	Snails             []SnailState // The number of snails depends on the map.
	WarriorGates       []GateState
	SpeedGates         []GateState
	BerriesUsed        int
	FamineStart        time.Time
	NumFamines         int

	Winner       Side
	EndCondition WinType
}

func (game *GameState) InGame() bool {
	return !game.Start.IsZero() && game.End.IsZero()
}
func (game *GameState) StartFamine(when time.Time) {
	game.FamineStart = when
	game.NumFamines++
}
func (game *GameState) EndFamine() {
	game.FamineStart = time.Time{}
	game.BerriesUsed = 0
}
func (game *GameState) InFamine() bool { return !game.FamineStart.IsZero() }

type PlayerState struct {
	Type     PlayerType
	HasSpeed bool
	HasBerry bool
	OnSnail  int
}

func (ps *PlayerState) Respawn() {
	if ps.Type == Warrior { ps.Type = Drone }
	ps.HasSpeed = false
	ps.HasBerry = false
	ps.OnSnail = 0
}
func (ps *PlayerState) IsOnSnail() bool {
	return ps.OnSnail > 0
}
func (ps *PlayerState) IsRobot() bool {
	return ps.Type == Robot
}

type TeamState struct {
	Warriors      int  // Vanilla + speed.
	SpeedWarriors int
	QueenDeaths   int
	BerriesIn     int
}

type SnailState struct {
	// Position 0 is neutral, extends to ±MaxPos with positive to the left.
	Pos, MaxPos int
}

type GateState struct {
	ClaimedBy Side
}
