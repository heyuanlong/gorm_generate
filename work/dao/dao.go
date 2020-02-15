package dao

import (
	kinit "gorm_generate/initialize"

	jgorm "github.com/jinzhu/gorm"
)

func Create(tx *jgorm.DB, dropsql string, createsql string) {
	if tx == nil {
		tx = kinit.Gorm
	}

	if err := tx.Exec(dropsql).Error; err != nil {
		kinit.LogError.Println(err)
		return
	}
	if err := tx.Exec(createsql).Error; err != nil {
		kinit.LogError.Println(err)
		return
	}
	return
}
