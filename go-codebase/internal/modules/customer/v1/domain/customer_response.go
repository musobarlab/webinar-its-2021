package domain

// CustomerResponse request model
type CustomerResponse struct {
	Email    string  `json:"email"`
	Fullname string  `json:"fullname"`
	Status   string  `json:"status"`
	Rank     string  `json:"rank"`
	Bonus    float64 `json:"bonus"`
}

func CustomerResponseFromCustomer(c *Customer) CustomerResponse {
	var cr CustomerResponse
	cr.Email = c.Email
	cr.Fullname = c.Fullname
	cr.Status = c.Status
	cr.Rank = string(c.Rank)
	cr.Bonus = c.Bonus
	return cr
}
