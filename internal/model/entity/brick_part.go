package entity

import "time"

type BrickPart struct {
	ID                  int64     `json:"id"`
	Number              string    `json:"number"`
	EnglishName         string    `json:"englishName"`
	ChineseName         string    `json:"chineseName"`
	SetOccurrenceNumber int32     `json:"setOccurrenceNumber"`
	ProducedStart       time.Time `json:"producedStart"`
	ProducedEnd         time.Time `json:"producedEnd"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	DeletedAt           time.Time `json:"deletedAt"`
}

type BrickPartCategory struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type BrickPartColor struct {
	ID            int64     `json:"id"`
	Number        string    `json:"number"`
	EnglishName   string    `json:"englishName"`
	ChineseName   string    `json:"chineseName"`
	BrickLinkName string    `json:"brickLinkName"`
	RGB           string    `json:"rgb"`
	ElementCount  int32     `json:"elementCount"`
	ProducedStart time.Time `json:"producedStart"`
	ProducedEnd   time.Time `json:"producedEnd"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
}

type BrickPartColorFamily struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type BrickPartDesign struct {
	ID           int64     `json:"id"`
	Number       string    `json:"number"`
	ElementCount int32     `json:"elementCount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}

type BrickPartPrice struct {
	ID        int64     `json:"id"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type BrickPartTag struct {
	ID          int64     `json:"id"`
	EnglishName string    `json:"englishName"`
	ChineseName string    `json:"chineseName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}
