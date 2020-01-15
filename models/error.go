package models

type (
	ResponseStatus struct {
		Status string `json:"status"`
	}
)

func InternelServerError() interface{} {
	response := new(ResponseStatus)
	response.Status = "Internel server error, please check logs"
	return response
}

func Created() interface{} {
	response := new(ResponseStatus)
	response.Status = "Successful created"
	return response
}

func Unauthorized() interface{} {
	response := new(ResponseStatus)
	response.Status = "Not authenticated"
	return response
}
