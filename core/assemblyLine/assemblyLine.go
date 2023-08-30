package assemblyLine

import (
	"nave/core/assemblyLine/flow"
	"nave/tools/log"
	"nave/types/blueprint"
	"nave/types/variables"
	"sync"
)

// 全局变量，全部流水线都可访问，后续做了并发处理后再具体使用
var svars = variables.Variables{String: map[string]string{}, Int: map[string]int{}, Bool: map[string]bool{}}

// AssemblyLine 流水线，简称流
// 可以根据Flow文件，执行具体的流水线
func AssemblyLine(bluePrint blueprint.BluePrint, wg *sync.WaitGroup) {
	// 当前蓝图
	var localBluePrint blueprint.BluePrint
	// 流水线变量，只有流水线内部可以访问
	var vars = variables.Variables{String: map[string]string{}, Int: map[string]int{}, Bool: map[string]bool{}}

	// 赋予全局变量
	// TODO： 临时实现，后续优化
	localBluePrint = bluePrint
	log.Info("BluePrint convert to AssemblyLine " + localBluePrint.Label + " success")

	// 分析需要的插件，并尝试从Plugins中读取
	for _, mod := range localBluePrint.Mods {
		// 如果Plugins中不能满足需求，尝试向公共版本库中索引下载
		// TODO：尝试加载插件（优先保证Windows，首选GRPC方式）
		//plugins.Load(mod)
		println(mod)
	}

	// 实例化变量
	for _, variable := range localBluePrint.Vars {
		// 一定要注意默认值，默认不允许声明没有默认值的变量，如果蓝图中出现没有默认值的变量，按下述规则赋值
		if variable.Default == nil {
			// 初始默认值
			switch variable.Type {
			case "String":
				vars.String[variable.Name] = ""
			case "Int":
				vars.Int[variable.Name] = 0
			case "Bool":
				vars.Bool[variable.Name] = false
			}
			log.Warn("Cannot find var default value, Set " + variable.Name + " default")
		} else {
			switch variable.Type {
			case "String":
				vars.String[variable.Name] = variable.Default.(string)
			case "Int":
				vars.Int[variable.Name] = variable.Default.(int)
			case "Bool":
				vars.Bool[variable.Name] = variable.Default.(bool)
			}
		}

	}
	println(vars.String["someone"])
	// 检测流水线是否开启了端口监听
	if localBluePrint.FlowType == "service" && localBluePrint.Port != "" {
		flow.Serve(localBluePrint)
	} else if localBluePrint.FlowType == "crontab" {
		// 定时任务
		flow.Cron()
	} else {
		log.Warn("This blueprint looks blank, Perhaps forgotten what to do?")
	}
	log.Success("Exec assemblyLine " + localBluePrint.Label + " complete")
	// 声明协程处理完毕
	wg.Done()
}
