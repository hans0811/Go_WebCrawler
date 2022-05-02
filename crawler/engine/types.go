package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct{
	Url string
	Type string // for saving website and unique Id
	Id string
	Payload interface{}
}

// fot return empty obj
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
