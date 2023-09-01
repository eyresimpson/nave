package flow

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"nave/core/worker"
	"nave/tools/log"
	"nave/types/blueprint"
	"nave/types/variables"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Serve(localBluePrint blueprint.BluePrint, vars *variables.Variables) {
	router := mux.NewRouter()
	for _, route := range localBluePrint.Route {
		router.HandleFunc("/"+localBluePrint.Path+"/"+route.Path, func(w http.ResponseWriter, r *http.Request) {
			log.Info("AssemblyLine " + localBluePrint.Label + " request accepted")
			// 开始执行操作器
			worker.Run(r.URL.Path, localBluePrint.Steps, w, r, vars)
			log.Success("AssemblyLine " + localBluePrint.Label + " request complete")
		}).Methods(route.Type)
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
}
