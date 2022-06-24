/*
 * @PackageName: impl
 * @Description:
 * @Author: gabbymrh
 * @Date: 2022-06-24 09:46:54
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-24 09:46:54
 */

package impl

type UserImpl struct {
}

// FindById 根据 id 查找实体数据
func (ui *UserImpl) FindById(id int64) error {
	return nil
}

// GetAll 获取所有实体数据
func (ui *UserImpl) GetAll() ([]interface{}, error) {
	return nil, nil
}

// UpdateById 更新实体数据
func (ui *UserImpl) UpdateById(id int64, data interface{}) error {
	return nil
}

// DeleteById 删除实体数据
func (ui *UserImpl) DeleteById(id int64) error {
	return nil
}
