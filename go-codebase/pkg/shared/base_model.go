package shared

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BaseModel a parent struct for all model
type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}
