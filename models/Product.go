package models

type Product struct {
	Id               int64  `gorm:"primaryKey" json:"id"`
	UUID             string `gorm:"type:uuid;default:gen_random_uuid()"`
	ProductName      string `gorm:"type:varchar(300)" json:"productName"`
	ProductPrice     int64  `gorm:"type:integer" json:"productPrice"`
	ProductImageName string `gorm:"type:text" json:"imageName"`
	ProductImageURL  string `gorm:"type:text" json:"imageURL"`
}
