package entity

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
