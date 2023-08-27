package plugins

import "nave/tools/log"

// Exec 执行操作
func Exec(path string, params []string) {
	// ----------------------------------------------
	// TODO : 此处临时测试使用，尽快改成真正的动态调用方法
	// ----------------------------------------------
	switch path {
	case "Print":
		println(params[0])
		break
	case "Info":
		log.Info(params[0])
	case "Warn":
		log.Warn(params[0])
	case "Err":
		log.Err(params[0], nil)
	case "SQL":
		log.Info("MySQL")
	default:
		log.Warn("Cannot find plugins " + path)
	}
}
