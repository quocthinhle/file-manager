package entity

import "github.com/google/uuid"

type Content struct {
	ID          uuid.UUID
	Name        string
	Description string
	Type        string
	ParentID    uuid.UUID
	OwnerID     uuid.UUID
	Children    []Content
}
