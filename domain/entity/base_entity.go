package entity

type BaseEntity struct {
	ID uint
}

func NewBaseEntity(id uint) BaseEntity {
	return BaseEntity{ID: id}
}
