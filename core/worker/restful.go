package worker

import (
	"nave/plugins"
	"nave/types/blueprint"
	"net/http"
	"strings"
)

var xw http.ResponseWriter
var xr *http.Request

func RestfulWorker(opt blueprint.Worker, w http.ResponseWriter, r *http.Request) {
	// 提供给执行操作的参数
	//var params []string
	// 全局变量赋值（其他地方可能用到），如果不是Restful，这两个字段就为nil
	xw = w
	xr = r
	// 执行参数表达式转换
	opt.OptParams = convertParamsExp(opt)
	// 调用具体实现
	// 理论上，所有的功能都是插件功能，用户的所有可执行操作都由插件定义
	plugins.Exec(opt.Name, opt.OptParams)
}

func convertParamsExp(opt blueprint.Worker) []string {
	var params []string
	for index, param := range opt.OptParams {

		// 转换参数表达式（#{}）
		if strings.Contains(param, "#{") {
			exp := strings.Split(param, "#{")
			text := param
			for _, s := range exp {
				if len(s) != 0 && strings.Contains(s, "}") {
					s2 := strings.Split(s, "}")[0]
					// TODO: 后续支持路径Path语法，并检测是否报错
					nParam := strings.Replace(text, "#{"+s2+"}", xr.FormValue(s2), -1)
					text = nParam
				}
			}
			params = append(params, text)
		} else {
			params = append(params, opt.OptParams[index])
		}
	}
	return params
}
