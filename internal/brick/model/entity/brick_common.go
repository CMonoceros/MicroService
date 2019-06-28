package entity

import (
	"time"
)

type BrickBrand struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickBrand) TableName() string {
	return "brick_brand"
}

type BrickPackaging struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickPackaging) TableName() string {
	return "brick_packaging"
}

type Country struct {
	ID             int64     `json:"id"`
	EnglishName    string    `json:"englishName"`
	ChineseName    string    `json:"chineseName"`
	Code           string    `json:"code"`
	CurrencyUnit   string    `json:"currencyUnit"`
	CurrencySymbol string    `json:"currencySymbol"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

type BrickMedia struct {
	BaseModel
	Type     int8    `json:"type" gorm:"column:type"`
	Src      string  `json:"src" gorm:"column:src"`
	High     int32   `json:"high" gorm:"column:high"`
	Width    int32   `json:"width" gorm:"column:width"`
	Duration float32 `json:"duration" gorm:"column:duration"`
}

func (*BrickMedia) TableName() string {
	return "brick_media"
}
