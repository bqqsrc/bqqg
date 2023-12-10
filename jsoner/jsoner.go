package jsoner

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

type ConfigReader interface {
	Bytes2Config(string, []byte) error
}

func ReadAllConfig(serverConfig []string, configReader ConfigReader) error {
	//funcName := "readListenConf"
	//log.Debugf("%s, serverConfig is %s", funcName, serverConfig)
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd error ,err is: %s", err)
	}
	for _, value := range serverConfig {
		dir := value
		if _, err := os.Lstat(dir); os.IsNotExist(err) { //如果路径不存在，将路径拼接到当前路径，再看看是否存在
			dir = path.Join(wd, dir)
			if _, err = os.Lstat(dir); os.IsNotExist(err) {
				log.Printf("warning: path not exist %s", value)
				continue
			}
		}
		suffix := ".json"
		//log.Debugf("%s, dir is %s, suffix is %s", funcName, dir, suffix)
		if file, err := os.Stat(dir); file.IsDir() {
			//log.Debugf("%s, domains is before walkConfigDir %s", funcName, d)
			if err = ReadConfigDir(dir, suffix, configReader); err != nil {
				return err
			}
			//log.Debugf("%s, domains is after walkConfigDir %s", funcName, d)
		} else {
			//log.Debugf("%s, isFile", funcName)
			if strings.HasSuffix(strings.ToLower(dir), suffix) { //只遍历后缀为json的文件
				//log.Debugf("%s, domains is before %s", funcName, d)
				if err = ReadFileConfig(dir, configReader); err != nil {
					return err
				}
				//log.Debugf("%s, domains is after file2ListenArr %s", funcName, d)
			} else {
				//log.Warnf("%s warn: %s is not end with %s", funcName, value, suffix)
			}
		}
	}
	return nil

}

//遍历一个目录下的配置文件
func ReadConfigDir(dir, suffix string, configReader ConfigReader) error {
	//funcName := "walkConfigDir"
	dirInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ioutil.ReadDir(%s) error, err is: %s", dir, err)
	}
	for _, fileInfo := range dirInfo {
		newPath := fileInfo.Name()
		newPath = path.Join(dir, newPath)
		if fileInfo.IsDir() {
			if err = ReadConfigDir(newPath, suffix, configReader); err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(strings.ToLower(newPath), suffix) {
				//log.Debugf("%s, before walkConfigDir, domains is %s", funcName, d)
				if err = ReadFileConfig(newPath, configReader); err != nil {
					return err
				}
				//log.Debugf("%s, after walkConfigDir, domains is %s", funcName, d)
			} else {
				log.Printf("info: %s is not end with %s", newPath, suffix)
			}
		}
	}
	return nil
}

//将一个配置文件转换为一组监听列表
func ReadFileConfig(file string, configReader ConfigReader) error {
	//funcName := "file2ListenArr"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("ioutil.ReadFile(%s) error, err is: %s", file, err)
	}
	return configReader.Bytes2Config(file, content)
}
