package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ResponseWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorWithData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type ResponseWithOutData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type StructKosong struct {
}

func ApiResponse(status bool, message string, data interface{}) interface{} {
	if data == nil {
		var structkosong StructKosong
		jsonResponse := ResponseErrorWithData{
			Success: status,
			Message: message,
			Error:   structkosong,
		}
		return jsonResponse
	} else {
		jsonResponse := ResponseWithData{
			Success: status,
			Message: message,
			Data:    data,
		}
		return jsonResponse
	}

}

func ApiWithOutData(status bool, message string) interface{} {
	jsonResponse := ResponseWithOutData{
		Success: status,
		Message: message,
	}
	return jsonResponse
}

func ValidationError(err error) []string {
	var resultMessage []string
	if castedObject, ok := err.(validator.ValidationErrors); ok {

		for _, v := range castedObject {
			var message string
			switch v.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", v.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid email", v.Field())
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param())
			default:
				message = v.Error()
			}

			resultMessage = append(resultMessage, message)
		}
	}

	return resultMessage

}
