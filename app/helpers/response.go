package helpers

import (
	"api-boilerplate/domain"
	"encoding/json"
)

// ErrorResponse :
func ErrorResponse(code int, message string, err error, validation map[string]interface{}) domain.Response {
	if err != nil {
		// fmt.Println(err)
		ErrResp2Struct(err)
	}
	return domain.Response{
		Data:       json.NewEncoder(nil),
		Message:    message,
		Status:     code,
		Validation: validation,
	}
}

func SuccessResponse(message string, data interface{}) domain.Response {
	return domain.Response{
		Data:       data,
		Message:    message,
		Status:     200,
		Validation: map[string]interface{}{},
	}
}

// logging into sentry
func ErrResp2Struct(err error) {
	// if err != nil && os.Getenv("LOG_PARAM_FILE") == "ON" {
	// 	WriteLogError(inputFmt, err.Error()+", Details: "+errTitle)
	// }
}
