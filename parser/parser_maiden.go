package parser

import . "github.com/ughoavgfhw/libkq/common"

type ClaimGateMessage struct {
	Pos  Position
	Side Side
}

type EnterGateMessage struct {
	Pos    Position
	Player PlayerId
}

type LeaveGateMessage struct {
	Pos    Position
	Player PlayerId
}

type UseGateMessage struct {
	Pos    Position
	Type   GateType
	Player PlayerId
}

func init() {
	registerRegexpValueParser(
		"blessMaiden", `^(\d+),(\d+),(Blue|Gold|Red)$`,
		func(match [][]byte) (msg interface{}, err error) {
			pos, err := parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			side, err := parseSide(match[3])
			if err != nil {
				return
			}
			return ClaimGateMessage{pos, side}, nil
		},
	)
	registerRegexpValueParser(
		"reserveMaiden", `^(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m EnterGateMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Player, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"unreserveMaiden", `^(\d+),(\d+),,(\d+)$`, // The message actually has an extra comma
		func(match [][]byte) (msg interface{}, err error) {
			var m LeaveGateMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Player, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"useMaiden", `^(\d+),(\d+),([^,]*),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m UseGateMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Type, err = parseMaidenType(match[3])
			if err != nil {
				return
			}
			m.Player, err = parsePlayer(match[4])
			if err != nil {
				return
			}
			return m, nil
		},
	)
}
