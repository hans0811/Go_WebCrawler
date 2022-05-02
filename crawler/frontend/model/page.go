package model

import "001_go_env/crawler/engine"

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}

type SearchTest struct{
	Hits int
	Start int
	Items []engine.Item
}
