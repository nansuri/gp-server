package model

type GeneralRequest struct {
	DataInput string `json:"data_input"`
}

type GeneralResponse struct {
	DataOutput string `json:"data_output"`
	Status     string `json:"success"`
	Error      string `json:"error_detail"`
}
