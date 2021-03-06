package domain

import (
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// CustomerRequest request model
type CustomerRequest struct {
	CustomerID string
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	Fullname   string       `json:"fullname"`
	Status     string       `json:"status"`
	Rank       CustomerRank `json:"rank"`
}

func (cr *CustomerRequest) ToCustomerModel() (*Customer, error) {
	var c Customer
	hashed := shared.Pbkdf2Hasher.HashPassword(cr.Password)
	c.Password = hashed.CipherText
	c.Salt = hashed.Salt

	c.Email = cr.Email
	c.Fullname = cr.Fullname
	c.Status = "ACTIVE"
	c.Rank = CustomerRankBasic

	return &c, nil
}
