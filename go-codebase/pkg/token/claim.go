package token

const (

	// HS256 const
	HS256 = "HS256"

	// RS256 const
	RS256 = "RS256"
)

// Claim model
type Claim struct {
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
