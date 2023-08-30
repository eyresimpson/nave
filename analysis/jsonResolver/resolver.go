package jsonResolver

import (
	"encoding/json"
	"io/ioutil"
	"nave/tools/log"
	"nave/types/blueprint"
	"os"
	"strings"
)

// Resolver 将Json文件解析为业务流
func Resolver(filePath string) blueprint.BluePrint {
	var flow blueprint.BluePrint
	// 读取Json中的配置
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Err("Cannot Open File", err)
		return flow
	}
	//log.Info("Success Load BluePrint BluePrint File", nil)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &flow)
	// 忽略字段不匹配的错误
	if err != nil && !strings.Contains(err.Error(), "cannot unmarshal object into Go struct field") {
		log.Err("Cannot Resolver BluePrint BluePrint File", err)
		return flow
	}
	return flow
}
