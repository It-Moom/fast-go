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
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

// BaseEntity 基础实体
type BaseEntity struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time      `gorm:"column:created_at;index;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;comment:更新时间" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (be BaseEntity) GetStringID() string {
	return cast.ToString(be.ID)
}
