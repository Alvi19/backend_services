package reqresp

type SortRequest struct {
	Selector string `json:"selector"`
	Desc     bool   `json:"desc"`
}
