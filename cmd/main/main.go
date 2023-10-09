package main

import (
	"advanced-rest-yt/internal/config"
	"advanced-rest-yt/pkg/logging"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"advanced-rest-yt/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	router := httprouter.New()
	logger.Info("create router success!")

	cfg := config.GetConfig()
	_ = cfg

	userHandler := user.NewHandler(logger)
	userHandler.Register(router)

	start(router, logger, cfg)
}

func start(r *httprouter.Router, logger *logging.Logger, cfg *config.Config) {

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "socket" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal("cant find filepath", err)
		}

		logger.Info("start create socket")
		socketPath := filepath.Join(appDir, "app.socket")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("create unix socket listener")
		listener, listenErr = net.Listen("unix", socketPath)
		if listenErr != nil {
			logger.Fatalf("cant start socket: %s", listenErr)
		}
		logger.Debugf("created socket, listener is: %s", socketPath)
	} else {
		logger.Info("create tcp listener")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	}

	if listenErr != nil {
		log.Fatalln("Cant create logger", listenErr)

	}

	logger.Println("logger success created!")

	server := &http.Server{
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("Server started! on %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	err := server.Serve(listener)
	if err != nil {
		logger.Fatalln("Server was closed!", err)
	}

}
