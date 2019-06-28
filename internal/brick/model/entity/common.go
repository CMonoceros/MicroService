package entity

import "time"

type BaseModel struct {
	ID        uint64     `json:"id" gorm:"column:id;primary_key"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

type BasePage struct {
	offset uint
	Limit  uint `json:"limit" form:"limit"`
	Page   uint `json:"page" form:"page"`
}

func (page *BasePage) GetOffset() (offset uint) {
	if page.offset != 0 {
		offset = page.offset
	} else {
		offset = (page.Page - 1) * page.Limit
		page.offset = offset
	}
	return
}
