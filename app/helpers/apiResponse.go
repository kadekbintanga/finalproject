package helpers

import()



type Response struct{
	Status		Status		`json:"status"`
	Meta		Meta		`json:"meta"`
	Data		any			`json:"data"`
}

type Meta struct{
	Page		int			`json:"page"`
	Limit		int			`json:"limit"`
	Total		int			`json:"total"`
}

type Status struct{
	Message		string		`json:"message"`
	Code		int			`json:"code"`
}


func APIResponse[D any](message string, code int, page int, limit int, total int, data D) Response{
	status := Status{
		Message : message,
		Code : code,
	}

	meta := Meta{
		Page : page,
		Limit : limit,
		Total : total,
	}

	jsonResponse := Response{
		Status : status,
		Meta : meta,
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