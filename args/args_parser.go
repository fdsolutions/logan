package args

import (
	"regexp"
)

// ParamParser is used to parse arg parameters.
// It implements the parser.Parser interface to define its own way to parse input.
type ParamParser interface {
	Parse(string) [][]string
}

// ParamParserImpl is a concrete params parser
type ParamParserImpl struct{}

// NewParamParser is a param parser contructor
func NewParamParser() *ParamParserImpl {
	return &ParamParserImpl{}
}

// Parse define the way arg parameters must be parsed
func (pp *ParamParserImpl) Parse(input string) (params [][]string) {
	r := regexp.MustCompile(argsParamsRegexPattern)
	params = r.FindAllStringSubmatch(input, -1)
	return
}
