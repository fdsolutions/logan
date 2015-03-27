package command

// Parser is a set of parsing operations
type Parser interface {
	Parse(input string) interface{}
}
