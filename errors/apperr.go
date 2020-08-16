package errors

type AppError struct {
	Code         int     `json:"code"`
	HumanMessage string  `json:"message"`
	DebugError   *string `json:"debug,omitempty"`
	sys          error
}

func (e *AppError) Error() string {
	return e.sys.Error()
}
