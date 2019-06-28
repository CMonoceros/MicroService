package entity

import (
	"time"
)

type BrickSet struct {
	BaseModel
	Number             string            `json:"number" gorm:"column:number"`
	EnglishName        string            `json:"englishName" gorm:"column:english_name"`
	ChineseName        string            `json:"chineseName" gorm:"column:chinese_name"`
	ReleaseDate        time.Time         `json:"releaseDate" gorm:"column:release_date"`
	PartCount          int32             `json:"partCount" gorm:"column:part_count"`
	MinifigCount       int32             `json:"minifigCount" gorm:"column:minifig_count"`
	BarCode            string            `json:"barCode" gorm:"column:bar_code"`
	EnglishDescription string            `json:"englishDescription" gorm:"column:english_description"`
	ChineseDescription string            `json:"chineseDescription" gorm:"column:chinese_description"`
	AgeMin             int8              `json:"ageMin" gorm:"column:age_min"`
	AgeMax             int8              `json:"ageMax" gorm:"column:age_max"`
	TypeID             uint64            `json:"typeID" gorm:"column:type_id"`
	Type               *BrickSetType     `json:"type" gorm:"foreignkey:type_id;"`
	ThemeID            uint64            `json:"themeID" gorm:"column:theme_id"`
	Theme              *BrickSetTheme    `json:"theme" gorm:"foreignkey:theme_id;"`
	SubThemeID         uint64            `json:"subThemeID" gorm:"column:subtheme_id"`
	SubTheme           *BrickSetSubTheme `json:"subTheme" gorm:"foreignkey:subtheme_id;"`
	BrandID            uint64            `json:"brandID" gorm:"column:brand_id"`
	Brand              *BrickBrand       `json:"brand" gorm:"foreignkey:brand_id;"`
	PackagingID        uint64            `json:"packagingID" gorm:"column:packaging_id"`
	Packaging          *BrickPackaging   `json:"packaging" gorm:"foreignkey:packaging_id;"`
	Tags               []*BrickSetTag    `json:"tags" gorm:"many2many:brick_set_re_brick_tag;association_jointable_foreignkey:tag_id;jointable_foreignkey:set_id;"`
	Media              []*BrickMedia     `json:"media" gorm:"many2many:brick_set_re_brick_media;association_jointable_foreignkey:media_id;jointable_foreignkey:set_id;"`
}

func (*BrickSet) TableName() string {
	return "brick_set"
}

type BrickSetPrice struct {
	BaseModel
	Price float32 `json:"price" gorm:"column:price"`
}

func (*BrickSetPrice) TableName() string {
	return "brick_set_price"
}

type BrickSetTag struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickSetTag) TableName() string {
	return "brick_set_tag"
}

type BrickSetSubTheme struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickSetSubTheme) TableName() string {
	return "brick_set_subtheme"
}

type BrickSetTheme struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickSetTheme) TableName() string {
	return "brick_set_theme"
}

type BrickSetType struct {
	BaseModel
	EnglishName string `json:"englishName" gorm:"column:english_name"`
	ChineseName string `json:"chineseName" gorm:"column:chinese_name"`
}

func (*BrickSetType) TableName() string {
	return "brick_set_type"
}
