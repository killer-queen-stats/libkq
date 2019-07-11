package parser

import (
	"strconv"
	"time"

	. "github.com/ughoavgfhw/libkq/common"
)

type GameStartMessage struct {
	Map Map
	_   bool
	_   bool
}

type GameEndMessage struct {
	Map      Map
	_        bool
	Duration time.Duration
	_        bool
}

type GameResultMessage struct {
	Winner       Side
	EndCondition WinType
}

func init() {
	registerRegexpValueParser(
		"gamestart", `^([^,]*),(True|False),0,(True|False)$`, // Time should always be 0 at the start
		func(match [][]byte) (msg interface{}, err error) {
			// Booleans unknown, only seen as false. The reddit thread suggests one may be whether the game is in the attract loop.
			var m GameStartMessage
			m.Map, err = parseMap(match[1])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"gameend", `^([^,]*),(True|False),([\d.]+),(True|False)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m GameEndMessage
			m.Map, err = parseMap(match[1])
			if err != nil {
				return
			}
			durSeconds, err := strconv.ParseFloat(string(match[3]), 64)
			if err != nil {
				return
			}
			durNanos := durSeconds * float64(time.Second/time.Nanosecond)
			m.Duration = time.Duration(durNanos) * time.Nanosecond
			return m, nil
		},
	)
	registerRegexpValueParser(
		"victory", `^(Blue|Gold),([^,]*)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m GameResultMessage
			m.Winner, err = parseSide(match[1])
			if err != nil {
				return
			}
			m.EndCondition, err = parseWinType(match[2])
			if err != nil {
				return
			}
			return m, nil
		},
	)
}
