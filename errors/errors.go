package errors

import (
    "encoding/json"
)

type Error struct {
    Code    int32   `json:"code"`
    Detail  string  `json:"detail"`
}

func (e *Error) Error() string {
    b, _ := json.Marshal(e)
    return string(b)
}

func New(code int32, detail string) error {
    return &Error{
        Code:   code,
        Detail: detail,
    }
}
