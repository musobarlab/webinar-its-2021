package domain

import "time"

const (
	// LoginAttempt const
	LoginAttempt = 4
	// LoginAttemptExpired const
	LoginAttemptExpired = 30 * time.Second

	// HS256 const
	HS256 = "HS256"

	// RS256 const
	RS256 = "RS256"
)
