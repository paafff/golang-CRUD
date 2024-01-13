package models

type Product struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	UUID         string `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	ProductName  string `gorm:"type:varchar(300)" json:"product_name"`
	ProductPrice int64  `gorm:"type:integer" json:"product_price"`
}

// UUID string `gorm:"type:uuid;default:uuid_generate_v4()"`
