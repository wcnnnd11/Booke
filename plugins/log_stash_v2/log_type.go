package log_stash_v2

import "encoding/json"

type LogType int

const (
	LoginType   LogType = 1
	ActionType  LogType = 2
	RuntimeType LogType = 3
)

// String 转字符串
func (t LogType) String() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RuntimeType:
		return "runtimeType"
	}
	return ""
}

// MarshalJSON 自定义类型转换为json
func (t LogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}
