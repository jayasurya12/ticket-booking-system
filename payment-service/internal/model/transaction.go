package model

type PaymentRequest struct {
	Token     string  `json:"token"`
	Amount    float64 `json:"amount"`
	Reference string  `json:"reference"`
}

type PaymentResponse struct {
	Status string `json:"status"`
	TxnId  string `json:"txnId"`
}
