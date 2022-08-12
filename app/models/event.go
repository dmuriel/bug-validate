package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Type      string    `json:"type" db:"type"`
}
