package server

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rsoi-2021-lab1-ci-cd-Kotyarich/common"
	"rsoi-2021-lab1-ci-cd-Kotyarich/db"
	"rsoi-2021-lab1-ci-cd-Kotyarich/user"
	userHttp "rsoi-2021-lab1-ci-cd-Kotyarich/user/delivery/http"
	userPostgres "rsoi-2021-lab1-ci-cd-Kotyarich/user/repository/postgres"
	userUceCase "rsoi-2021-lab1-ci-cd-Kotyarich/user/usecase"
	"time"
)

type App struct {
	httpServer *http.Server

	userUC user.UseCase
}

func NewApp() *App {
	userRepo := userPostgres.NewUserRepository(db.GetDB())

	return &App{
		userUC: userUceCase.NewUserUseCase(userRepo),
	}
}

func (a *App) Run(port string) error {
	router := mux.NewRouter()

	userHttp.RegisterHTTPEndpoints(router, a.userUC)

	router.Use(common.CORSMiddlware)
	router.Use(mux.CORSMethodMiddleware(router))
	a.httpServer = &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
