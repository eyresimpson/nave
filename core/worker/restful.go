package worker

import (
	"nave/mods"
	"nave/types/blueprint"
	"nave/types/variables"
	"net/http"
	"strings"
)

var xw http.ResponseWriter
var xr *http.Request

func RestfulWorker(opt blueprint.Worker, w http.ResponseWriter, r *http.Request, vars *variables.Variables) {

	// 提供给执行操作的参数
	//var params []string
	// 全局变量赋值（其他地方可能用到），如果不是Restful，这两个字段就为nil
	xw = w
	xr = r
	// 执行参数表达式转换
	opt.OptParams = convertParamsExp(opt, vars)
	// 执行变量表达式转换

	// 调用具体实现
	// 理论上，所有的功能都是插件功能，用户的所有可执行操作都由插件定义
	mods.Exec(opt.Name, opt.OptParams)
}

// TODO：这种垃圾解析方法等我写完Artist，立刻修复
func convertParamsExp(opt blueprint.Worker, vars *variables.Variables) []string {
	var params []string
	var ret []string
	for index, param := range opt.OptParams {

		// 转换参数表达式 #{}
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

	for index, param := range params {
		if strings.Contains(param, "@{") {
			// 处理变量表达式 @{}
			exp := strings.Split(param, "@{")
			text := param
			for _, s := range exp {
				if len(s) != 0 && strings.Contains(s, "}") {
					s2 := strings.Split(s, "}")[0]
					// TODO: 后续支持路径Path语法，并检测是否报错
					nParam := strings.Replace(text, "@{"+s2+"}", vars.String[s2], -1)
					text = nParam
				}
			}
			ret = append(ret, text)
		} else {
			ret = append(ret, params[index])
		}
	}
	return ret
}
