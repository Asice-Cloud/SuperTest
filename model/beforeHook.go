package model

import (
	"Chat/pkg"
	"errors"
	"gorm.io/gorm"
)

/*
	BeforeUpdate

乐观锁的实现是通过 BeforeUpdate 钩子和 UUID 字段来实现的。在每次更新 UserBasic 模型之前，都会检查当前 UUID 是否和数据库中的 UUID 一致，
如果一致，那么就更新 UUID 为新的 UUID，然后进行更新；
如果不一致，那么就返回一个错误，这将导致 GORM 回滚事务，从而保证数据的一致性
*/
func (u *UserBasic) BeforeUpdate(tx *gorm.DB) (err error) {
	var exist UserBasic
	if err = tx.Model(u).Where("id = ?", u.ID).First(&exist).Error; err != nil {
		return err
	}

	if u.UUID != exist.UUID {
		return errors.New("data has been updated by another transaction")
	}

	snowflake, _ := pkg.NewSnowflake(1, 1)
	newUUID, _ := snowflake.NextID()
	u.UUID = newUUID

	return nil
}
