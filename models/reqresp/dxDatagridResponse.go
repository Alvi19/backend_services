package reqresp

type DxDatagridResponse struct {
	Data       interface{} `json:"data"`
	TotalCount int         `json:"totalCount"`
	Summary    []int       `json:"summary"`
	Meta       interface{} `json:"meta"`
}
