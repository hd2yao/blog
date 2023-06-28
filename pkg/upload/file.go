package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/hd2yao/blog/global"
	"github.com/hd2yao/blog/pkg/util"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileExt 获取文件扩展名
func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

/*
	检查文件的相关方法
*/

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	// func Stat(name string) (fi FileInfo, err error)
	_, err := os.Stat(dst)
	// func IsNotExist(err error) bool
	// 返回一个布尔值说明该错误是否表示一个文件或目录不存在
	return os.IsNotExist(err)
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	// 上传的文件的后缀有可能是大写、小写、大小写,因此统一转为大写
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

/*
	涉及文件写入/创建的相关操作
*/

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 保存所上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	// 打开源地址文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标地址文件
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// 文件内容拷贝
	_, err = io.Copy(out, src)
	return err
}
