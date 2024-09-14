package model

type APIResponse struct {
	Status int         `json:"status"`
	Des    string      `json:"des"`
	Data   interface{} `json:"data,omitempty"`
	Page   int         `json:"page,omitempty"`
	Size   int         `json:"size,omitempty"`
	Total  int64       `json:"total,omitempty"`
}
