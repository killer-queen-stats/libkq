package parser

import "strconv"
import . "github.com/ughoavgfhw/libkq/common"

type AliveMessage string

type PlayerNamesMessage [10]string

type PlayerSpawnMessage struct {
	Player PlayerId
	Type   PlayerType
}

func init() {
	registerRegexpValueParser(
		"alive", `^\d\d?:\d\d?:\d\d? [AP]M$`,
		func(match [][]byte) (interface{}, error) {
			return AliveMessage(string(match[0])), nil
		},
	)
	registerRegexpValueParser(
		"playernames", `^([^,]*),([^,]*),([^,]*),([^,]*),([^,]*),([^,]*),([^,]*),([^,]*),([^,]*),([^,]*)$`,
		func(match [][]byte) (interface{}, error) {
			return PlayerNamesMessage{
				string(match[1]), string(match[2]), string(match[3]), string(match[4]), string(match[5]),
				string(match[6]), string(match[7]), string(match[8]), string(match[9]), string(match[10]),
			}, nil
		},
	)
	registerRegexpValueParser(
		"spawn", `^(\d+),(True|False)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m PlayerSpawnMessage
			m.Player, err = parsePlayer(match[1])
			if err != nil {
				return
			}
			isBot, err := strconv.ParseBool(string(match[2]))
			if err != nil {
				return
			}
			if isBot {
				m.Type = Robot
			} else if m.Player.IsQueen() {
				m.Type = Queen
			} else {
				m.Type = Drone
			}
			return m, nil
		},
	)
}
