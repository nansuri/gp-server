package model

// Get Token API
type GeneralRequest struct {
	Data     string `json:"data"`
	DataByte string `json:"data_byte"`
}

type GeneralResponse struct {
	Data          string `json:"data"`
	EncryptedData []byte `json:"encrypted_data"`
}
