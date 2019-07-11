package parser

type valueParser interface {
	parse([]byte) (interface{}, error)
}

type valueParserFunc func([]byte) (interface{}, error)

func (f valueParserFunc) parse(val []byte) (interface{}, error) {
	return f(val)
}

var valueParserMap = map[string]valueParser{}
