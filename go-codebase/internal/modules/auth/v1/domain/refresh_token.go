package domain

// RefreshTokenResponse model
type RefreshTokenResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn int    `json:"refreshTokenExpiresIn"`
	AccessTokenExpiresIn  int    `json:"accessTokenExpiresIn"`
}
