package errors

import "github.com/joomcode/errorx"

var (
	Code = errorx.RegisterProperty("code")
)

var (
	DataBaseError     = errorx.NewNamespace("_DATA_BASE_ERROR_")
	InvalidInputError = errorx.NewNamespace("_INVALID_INPUT_ERROR_")
)

// defining different error codes
