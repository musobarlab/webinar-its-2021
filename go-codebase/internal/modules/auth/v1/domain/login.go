package domain

// RequestLogin model
type RequestLogin struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ApplicationType string `json:"applicationType"`
}

// ResponseLogin model
type ResponseLogin struct {
	AccessToken           string   `json:"accessToken,omitempty"`
	RefreshToken          string   `json:"refreshToken,omitempty"`
	AccessTokenExpiresIn  int      `json:"accessTokenExpiresIn,omitempty"`
	RefreshTokenExpiresIn int      `json:"refreshTokenExpiresIn,omitempty"`
	Subscribe             []string `json:"subscribe"`
	Role                  string   `json:"role,omitempty"`
}
