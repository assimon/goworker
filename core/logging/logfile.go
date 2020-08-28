package logging

import (
	"fmt"
	"goworker/core/config"
	"log"
	"os"
	"time"
)


/**
获取日志保存路径.
*/
func getLogsFilePath() string {
	return fmt.Sprintf("%s%s", config.AppConfig.RuntimeRootPath, config.AppConfig.LogSavePath)
}

/**
获取文件名，根据时间决定文件名.
*/
func getLogsFileName() string {
	return fmt.Sprintf("%s%s.%s",
		config.AppConfig.LogSaveName,
		time.Now().Format(config.AppConfig.LogTimeFormat),
		config.AppConfig.LogFileExt,
	)
}

/**
检测是否具有路径权限.
 */
func CheckPermission(path string) bool {
	_, err := os.Stat(path)
	return os.IsPermission(err)
}

/**
检查目录是否存在.
 */
func CheckNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

/**
创建目录
 */
func MkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

/**
检测是否有目录不存在则创建.
 */
func IsNotExistMkDir(path string) error {
	if notExist := CheckNotExist(path); notExist == true{
		if err := MkDir(path); err != nil {
			return err
		}
	}
	return nil
}

/**
打开文件流.
 */
func Open(filename string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func MustOpen(fileName,  filePath string) (*os.File, error)  {
	dir, err := os.Getwd();
	if err != nil {
		return nil, fmt.Errorf("os.Getwd() err: %v", err)
	}
	src := dir + "/" + filePath
	// 判断有无权限.
	permission := CheckPermission(src)
	if permission == true {
		log.Fatalf("file.CheckPermission Permission denied src: %s", src)
	}
	// 判断有无目录.
	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	// 打开文件流.
	file, err := Open(src + fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return file, nil
}
