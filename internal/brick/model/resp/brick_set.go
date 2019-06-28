package resp

type ListSetResp struct {
	ID                 string                `json:"id"`
	CreatedAt          string                `json:"createdAt"`
	UpdatedAt          string                `json:"updatedAt"`
	Number             string                `json:"number"`
	EnglishName        string                `json:"englishName"`
	ChineseName        string                `json:"chineseName"`
	ReleaseDate        string                `json:"releaseDate"`
	PartCount          int                   `json:"partCount"`
	MinifigCount       int                   `json:"minifigCount"`
	ChineseDescription string                `json:"chineseDescription" `
	AgeRange           string                `json:"ageRange"`
	Type               *ListSetTypeResp      `json:"type"`
	Theme              *ListSetThemeResp     `json:"theme"`
	SubTheme           *ListSetSubThemeResp  `json:"subTheme"`
	Brand              *ListSetBrandResp     `json:"brand"`
	Packaging          *ListSetPackagingResp `json:"packaging"`
	Tags               []*ListSetTagResp     `json:"tags"`
	Cover              *ListSetMediaResp     `json:"cover"`
}

type ListSetTypeResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetBrandResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetThemeResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetSubThemeResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetPackagingResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetTagResp struct {
	ID          string `json:"id"`
	ChineseName string `json:"chineseName"`
}

type ListSetMediaResp struct {
	ID       string  `json:"id"`
	Type     int8    `json:"type"`
	Src      string  `json:"src"`
	High     int32   `json:"high"`
	Width    int32   `json:"width"`
	Duration float32 `json:"duration"`
}
