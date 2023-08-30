package core

import (
	"nave/analysis"
	"nave/core/assemblyLine"
	"nave/mods"
	"nave/tools/log"
	"sync"
)

// Engine 核心控制器
/*
	优先使用用户配置文件（运行目录下的app.yml），如果当前目录下没有检测到配置文件，将使用默认配置
	默认配置：
		- 默认日志存放在运行目录的logs下
		- 默认执行运行目录和运行目录下exec目录中的所有符合规范的json、xml、yml文件
		- 默认以文件名称顺序执行，从小到大，以Json - Yml - Xml顺序执行
		- 默认开启高安全性，不开放系统API
		- 默认加载系统默认的操作单元
	如果未检测到任何可执行exec文件，将退出运行
*/
func Engine() {
	log.Info("Engine Running...")
	// 协程等待
	var wg sync.WaitGroup
	// 读取配置文件（同步）
	log.Info("Engine Load Config...")

	// 运行监听模块（如果需要）（异步）
	log.Info("Engine Skip Listener...")
	//handle.Run()

	// 加载可执行Exec文件/项目
	log.Info("Engine Load Blueprint...")
	bluePrints := analysis.LoadExecFiles("auto")

	// 分析需要的插件，并尝试从Plugins中读取
	log.Info("Engine Load Mods...")
	mods.Load("mod")

	log.Info("Engine Start Flows...")
	// 尝试根据执行计划执行（for）
	for _, bluePrint := range bluePrints {
		wg.Add(1)
		// 通过协程执行
		go assemblyLine.AssemblyLine(bluePrint, &wg)
	}
	wg.Wait()
	log.Info("Engine Stopped")
}
