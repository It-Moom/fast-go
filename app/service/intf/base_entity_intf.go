/*
 * @PackageName: intf
 * @Description: 基础实体类接口
 * @Author: gabbymrh
 * @Date: 2022-06-24 09:18:37
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-24 09:18:37
 */

package intf

// BaseEntityIntf 基础实体类接口
type BaseEntityIntf interface {
	// FindById 根据 id 查找实体数据
	FindById(id int64) error
	// GetAll 获取所有实体数据
	GetAll() ([]interface{}, error)
	// UpdateById 更新实体数据
	UpdateById(id int64, data interface{}) error
	// DeleteById 删除实体数据
	DeleteById(id int64) error
}
