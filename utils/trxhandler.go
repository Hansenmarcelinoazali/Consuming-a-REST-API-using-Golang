package utils

import (
	"external_api/model"

	"github.com/jinzhu/gorm"
)

func RollbackandCommitHandler(tx *gorm.DB, data *model.Product) error {
	err := tx.Create(data).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
