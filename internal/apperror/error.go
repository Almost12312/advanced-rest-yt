package apperror

import "encoding/json"

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000003")
)

type AppError struct {
	Err              error  `json:"-"`
	Msg              string `json:"msg,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e AppError) Error() string {
	return e.Msg
}

func (e AppError) Unwrap() error {
	return e.Err
}

func (e AppError) Marshal() []byte {
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
