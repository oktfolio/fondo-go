package common

type Result struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
}
