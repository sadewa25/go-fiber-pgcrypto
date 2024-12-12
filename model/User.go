package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	BirthDate string
	KTP       string
	Gender    string
	Phone     string
}
