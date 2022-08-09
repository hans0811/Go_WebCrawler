package services

// import "fmt"

type lifecycle int

const (
	Transient lifecycle = iota
	Singleton
	Scoped
)