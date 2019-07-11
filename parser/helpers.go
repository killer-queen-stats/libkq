package parser

import (
	"fmt"
	"strconv"

	. "github.com/ughoavgfhw/libkq/common"
)

type InvalidSideError string

func (e InvalidSideError) Error() string {
	return fmt.Sprintf("invalid side %q", string(e))
}
func parseSide(color []byte) (side Side, err error) {
	s := string(color)
	switch s {
	case "Blue":
		side = BlueSide
	case "Gold", "Red":
		side = GoldSide
	default:
		err = InvalidSideError(s)
	}
	return
}

func parsePosition(xBuf, yBuf []byte) (p Position, err error) {
	x, err := strconv.ParseInt(string(xBuf), 10, 0)
	if err != nil {
		return
	}
	y, err := strconv.ParseInt(string(yBuf), 10, 0)
	if err != nil {
		return
	}
	return Position{X: int(x), Y: int(y)}, nil
}

type InvalidPlayerError PlayerId

func (e InvalidPlayerError) Error() string {
	return fmt.Sprintf("invalid player id %v", int(e))
}
func parsePlayer(buf []byte) (id PlayerId, err error) {
	raw, err := strconv.ParseInt(string(buf), 10, 0)
	id = PlayerId(raw)
	if err == nil && !id.IsValid() {
		err = InvalidPlayerError(id)
	}
	return
}

type InvalidGateTypeError string

func (e InvalidGateTypeError) Error() string {
	return fmt.Sprintf("invalid gate type %q", string(e))
}
func parseMaidenType(buf []byte) (typ GateType, err error) {
	t := string(buf)
	switch t {
	case "maiden_wings":
		typ = WarriorGate
	case "maiden_speed":
		typ = SpeedGate
	default:
		err = InvalidGateTypeError(t)
	}
	return
}

type InvalidPlayerTypeError string

func (e InvalidPlayerTypeError) Error() string {
	return fmt.Sprintf("invalid player type %q", string(e))
}
func parsePlayerType(buf []byte) (typ PlayerType, err error) {
	t := string(buf)
	switch t {
	case "Worker":
		typ = Drone
	case "Soldier":
		typ = Warrior
	case "Queen":
		typ = Queen
	default:
		err = InvalidPlayerTypeError(t)
	}
	return
}

type InvalidMapError string

func (e InvalidMapError) Error() string {
	return fmt.Sprintf("invalid map %q", string(e))
}
func parseMap(buf []byte) (m Map, err error) {
	t := string(buf)
	switch t {
	case "map_day":
		m = DayMap
	case "map_night":
		m = NightMap
	case "map_dusk":
		m = DuskMap
	default:
		err = InvalidMapError(t)
	}
	return
}

type InvalidWinTypeError string

func (e InvalidWinTypeError) Error() string {
	return fmt.Sprintf("invalid win condition %q", string(e))
}
func parseWinType(buf []byte) (wc WinType, err error) {
	t := string(buf)
	switch t {
	case "economic":
		wc = EconomicWin
	case "military":
		wc = MilitaryWin
	case "snail":
		wc = SnailWin
	default:
		err = InvalidWinTypeError(t)
	}
	return
}
