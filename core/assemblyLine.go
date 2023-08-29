package core

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"nave/tools/log"
	"nave/types/basic"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var localBluePrint basic.BluePrint

// AssemblyLine 流水线，简称流
// 可以根据Flow文件，执行具体的流水线
func AssemblyLine(bluePrint basic.BluePrint, wg *sync.WaitGroup) {
	// 赋予全局变量
	localBluePrint = bluePrint
	log.Info("Exec localBluePrint " + localBluePrint.Label + " start")

	// 分析需要的插件，并尝试从Plugins中读取
	for _, mod := range bluePrint.Mods {
		// 如果Plugins中不能满足需求，尝试向公共版本库中索引下载
		//os.DirFS()
		// 尝试加载插件
		//plugins.Load(mod)
		println(mod)
	}

	// 检测流水线是否开启了端口监听
	if localBluePrint.FlowType == "service" && localBluePrint.Port != "" {

		router := mux.NewRouter()
		for _, route := range localBluePrint.Route {
			router.HandleFunc("/"+localBluePrint.Path+"/"+route.Path, handler).Methods(route.Type)
		}

		server := &http.Server{
			Addr:    ":" + localBluePrint.Port,
			Handler: router,
		}

		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			log.Success("AssemblyLine start listening for port " + localBluePrint.Port + " for " + localBluePrint.Label)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Err("Cannot starting HTTP server", err)
			}
		}()

		<-stopChan
		log.Info("Received termination signal. Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Err("Error shutting down HTTP server", err)
		} else {
			log.Success("HTTP server has been shut down")
		}
	} else if localBluePrint.FlowType == "crontab" {
		// 定时任务
	} else {
		log.Warn("This blueprint looks blank, Perhaps forgotten what to do?")
	}
	log.Success("Exec assemblyLine " + localBluePrint.Label + " complete")
	// 声明协程处理完毕
	wg.Done()
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("AssemblyLine request accepted")
	// 开始执行操作器
	Run(r.URL.Path, localBluePrint.Steps, w, r)
	log.Success("AssemblyLine request complete")
}
