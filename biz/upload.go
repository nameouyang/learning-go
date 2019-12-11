package biz

import (
	"fmt"
	"github.com/nameouyang/learning-go/conf"
	"github.com/nameouyang/learning-go/lib/utils"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type UploadService struct{}

// 获取图片相对目录
func (us *UploadService) GetImgPath() string {
	return conf.ServerConf.StaticRootPath
}

// GetImgFullPath 获取图片完整目录
func (upload *UploadService) GetImgFullPath() string {
	return conf.ServerConf.StaticRootPath + conf.ServerConf.UploadImagePath
}

// GetImgName 获取图片名称
func (us *UploadService) GetImgName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.MakeSha1(fileName)
	return fileName + ext
}

// GetImgFullURL 获取图片完整URL
func (us *UploadService) GetImgFullURL(name string) string {
	return conf.ServerConf.PrefixURL + conf.ServerConf.UploadImagePath + name
}

// CheckImgExt 检查图片后缀是否满足要求
func (us *UploadService) CheckImgExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range conf.ServerConf.ImageFormats {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImgSize 检查图片大小是否超出
func (us *UploadService) CheckImgSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error().Msg(err.Error())
		return false
	}
	// 单位转换 bytes to Megabyte
	const converRatio float64 = 1024 * 1024
	fileSize := float64(len(content)) / converRatio
	// 文件大小不得超出上传限制：5M
	return fileSize <= conf.ServerConf.UploadLimit
}

// CheckImgPath 检测图片路径是否创建及权限是否满足
func (us *UploadService) CheckImgPath(path string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	isExist, err := utils.IsExist(dir + "/" + path)
	if err != nil {
		return fmt.Errorf("utils.IsExist err: %v", err)
	}
	if isExist == false {
		// 若路径不存在，则创建
		err := os.MkdirAll(dir+"/"+path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("os.MkdirAll err: %v", err)
		}
	}
	isPerm := utils.IsPerm(path)
	if isPerm {
		return fmt.Errorf("utils.IsPerm Permission denied src: %s", path)
	}
	return nil
}
