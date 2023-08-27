package core

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"nave/plugins"
	"nave/types/basic"
	"net/http"
	"strings"
)

func Run(sid string, parent interface{}, w http.ResponseWriter, r *http.Request) {
	switch parent.(type) {
	case basic.Worker:
		//fmt.Println("Received a string:", v)
		parents := parent.(basic.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents.Children {
			if child.Sid == sid {
				if child.Type == 0 {
					RestfulWorker(child, w, r)
				} else {
					RestfulConditionWorker(child, w, r)
				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r)
				} else {
					break
				}
			}
		}
	case []basic.Worker:
		parents := parent.([]basic.Worker)
		// 这个循环的是需要执行节点和其兄弟节点
		for _, child := range parents {
			if child.Sid == sid {
				if child.Type == 0 {
					RestfulWorker(child, w, r)
				} else {
					RestfulConditionWorker(child, w, r)
				}
				// 如果有后续节点
				if child.Next != "-1" {
					Run(child.Next, parent, w, r)
				} else {
					break
				}
			}
		}
	default:
		fmt.Println("Received an unknown type")
	}

}

func RestfulWorker(opt basic.Worker, w http.ResponseWriter, r *http.Request) {
	var params []string
	// 执行参数表达式转换
	for index, param := range opt.OptParams {

		if strings.Contains(param, "#{") {
			exp := strings.Split(param, "#{")
			text := param
			for _, s := range exp {
				if len(s) != 0 && strings.Contains(s, "}") {
					s2 := strings.Split(s, "}")[0]
					// TODO: 后续支持路径Path语法，并检测是否报错
					nParam := strings.Replace(text, "#{"+s2+"}", r.FormValue(s2), -1)
					text = nParam
				}
			}
			params = append(params, text)
		} else {
			params = append(params, opt.OptParams[index])
		}
	}
	plugins.Exec(opt.Name, params)
}

func RestfulConditionWorker(opt basic.Worker, w http.ResponseWriter, r *http.Request) {
	// 执行判断表达式
	for _, c := range opt.Condition {
		// 执行列表
		var execList []string
		// 切分三元表达式
		Exp := strings.Split(c, "?")
		// 代表表达式前面的值，即判断条件
		CE := strings.Split(Exp[1], ":")
		// 执行条件表达式
		if ConditionExp(Exp[0]) {
			// 执行列表
			execList = strings.Split(CE[0], ",")
		} else {
			// 执行列表
			execList = strings.Split(CE[1], ",")
		}
		// 循环执行列表
		for _, item := range execList {
			Run(item, opt, w, r)
			// 循环子列表，以执行对应SID的Child
			//for _, child := range opt.Children {
			//	if child.Sid == item {
			//		if child.Type == 0 {
			//			RestfulWorker(child, w, r)
			//		} else {
			//			RestfulConditionWorker(child, w, r)
			//		}
			//	}
			//}
		}
	}
}

func ConditionExp(expression string) bool {
	// 取出表达式部分
	expr := strings.Split(strings.Split(expression, "${")[1], "}")[0]

	// 解析字符串表达式
	result, err := evaluateExpression(expr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return result.(bool)
}

func evaluateExpression(expression string) (interface{}, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return nil, err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
