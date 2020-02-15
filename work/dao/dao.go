package dao

import (
	kinit "gorm_generate/initialize"

	jgorm "github.com/jinzhu/gorm"
)

func DropTable(tx *jgorm.DB, tablename string) error {
	if tx == nil {
		tx = kinit.Gorm
	}
	if err := tx.Exec("DROP TABLE IF EXISTS `" + tablename + "`;").Error; err != nil {
		kinit.LogError.Println(err)
		return err
	}
	return nil
}
func Create(tx *jgorm.DB, createsql string) error {
	if tx == nil {
		tx = kinit.Gorm
	}

	if err := tx.Exec(createsql).Error; err != nil {
		kinit.LogError.Println(err)
		return err
	}
	return nil
}
