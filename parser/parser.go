package parser

import (
	"fmt"
	"regexp"
)

type UnknownMessage struct {
	Value string
}

type InvalidMessageError struct {
	Err     string
	Message string
}

func (e *InvalidMessageError) Error() string {
	return fmt.Sprint(e.Err, ": ", e.Message)
}

type Parser struct{}

var baseRegexp = regexp.MustCompile(`^!\[k\[([^]]+)\],v\[(.*)\]\]!$`)

func (p Parser) Parse(msg []byte) (key []byte, val interface{}, err error) {
	match := baseRegexp.FindSubmatch(msg)
	if match == nil {
		err = &InvalidMessageError{"Could not parse into key and value",
			string(msg)}
		return
	}
	key = match[1]
	keyStr := string(key)
	valParse, ok := valueParserMap[keyStr]
	if !ok {
		valParse = valueParserFunc(func(val []byte) (interface{}, error) {
			return &UnknownMessage{string(val)}, nil
		})
	}
	val, err = valParse.parse(match[2])
	return
}

type regexpValueParser struct {
	key     string
	re      *regexp.Regexp
	builder func([][]byte) (interface{}, error)
}

func (rvp *regexpValueParser) parse(val []byte) (interface{}, error) {
	match := rvp.re.FindSubmatch(val)
	if match == nil {
		return nil, &InvalidMessageError{
			fmt.Sprint("could not parse value of type ", rvp.key),
			string(val),
		}
	}
	return rvp.builder(match)
}
func registerRegexpValueParser(
	key, re string, builder func([][]byte) (interface{}, error)) {
	valueParserMap[key] = &regexpValueParser{
		key:     key,
		re:      regexp.MustCompile(re),
		builder: builder,
	}
}
