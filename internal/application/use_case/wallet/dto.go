package wallet

type Dto struct {
	UId           string `json:"user_id"`
	AccountNumber string `json:"account_number"`
	Balance       int64  `json:"balance"`
}
