package model

// Get Token API
type GeneralRequest struct {
	Data     string `json:"data"`
	DataByte string `json:"data_byte"`
}

type GeneralResponse struct {
	Data          string `json:"data"`
	EncryptedData []byte `json:"encrypted_data"`
	ExtendInfo    string `json:"extend_info"`
}

type GeneralHeader struct {
	Token string `json:token`
}
