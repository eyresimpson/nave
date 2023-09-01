package worker

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"nave/types/blueprint"
	"nave/types/variables"
	"strings"
)

func ConditionWorker(opt blueprint.Worker, vars *variables.Variables) {
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
			Run(item, opt, xw, xr, vars)
		}
	}
}

// ConditionExp 取出表达式
func ConditionExp(expression string) bool {
	// 取出表达式部分
	expr := strings.Split(strings.Split(expression, "${")[1], "}")[0]

	// 解析字符串表达式
	result, err := EvaluateExpression(expr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return result.(bool)
}

// EvaluateExpression 执行表达式
func EvaluateExpression(expression string) (interface{}, error) {
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
