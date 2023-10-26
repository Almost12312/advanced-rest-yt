package apperror

import (
	"encoding/json"
	"fmt"
)

type ErrFields map[string]string
type ErrParams map[string]string

var (
	ErrNotFound   = NewAppError(nil, "not found", "", "ADV-000003")
	ErrBadRequest = BadRequest("bad request", "bad request")
)

type AppError struct {
	Err              error     `json:"-"`
	Msg              string    `json:"msg,omitempty"`
	DeveloperMessage string    `json:"developer_message,omitempty"`
	Code             string    `json:"code,omitempty"`
	Fields           ErrFields `json:"fields,omitempty"`
	Params           ErrParams `json:"params,omitempty"`
}

func (e *AppError) WithFields(fields ErrFields) {
	e.Fields = fields
}
func (e *AppError) WithParams(params ErrParams) {
	e.Params = params
}

func (e *AppError) Error() string {
	return e.Msg
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	m, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return m
}

func NewAppError(error error, msg, devMsg, code string) *AppError {
	return &AppError{
		Err:              error,
		Msg:              msg,
		DeveloperMessage: devMsg,
		Code:             code,
	}
}

func systemError(err error) *AppError {
	return NewAppError(err, "internal system error", err.Error(), "US-000000")
}

func BadRequest(devMsg string, msg string) *AppError {
	return NewAppError(fmt.Errorf(msg), msg, devMsg, "ADV-000004")
}
