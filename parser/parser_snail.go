package parser

import . "github.com/ughoavgfhw/libkq/common"

type GetOnSnailMessage struct {
	Pos   Position
	Rider PlayerId
}

type GetOffSnailMessage struct {
	Pos   Position
	Rider PlayerId
}

type SnailStartEatMessage struct {
	Pos   Position
	Rider PlayerId
	Snack PlayerId
}

type SnailEscapeEatMessage struct {
	Pos     Position
	Escapee PlayerId
}

func init() {
	registerRegexpValueParser(
		"getOnSnail: ", `^(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m GetOnSnailMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Rider, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"getOffSnail: ", `^(\d+),(\d+),,(\d+)$`, // Extra comma is actually in the message
		func(match [][]byte) (msg interface{}, err error) {
			var m GetOffSnailMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Rider, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"snailEat", `^(\d+),(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m SnailStartEatMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Rider, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			m.Snack, err = parsePlayer(match[4])
			if err != nil {
				return
			}
			return m, nil
		},
	)
	registerRegexpValueParser(
		"snailEscape", `^(\d+),(\d+),(\d+)$`,
		func(match [][]byte) (msg interface{}, err error) {
			var m SnailEscapeEatMessage
			m.Pos, err = parsePosition(match[1], match[2])
			if err != nil {
				return
			}
			m.Escapee, err = parsePlayer(match[3])
			if err != nil {
				return
			}
			return m, nil
		},
	)
}
