package domain

// TokenClaim model
type TokenClaim struct {
	RoleUser string `json:"roleUser"`
	DeviceID string `json:"did"`
	User     struct {
		ID           string `json:"id"`
		DivisionCode string `json:"divisionCode"`
		Email        string `json:"email"`
		AppID        string `json:"appID"`
	} `json:"user"`
	JTI        string `json:"jti"`
	RefreshJTI string `json:"refreshJti"`
	Alg        string `json:"-"`
}
