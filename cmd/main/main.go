package main

import (
	"advanced-rest-yt/internal/config"
	"advanced-rest-yt/internal/user/db"
	"advanced-rest-yt/pkg/client/mongodb"
	"advanced-rest-yt/pkg/logging"
	"context"
	"errors"
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

func init() {
	// if using socket
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	err := os.Remove(p + "/app.socket")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Err is: %v", err)
		}
	}
}

func main() {
	ctx := context.Background()
	logger := logging.GetLogger()

	router := httprouter.New()
	logger.Info("create router success!")

	cfg := config.GetConfig()

	database, err := mongodb.NewClient(
		ctx,
		cfg.Storage.MongoDB.Host,
		cfg.Storage.MongoDB.Port,
		cfg.Storage.MongoDB.Username,
		cfg.Storage.MongoDB.Password,
		cfg.Storage.MongoDB.Database,
		cfg.Storage.MongoDB.AuthDB,
		logger,
	)
	if err != nil {
		logger.Fatalf("cant create client, error: %v", err)
	}

	storage := db.NewStorage(database, cfg.Storage.MongoDB.Collection, logger)

	testDatabase(ctx, storage, logger)

	userHandler := user.NewHandler(logger)
	userHandler.Register(router)

	start(router, logger, cfg)
}

func testDatabase(ctx context.Context, storage user.Storage, logger *logging.Logger) {
	id, err := storage.Create(ctx, *createTestUser())
	if err != nil {
		panic(err)
	}
	logger.Infof("created user, id: %s", id)

	foundedUser, err := storage.FindOne(ctx, id)
	if err != nil {
		panic(err)
	}
	logger.Infof("finded user is: %v", foundedUser)

	foundedUser.Username = "bibki!"
	err = storage.Update(ctx, foundedUser)
	if err != nil {
		panic(err)
	}
	logger.Infof("user was updated")

	//t, _ := context.WithTimeout(ctx, time.Second*1)
	//err = storage.Delete(t, foundedUser.ID)
	//_ = storage.Delete(t, "652699222e181c4337fe888a")
	//if err != nil {
	//	panic(err)
	//}

	res, err := storage.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Debugf("All users: %s", res)
}

func createTestUser() *user.User {
	return &user.User{
		ID:           "",
		Email:        "mail@mail.com",
		Username:     "fids",
		PasswordHash: "r23c",
	}
}

func start(r *httprouter.Router, logger *logging.Logger, cfg *config.Config) {

	var listener net.Listener
	var listenErr error
	var msg string

	if cfg.Listen.Type == "socket" {
		logger.Info("start detect app path to socket")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal("cant find filepath", err)
		}
		logger.Debugf("path is %s", appDir)

		logger.Info("start create socket")
		socketPath := filepath.Join(appDir, "app.socket")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("start create unix socket listener")
		listener, listenErr = net.Listen("unix", socketPath)
		if listenErr != nil {
			logger.Fatalf("cant start socket: %s", listenErr)
		}
		msg = fmt.Sprintf("created socket, listener is: %s", socketPath)
	} else {
		logger.Info("create tcp listener")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		msg = fmt.Sprintf("Server started! on %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
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

	logger.Infof(msg)
	err := server.Serve(listener)
	if err != nil {
		logger.Fatalln("Server was closed!", err)
	}

}
