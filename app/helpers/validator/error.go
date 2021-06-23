package validator

import "fmt"

func Required(attribute string) string {
	return fmt.Sprintf("Sorry, The %s field is required", attribute)
}

func Error(message string) string {
	return fmt.Sprintf("Sorry, Error : %s", message)
}

func GeneralFail(message interface{}) string {
	if message != nil {
		return fmt.Sprintf(message.(string))
	}
	return fmt.Sprintf("Sorry, something went wrong, please try again")
}

func FailedToGetData(message interface{}) string {
	text := "Sorry, failed to get data"
	if message == nil {
		return fmt.Sprintf(text)
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToCreateData(message interface{}) string {
	text := "Sorry, failed to create data"
	if message == nil {
		return fmt.Sprintf(text)
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToDeleteData(message interface{}) string {
	text := "Sorry, failed to delete data"
	if message == nil {
		return fmt.Sprintf(text)
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func FailedToUpdateData(message interface{}) string {
	text := "Sorry, failed to update data"
	if message == nil {
		return fmt.Sprintf(text)
	}
	return fmt.Sprintf(text+" %s", message.(string))
}

func DataNotFound(message string) string {
	return fmt.Sprintf("Sorry, data %s not found", message)
}

func DataNotValid(message string) string {
	return fmt.Sprintf("Sorry, data %s not valid", message)
}

func NotValid(message string) string {
	return fmt.Sprintf("Sorry, %s not valid", message)
}
