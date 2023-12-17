package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id       uuid.UUID
	CreateAt time.Time
	Name     string
	Price    float32
}
