package parser

import . "github.com/ughoavgfhw/libkq/common"

type PlayerKillMessage struct {
	Pos        Position
	Killer     PlayerId
	Victim     PlayerId
	VictimType PlayerType
}

type GlanceMessage struct {
	Player1 PlayerId // Initiator?
	Player2 PlayerId
}

func init() {
	registerRegexpValueParser(
		"playerKill", `^(\d+),(\d+),(\d+),(\d+)(?:,([^,]*)?)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var pkm PlayerKillMessage
			pkm.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			pkm.Killer, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			pkm.Victim, err = parsePlayer(match[4])
			if err != nil {
				return
			}
			if len(match) >= 6 && match[5] != nil {
				pkm.VictimType, err = parsePlayerType(match[5])
				if err != nil {
					return
				}
			}
			return pkm, nil
		},
	)
	registerRegexpValueParser(
		"glance", `^(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var gm GlanceMessage
			gm.Player1, err = parsePlayer(match[1])
			if err != nil {
				return
			}
			gm.Player2, err = parsePlayer(match[2])
			if err != nil {
				return
			}
			return gm, nil
		},
	)
}
