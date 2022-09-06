package errors

import "golang_bootstrap_project/environment"

type BaseError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type BaseErrorReturn struct {
	Success bool       `json:"success"`
	Error   *BaseError `json:"error"`
}
type ErrorHandlerDTO struct {
	ErrorCode string `json:"error_code" xml:"error_code" form:"error_code" query:"error_code"`
	Message   string `json:"message" xml:"message" form:"message" query:"message"`
}

func TechnicalErrorNoSubWasGiven() BaseErrorReturn {
	return BaseErrorReturn{
		Success: false,
		Error: &BaseError{
			Code:    "T003",
			Message: "No sub was given in the query parameter",
		},
	}
}

func TechnicalErrorWasOccurred() BaseErrorReturn {
	return BaseErrorReturn{
		Success: false,
		Error: &BaseError{
			Code:    environment.TechnicalErrorCode,
			Message: environment.TechnicalErrorBusinessMessage,
		},
	}
}
