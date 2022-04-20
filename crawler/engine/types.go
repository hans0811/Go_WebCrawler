package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// fot return empty obj
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
