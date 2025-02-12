package reqresp

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponsePagination struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta"`
}

// Meta untuk metadata paginasi
type Meta struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	TotalRecords int64 `json:"totalRecords"`
	TotalPages   int   `json:"totalPages"`
}

type SuccessResponsePerangkat struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Perngkat interface{} `json:"totalRecords"`
}
