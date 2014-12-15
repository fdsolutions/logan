package logan

type LoganAgent struct {}

type LoganAction struct {
	intent string
	target string
	context string
	parameters map[string]string
}


var Usage = `

`

func NewAgent() *LoganAgent {
	return &LoganAgent{}
}
