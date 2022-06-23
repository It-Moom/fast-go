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
	"fast-go/pkg/app"
	"fast-go/pkg/auth"
	"fast-go/pkg/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
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

// SaveUploadFile 保存上传的头像
func SaveUploadFile(c *gin.Context, file *multipart.FileHeader) (string, error) {

	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/%s/%s/", app.TimenowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/2021/12/22/1/nFDacgaWKpWWOmOt.png
	filePath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return avatar, err
	}

	return filePath, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
