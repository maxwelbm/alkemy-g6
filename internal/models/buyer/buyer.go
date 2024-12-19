package modelsBuyer

type Buyer struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
}

type BuyerDTO struct {
	Id           *int    `json:"id,omitempty"`
	CardNumberId *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}
