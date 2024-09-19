package log_stash

import "encoding/json"

type Leave int

const (
	DebugLeave Leave = 1
	InfoLeave  Leave = 2
	WarnLeave  Leave = 3
	ErrorLeave Leave = 4
)

func (s Leave) MarshalJSON() ([]byte, error) {

	return json.Marshal(s.String())
}

func (s Leave) String() string {
	var str string
	switch s {
	case DebugLeave:
		str = "debug"
	case InfoLeave:
		str = "info"
	case WarnLeave:
		str = "warn"
	case ErrorLeave:
		str = "error"
	default:
		str = "其他"
	}
	return str
}
