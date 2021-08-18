package entity

type Product struct {
	Id          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(191)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Price       uint   `gorm:"type:double" json:"price"`
	ImagePath   string `gorm:"type:varchar(191)" json:"image_path" `
}
