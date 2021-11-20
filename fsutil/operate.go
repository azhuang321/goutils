package fsutil

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Mkdir alias of os.MkdirAll()
func Mkdir(dirPath string, perm os.FileMode) error {
	return os.MkdirAll(dirPath, perm)
}

// MkParentDir quick create parent dir
func MkParentDir(fpath string) error {
	dirPath := filepath.Dir(fpath)
	if !IsDir(dirPath) {
		return os.MkdirAll(dirPath, 0775)
	}
	return nil
}

// MustReadFile read file contents, will panic on error
func MustReadFile(filePath string) []byte {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return bs
}

// ReadExistFile read file contents if exist, will panic on error
func ReadExistFile(filePath string) []byte {
	if IsFile(filePath) {
		bs, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		return bs
	}
	return nil
}

// ************************************************************
//	open/create files
// ************************************************************

// OpenFile like os.OpenFile, but will auto create dir.
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)

	// if err := os.Mkdir(dir, 0775); err != nil {
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// QuickOpenFile like os.OpenFile
func QuickOpenFile(filepath string) (*os.File, error) {
	return OpenFile(filepath, DefaultFileFlags, DefaultFilePerm)
}

/* TODO MustOpenFile() */

// CreateFile create file if not exists
// Usage:
// 	CreateFile("path/to/file.txt", 0664, 0666)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode) (*os.File, error) {
	dirPath := path.Dir(fpath)
	if !IsDir(dirPath) {
		err := os.MkdirAll(dirPath, dirPerm)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
}

// MustCreateFile create file, will panic on error
func MustCreateFile(filePath string, filePerm, dirPerm os.FileMode) *os.File {
	file, err := CreateFile(filePath, filePerm, dirPerm)
	if err != nil {
		panic(err)
	}

	return file
}

// ************************************************************
//	copy files
// ************************************************************

// CopyFile copy file to another path.
func CopyFile(src string, dst string) error {
	return errors.New("TODO")
}

// MustCopyFile copy file to another path.
func MustCopyFile(src string, dst string) {
	panic("TODO")
}

// ************************************************************
//	remove files
// ************************************************************

// alias methods
var (
	MustRm  = MustRemove
	QuietRm = QuietRemove
)

// MustRemove removes the named file or (empty) directory.
// NOTICE: if error will panic
func MustRemove(fpath string) {
	if err := os.Remove(fpath); err != nil {
		panic(err)
	}
}

// QuietRemove removes the named file or (empty) directory.
// NOTICE: will ignore error
func QuietRemove(fpath string) {
	_ = os.Remove(fpath)
}

// DeleteIfExist removes the named file or (empty) directory on exists.
func DeleteIfExist(fpath string) error {
	if !PathExists(fpath) {
		return nil
	}

	return os.Remove(fpath)
}

// DeleteIfFileExist removes the named file on exists.
func DeleteIfFileExist(fpath string) error {
	if !IsFile(fpath) {
		return nil
	}

	return os.Remove(fpath)
}

// ************************************************************
//	other operates
// ************************************************************

// Unzip a zip archive
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func Unzip(archive, targetDir string) (err error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return
	}

	if err = os.MkdirAll(targetDir, DefaultDirPerm); err != nil {
		return
	}

	for _, file := range reader.File {
		fullPath := filepath.Join(targetDir, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(fullPath, file.Mode())
			if err != nil {
				return err
			}

			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fileReader.Close()
			return err
		}

		_, err = io.Copy(targetFile, fileReader)

		// close all
		fileReader.Close()
		targetFile.Close()

		if err != nil {
			return err
		}
	}

	return
}

// GetCurrentDirectory 获得当前绝对路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

// LeftAddPathPos 检测并补全路径左边的反斜杠
func LeftAddPathPos(path string) string {
	if path[:0] != "/" {
		path = "/" + path
	}
	return path
}

// RightAddPathPos 检测并补全路径右边的反斜杠
func RightAddPathPos(path string) string {
	if path[len(path)-1:len(path)] != "/" {
		path = path + "/"
	}
	return path
}

// FileNameByDate 根据当天日期和给定dir返回log文件名路径
func FileNameByDate(dir string) string {
	fileName := time.Now().Format("2006-01-02")
	dir = RightAddPathPos(dir)
	return dir + fileName + ".log"
}

// CreateDir 不存在则创建目录
func CreateDir(folderPath string) {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
}

// CreateDateDir 根据当前日期，不存在则创建目录
func CreateDateDir(path string, prex string) string {
	folderName := time.Now().Format("20060102")
	if prex != "" {
		folderName = prex + folderName
	}
	folderPath := filepath.Join(path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// GetDateDir 根据路径获得按照"年月日"生成的子路径
func GetDateDir(path string) string {
	return path + time.Now().Format("20660102")
}

// CreateHourLogFile 根据当前小时创建目录和日志文件
func CreateHourLogFile(path string, prex string) string {
	folderName := time.Now().Format("2006010215")
	if prex != "" {
		folderName = prex + folderName
	}
	folderPath := filepath.Join(path, folderName)
	fmt.Println(folderPath)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, 0777) //0777也可以os.ModePerm
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// ReadDirAll 读取目录
// example ReadDirAll("/Users/why/Desktop/go/test", 0)
func ReadDirAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			ReadDirAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
