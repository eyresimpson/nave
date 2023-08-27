package analysis

import (
	"nave/analysis/jsonResolver"
	"nave/tools/log"
	"nave/types/basic"
	"os"
	"path/filepath"
	"strings"
)

// LoadProject 加载Nave Project (*.basic)
func LoadProject() {

}

// LoadExecFiles 加载指定的Exec文件（Business业务流）
func LoadExecFiles(pn string, arg ...string) []basic.BluePrint {
	var bluePrints []basic.BluePrint
	if pn == "auto" {
		// 获取当前工作目录
		currentDir, err := os.Getwd()
		if err != nil {
			log.Err("[ERROR E001] Cannot getting current directory", err)
			return nil
		}

		// 构建 basic 目录的路径
		execDir := filepath.Join(currentDir)

		// 遍历 basic 目录下的所有 JSON 文件
		err = filepath.Walk(execDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Err("[ERROR] Cannot walking through directory:", err)
				return err
			}

			// 如果是文件且扩展名为 .json
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				// 调用解析器进行解析
				bluePrints = append(bluePrints, jsonResolver.Resolver(path))
				log.Info("Success Load & Resolver BluePrint File: " + path)
			}

			return nil
		})

		if err != nil {
			log.Err("[ERROR] Cannot walking through directory:", err)
		}
	} else if len(arg) >= 1 {
		// 仅加载目标目录文件
	}
	return bluePrints
}

// LoadDBExec 加载数据表中可执行数据项
func LoadDBExec() {

}
