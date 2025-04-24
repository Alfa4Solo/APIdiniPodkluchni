package API_to_DB_Memesapp

type Role struct {
	IDRole   string `gorm:"primaryKey" json:"id_role"`
	NameRole string `json:"name"`
}

type User struct {
	IDUser       string `gorm:"primaryKey" json:"id_user"`
	IDRole       string `gorm:"primaryKey" json:"id_role"`
	NameUser     string `json:"user_name"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type Meme struct {
	IDMeme           string `gorm:"primaryKey" json:"id_meme"`
	NameMeme         string `json:"name_meme"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	IDUser           string `gorm:"primaryKey" json:"id_user"`
}

type TagMeme struct {
	IDTag  string `gorm:"primaryKey" json:"id_tag"`
	IDMeme string `gorm:"primaryKey" json:"id_meme"`
}

type Tag struct {
	IDTag   string `gorm:"primaryKey" json:"id_tag"`
	NameTag string `json:"name_tag"`
}

type CategoryMeme struct {
	IDCategory string `gorm:"primaryKey" json:"id_category"`
	IDMeme     string `gorm:"primaryKey" json:"id_meme"`
}

type Category struct {
	IDCategory   string `gorm:"primaryKey" json:"id_category"`
	NameCategory string `json:"name_category"`
}
