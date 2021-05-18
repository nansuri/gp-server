package model

// Get Token API
type CipherRequest struct {
	Data     string `json:"data"`
	DataByte string `json:"data_byte"`
}

type CipherResponse struct {
	Data          string `json:"data"`
	EncryptedData []byte `json:"encrypted_data"`
	ExtendInfo    string `json:"extend_info"`
}

type GeneralHeader struct {
	Token string `json:token`
}
