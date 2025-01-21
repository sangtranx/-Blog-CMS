package common

import (
	"fmt"
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"create_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"update_at,omitempty" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
	fmt.Printf("Debug: ID : %v\n FakeId generated: %v\n", m.Id, m.FakeId.String())
}
