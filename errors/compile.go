package errors

import "net/http"

func CompileError(err error, lang string, debugMode bool) (int, error) {
	var (
		appError *AppError
		httpCode int
	)
	var debugErr *string
	if debugMode {
		errStr := err.Error()
		if len(errStr) > 0 {
			debugErr = &errStr
		}
	}
	code := ErrCode(err)
	switch code {

	case CodeValueInvalid, CodeHTTPBadRequest, CodeSQLRecordDoesNotMatch, CodeSQLRecordIsExpired:
		httpCode = http.StatusBadRequest
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "badrequest"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeHTTPNotFound:
		httpCode = http.StatusNotFound
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "notfound"),
			sys:          err,
			DebugError:   debugErr,
		}
	case CodeSQLUniqueConstraint:
		httpCode = http.StatusConflict
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "uniqueconst"),
			sys:          err,
			DebugError:   debugErr,
		}

	case CodeHTTPUnauthorized:
		httpCode = http.StatusUnauthorized
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "unauthorized"),
			sys:          err,
			DebugError:   debugErr,
		}

	default:
		httpCode = http.StatusInternalServerError
		appError = &AppError{
			Code:         int(code),
			HumanMessage: EM.Message(lang, "internal"),
			sys:          err,
			DebugError:   debugErr,
		}
	}

	return httpCode, appError
}
