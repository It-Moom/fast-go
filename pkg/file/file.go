/*
 * @PackageName: file
 * @Description: 文件处理
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:52:49
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:52:49
 */

package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// FileNameWithoutExtension 获取文件名，不包含扩展名
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
