package files

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 判断文件夹是否存在
func IsPathExist(folderPath string) bool {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 文件夹不存在
		fmt.Println("文件夹不存在")
		return false
	} else {
		// 文件夹存在
		fmt.Println("文件夹存在")
		return true
	}
}
func IsExist(folderPath string) bool {
	f, err := os.Open(folderPath)
	defer f.Close()
	if err != nil {
		return false
	} else {
		return f != nil
	}
}
func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true, nil
	}
	return false, err
}
func IsFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err // 如果Stat函数返回错误，则路径可能不存在或者无法访问
	}
	// 使用Mode().IsRegular()方法判断是否为文件
	return fileInfo.Mode().IsRegular(), nil
}
func Create(path string) (*os.File, error) {
	exist := IsExist(path)
	if exist {
		return os.Create(path)
	}
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(path)
}

// CreateFile creates a file specified by path.
// If the file already exists, it is truncated.
// If the parent directory does not exist, it will be created with mode os.ModePerm(0777).
func CreateFile(path string) error {
	file, err := Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func ReadLinesV2(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			lines = append(lines, line)
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, line[:len(line)-1])
	}
	return lines, nil
}
func ReadLinesV3(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var lines []string
	r := bufio.NewReader(f)
	for {
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return lines, err
		}
		lines = append(lines, string(bytes))
	}
	return lines, nil
}
func ListDir(dir string) ([]string, error) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}
func ListFilenames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var names []string
	for i := 0; i < len(entries); i++ {
		if !entries[i].IsDir() {
			names = append(names, entries[i].Name())
		}
	}
	return names, nil
}
func FileMd5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
func FileSize(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
