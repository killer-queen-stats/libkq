package parser

import . "github.com/ughoavgfhw/libkq/common"

type PickUpBerryMessage struct {
	Player PlayerId
}

type DepositBerryMessage struct {
	Pos    Position
	Player PlayerId
}

type KickInBerryMessage struct {
	Pos    Position
	Player PlayerId
}

func init() {
	registerRegexpValueParser(
		"carryFood", `^\d+$`,
		func(match [][]byte) (msg interface{}, err error) {
			var pubm PickUpBerryMessage
			pubm.Player, err = parsePlayer(match[0])
			if err != nil {
				return
			}
			return pubm, nil
		},
	)
	registerRegexpValueParser(
		"berryDeposit", `^(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m DepositBerryMessage
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
		"berryKickIn", `^(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m KickInBerryMessage
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
}
