package helpers

import()



type Response struct{
	Status		Status		`json:"status"`
	Data		any			`json:"data"`
}

type Status struct{
	Message		string		`json:"message"`
	Code		int			`json:"code"`
}


func APIResponse[D any](message string, code int, data D) Response{
	status := Status{
		Message : message,
		Code : code,
	}

	jsonResponse := Response{
		Status : status,
		Data : data,
	}

	return jsonResponse
}

type ResponseFailed struct {
	Meta MetaFailed `json:"meta"`
	Data any  		`json:"data"`
}

type MetaFailed struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponseFailed[D any](message string, code int, status string, data D) ResponseFailed {
	meta := MetaFailed{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := ResponseFailed{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}