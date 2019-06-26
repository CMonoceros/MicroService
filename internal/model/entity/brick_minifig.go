package entity

import (
	"time"
)

type BrickMinifig struct {
	ID                   int64                   `json:"id"`
	Number               string                  `json:"number"`
	EnglishName          string                  `json:"englishName"`
	ChineseName          string                  `json:"chineseName"`
	CharacterEnglishName string                  `json:"characterEnglishName"`
	CharacterChineseName string                  `json:"characterChineseName"`
	ReleaseDate          time.Time               `json:"releaseDate"`
	SetOccurrenceNumber  int32                   `json:"setOccurrenceNumber"`
	CreatedAt            time.Time               `json:"createdAt"`
	UpdatedAt            time.Time               `json:"updatedAt"`
	DeletedAt            time.Time               `json:"deletedAt"`
	Brand                BrickBrand              `json:"brand"`
	Category             BrickMinifigCategory    `json:"category"`
	SubCategory          BrickMinifigSubCategory `json:"subCategory"`
	Tags                 []BrickMinifigTag       `json:"tags"`
	Media                []BrickMedia            `json:"media"`
}

type BrickMinifigCategory struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type BrickMinifigSubCategory struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type BrickMinifigPrice struct {
	ID        int64     `json:"id"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type BrickMinifigTag struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
