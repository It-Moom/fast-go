/*
 * @PackageName: entity
 * @Description: 模型基础结构体
 * @Author: gabbymrh
 * @Date: 2021-10-03 16:24:26
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 16:24:26
 */

package entity

import (
	"gorm.io/gorm"
	"time"
)

// BaseEntity 基础实体
type BaseEntity struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:主键ID"`
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间"`
}
