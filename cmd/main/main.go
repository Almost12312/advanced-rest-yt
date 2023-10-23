package main

import (
	"advanced-rest-yt/internal/author"
	authRepo "advanced-rest-yt/internal/author/db/postgresql"
	"advanced-rest-yt/internal/author/model"
	"advanced-rest-yt/internal/author/service"
	"advanced-rest-yt/internal/author/storage"
	bookDB "advanced-rest-yt/internal/book"
	book "advanced-rest-yt/internal/book/db"
	"advanced-rest-yt/internal/config"
	"advanced-rest-yt/pkg/api/sort"
	"advanced-rest-yt/pkg/client/postgresql"
	"advanced-rest-yt/pkg/logging"
	"advanced-rest-yt/pkg/strs"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
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

	//database, err := mongodb.NewClient(
	//	ctx,
	//	cfg.Storage.MongoDB.Host,
	//	cfg.Storage.MongoDB.Port,
	//	cfg.Storage.MongoDB.Username,
	//	cfg.Storage.MongoDB.Password,
	//	cfg.Storage.MongoDB.Database,
	//	cfg.Storage.MongoDB.AuthDB,
	//	logger,
	//)
	//if err != nil {
	//	logger.Fatalf("cant create client, error: %v", err)
	//}

	//mongo := db.NewStorage(database, cfg.Storage.MongoDB.Collection, logger)
	postgres, err := postgresql.NewClient(ctx, 3, cfg.Storage.PostgreSQL, logger)
	if err != nil {
		logger.Fatalf("cant create postgres client: %s", err)
	}

	authRepository := authRepo.NewRepository(postgres, logger)
	bookRepository := book.NewRepository(postgres, logger)
	_ = bookRepository
	//testMongoDB(ctx, mongo, logger)
	//testPostgreSQL(ctx, logger, authRepository, bookRepository)

	logger.Info("start creating handlers")
	userHandler := user.NewHandler(logger)
	userHandler.Register(router)

	authorService := service.NewService(authRepository, logger)
	authorHandler := author.NewHandler(logger, authorService)
	authorHandler.Register(router)
	logger.Info("end creating handlers")

	start(router, logger, cfg)
}

func testPostgreSQL(ctx context.Context, logger *logging.Logger, authRepo storage.Repository, bookRepo bookDB.Repository) {
	opt := storage.NewSortOptions("age", sort.ASC)

	authors, err := authRepo.FindAll(ctx, opt)
	if err != nil {
		logger.Fatalf("%s", err)
	}

	for _, a := range authors {
		logger.Infof("%s", a)
	}

	a, err := authRepo.FindOne(ctx, "cb7f6f2b-8663-467f-9c80-d3ea364be7ef")
	if err != nil {
		logger.Errorf("Cant test FindOne(), err: %s", err)
	}

	logger.Debugf("Author is: %s", a)

	ath := &model.Author{
		Name: strs.RandomString(int8(rand.Intn(99))),
	}

	id, err := authRepo.Create(ctx, ath)
	if err != nil {
		logger.Errorf("Cant test Create(), err: %s", err)
	}

	_ = id

	all, err := bookRepo.FindAll(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debugf("Authors is: %v", all)
}

func testMongoDB(ctx context.Context, mongo user.Repository, logger *logging.Logger) {
	id, err := mongo.Create(ctx, *createTestUser())
	if err != nil {
		panic(err)
	}
	logger.Infof("created user, id: %s", id)

	foundedUser, err := mongo.FindOne(ctx, id)
	if err != nil {
		panic(err)
	}
	logger.Infof("finded user is: %v", foundedUser)

	foundedUser.Username = "bibki!"
	err = mongo.Update(ctx, foundedUser)
	if err != nil {
		panic(err)
	}
	logger.Infof("user was updated")

	//t, _ := context.WithTimeout(ctx, time.Second*1)
	//err = mongo.Delete(t, foundedUser.ID)
	//_ = mongo.Delete(t, "652699222e181c4337fe888a")
	//if err != nil {
	//	panic(err)
	//}

	res, err := mongo.FindAll(ctx)
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
