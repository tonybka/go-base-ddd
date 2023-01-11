package entity

import "github.com/google/uuid"

type BaseEntity struct {
	ID uuid.UUID
}

func NewBaseEntity() (BaseEntity, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return BaseEntity{}, err
	}

	return BaseEntity{ID: id}, nil
}
