package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Company struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Status    string    `json:"status" db:"status"`
}
