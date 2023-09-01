package worker

import (
	"nave/tools/log"
	"nave/types/blueprint"
	"nave/types/variables"
	"net/http"
	"strings"
)

// Run Work编排
// TODO： 设计上这部分应该是流水线的工作（Work分配），后续整理代码时注意
func Run(sid string, parent interface{}, w http.ResponseWriter, r *http.Request, vars *variables.Variables) {
	switch parent.(type) {
	case blueprint.Worker:
		parents := parent.(blueprint.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents.Children {
			if child.Sid == sid {
				if child.Type == 0 {
					ConditionWorker(child, vars)
				} else if child.Type == 1 {
					RestfulWorker(child, w, r, vars)
				} else if child.Type == 2 {

				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r, vars)
				} else {
					break
				}
			}
		}
	case []blueprint.Worker:
		parents := parent.([]blueprint.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents {
			if child.Sid == sid {
				if child.Type == 0 {
					ConditionWorker(child, vars)
				} else {
					RestfulWorker(child, w, r, vars)
				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r, vars)
				} else {
					break
				}
			}
		}
	default:
		log.Warn("Received an unknown type")
	}

}

func convertVariableExp(exps string, vars *variables.Variables) string {
	//println(vars.String["someone"])
	if strings.Contains(exps, "@{") {
		exp := strings.Split(exps, "@{")
		text := exps
		for _, s := range exp {
			if len(s) != 0 && strings.Contains(s, "}") {
				s2 := strings.Split(s, "}")[0]
				// TODO: 后续支持路径Path语法，并检测是否报错
				nParam := strings.Replace(text, "@{"+s2+"}", vars.String[s2], -1)
				text = nParam
			}
		}
		return text
	}
	return ""
	//var params []string
	//for index, param := range opt.OptParams {
	//
	//	// 转换参数表达式（#{}）
	//	if strings.Contains(param, "#{") {
	//		exp := strings.Split(param, "#{")
	//		text := param
	//		for _, s := range exp {
	//			if len(s) != 0 && strings.Contains(s, "}") {
	//				s2 := strings.Split(s, "}")[0]
	//				// TODO: 后续支持路径Path语法，并检测是否报错
	//				nParam := strings.Replace(text, "#{"+s2+"}", xr.FormValue(s2), -1)
	//				text = nParam
	//			}
	//		}
	//		params = append(params, text)
	//	} else {
	//		params = append(params, opt.OptParams[index])
	//	}
	//}
	//return params
}
