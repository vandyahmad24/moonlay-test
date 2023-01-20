package database

import (
	"fmt"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	File        string `gorm:"type:varchar(255)"`
	IsParent    bool
	ParentId    int
}

func InitMigration() {
	fmt.Println("Jalankan migration")
	DB.AutoMigrate(
		&Todo{},
	)
	fmt.Println("Selesai migration")
}
