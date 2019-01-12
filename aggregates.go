package goes

import (
	"time"
)

type Aggregate interface {
	GetID() string
	incrementVersion()
	updateUpdatedAt(time.Time)
	Type() string
}

// BaseAggregate should be embedded in all your aggregates
type BaseAggregate struct {
	ID        string     `json:"id" gorm:"column:id;type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Version   uint64     `json:"version" gorm:"column:version"`
}

func (agg BaseAggregate) GetID() string {
	return agg.ID
}

func (agg *BaseAggregate) incrementVersion() {
	agg.Version += 1
}

func (agg *BaseAggregate) updateUpdatedAt(t time.Time) {
	agg.UpdatedAt = t
}
