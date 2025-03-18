package m_internal

type User struct {
	Created_At uint   `gorm:"not null;" json:"created_at"`
	Updated_At uint   `gorm:"not null;" json:"updated_at"`
	ID         uint   `gorm:"primarykey;" json:"id"`
	Email      string `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Fullname   string `gorm:"size:255;not null;" json:"fullname"`
	// Username   string  	`gorm:"size:30;not null;uniqueIndex" json:"username"`
	Password  *string `gorm:"not null;" json:"password"`
	Position  string  `gorm:"size:50;not null" json:"position"`
	Api_Token string  `gorm:"null" json:"api_token"`
}

func (User) TableName() string {
	return "users"
}
