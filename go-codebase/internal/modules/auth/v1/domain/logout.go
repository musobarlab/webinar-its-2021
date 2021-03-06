package domain

// ResponseLogout represent response after logout operation
type ResponseLogout struct {
	Subscribe []string `json:"subscribe"`
}
