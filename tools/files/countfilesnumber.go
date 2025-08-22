package files

import (
	"fmt"
	"io/ioutil"
	"time"
)

var match int
var workCount int
var maxWorkCount int
var searchChan chan string
var foundFileChan chan bool
var workDoneChan chan bool

func init() {
	searchChan = make(chan string)
	foundFileChan = make(chan bool)
	workDoneChan = make(chan bool)
	maxWorkCount = 32
	workCount = 0
}
func CountNumberOfFiles(searchFileName string, path string, isMaster bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			filename := file.Name()
			if filename == searchFileName {
				foundFileChan <- true
			}
			if file.IsDir() {
				if workCount < maxWorkCount {
					searchChan <- path + filename + "\\"
				} else {
					CountNumberOfFiles(searchFileName, path+filename+"\\", false)
				}
			}
		}
	}
	if isMaster {
		//走到这里说明已经结束当前协程的递归搜索
		workDoneChan <- true
	}
}
func CountFilesNumber(filename string, path string) int {
	start := time.Now()
	workCount++
	go CountNumberOfFiles(filename, path, true)
	for {
		select {
		case subpath := <-searchChan:
			workCount++
			go CountNumberOfFiles(filename, subpath, true)
		case <-foundFileChan:
			//fmt.Println(match)
			match++
		case <-workDoneChan:
			workCount--
		default:
		}
		if workCount == 0 {
			break
		}
	}
	//关闭管道
	close(searchChan)
	close(foundFileChan)
	close(workDoneChan)
	fmt.Println(match, "match")
	fmt.Println(time.Since(start))
	return match
}
