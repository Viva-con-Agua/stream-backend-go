package models

type (
	ResponseStatus struct {
		Status string `json:"status"`
	}
	ResponseStatusId struct {
		Status string `json:"status"`
		Uuid   string `json:"uuid"`
	}
)

func InternelServerError() interface{} {
	response := new(ResponseStatus)
	response.Status = "Internel server error, please check logs"
	return response
}
func Conflict() interface{} {
	response := new(ResponseStatus)
	response.Status = "Models already exists"
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
func NoContent(uuid string) interface{} {
	response := new(ResponseStatusId)
	response.Status = "Not found"
	response.Uuid = uuid
	return response
}
func Updated(uuid string) interface{} {
	response := new(ResponseStatusId)
	response.Status = "Successful updated"
	response.Uuid = uuid
	return response
}

func Deleted(uuid string) interface{} {
	response := new(ResponseStatusId)
	response.Status = "Successful deleted"
	response.Uuid = uuid
	return response
}
