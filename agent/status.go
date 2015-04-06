package agent

type StatusCode int

const (
	StatusSuccess StatusCode = iota // 0
	StatusFail
)

type Status struct {
	Code  StatusCode
	Input interface{}
}

func NewStatus(code StatusCode, input interface{}) Status {
	return Status{code, input}
}
