package domain

import (
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

type CustomerRank string

const (
	CustomerRankBasic CustomerRank = "BASIC"
	CustomerRankGold  CustomerRank = "GOLD"
)

// Customer model
type Customer struct {
	shared.BaseModel
	Email    string
	Password string
	Salt     string
	Fullname string
	Rank     CustomerRank
	Status   string
	Bonus    float64
}

// IsValidPassword function
func (c *Customer) IsValidPassword(password string) bool {
	return shared.Pbkdf2Hasher.VerifyPassword(password, c.Password, c.Salt)
}

func (c *Customer) CollectRequest(cr *CustomerRequest) {
	if cr.Fullname != "" {
		c.Fullname = cr.Fullname

	}

	if cr.Status != "" {
		c.Status = cr.Status
	}

	if cr.Rank != "" {
		c.Rank = cr.Rank
	}
}
