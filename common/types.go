package common

type Side int8

const (
	Neutral Side = iota
	BlueSide
	GoldSide
)

func (s Side) String() string {
	switch s {
	case Neutral:
		return "neutral"
	case BlueSide:
		return "blue"
	case GoldSide:
		return "gold"
	default:
		return "invalid_side"
	}
}

type Position struct {
	X int
	Y int
}

type PlayerId int8

const (
	// Values are explicitly assigned, since they specifically match values used in messages from the game.
	GoldQueen   PlayerId = 1
	BlueQueen   PlayerId = 2
	GoldStripes PlayerId = 3
	BlueStripes PlayerId = 4
	GoldAbs     PlayerId = 5
	BlueAbs     PlayerId = 6
	GoldSkulls  PlayerId = 7
	BlueSkulls  PlayerId = 8
	GoldChecks  PlayerId = 9
	BlueChecks  PlayerId = 10

	NumPlayers = 10
)

func (id PlayerId) IsValid() bool {
	return id >= GoldQueen && id <= BlueChecks
}
func (id PlayerId) IsQueen() bool {
	return id == GoldQueen || id == BlueQueen
}
func (id PlayerId) Team() Side {
	if id&1 == 0 {
		return BlueSide
	} else {
		return GoldSide
	}
}
func (id PlayerId) Index() int {
	return int(id) - 1
}
func (id PlayerId) String() string {
	switch id {
	case GoldQueen:
		return "gold queen"
	case BlueQueen:
		return "blue queen"
	case GoldStripes:
		return "gold stripes"
	case BlueStripes:
		return "blue stripes"
	case GoldAbs:
		return "gold abs"
	case BlueAbs:
		return "blue abs"
	case GoldSkulls:
		return "gold skulls"
	case BlueSkulls:
		return "blue skulls"
	case GoldChecks:
		return "gold checks"
	case BlueChecks:
		return "blue checks"
	default:
		return "invalid player"
	}
}

type GateType int8

const (
	_ GateType = iota
	WarriorGate
	SpeedGate
)

func (g GateType) String() string {
	switch g {
	case WarriorGate:
		return "warrior"
	case SpeedGate:
		return "speed"
	default:
		return "invalid_gate"
	}
}

type PlayerType int8

const (
	_ PlayerType = iota
	Robot
	Drone
	Warrior
	Queen
)

func (t PlayerType) String() string {
	switch t {
	case Robot:
		return "robot"
	case Drone:
		return "drone"
	case Warrior:
		return "warrior"
	case Queen:
		return "queen"
	default:
		return "invalid_player_type"
	}
}

type Map int8

const (
	_ Map = iota
	DayMap
	NightMap
	DuskMap

	WarriorBonusMap
	SnailBonusMap
)

func (m Map) String() string {
	switch m {
	case DayMap:
		return "day"
	case NightMap:
		return "night"
	case DuskMap:
		return "dusk"
	case WarriorBonusMap:
		return "warrior_bonus"
	case SnailBonusMap:
		return "snail_bonus"
	default:
		return "invalid_map"
	}
}

type WinType int8

const (
	_ WinType = iota
	EconomicWin
	MilitaryWin
	SnailWin
)

func (w WinType) String() string {
	switch w {
	case EconomicWin:
		return "economic"
	case MilitaryWin:
		return "military"
	case SnailWin:
		return "snail"
	default:
		return "invalid_win_type"
	}
}
