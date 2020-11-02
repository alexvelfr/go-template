package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexvelfr/go-template/app"
	apphttp "github.com/alexvelfr/go-template/app/delivery/http"
	apprepo "github.com/alexvelfr/go-template/app/repo/mock"
	appusecase "github.com/alexvelfr/go-template/app/usecase"
	"github.com/gin-gonic/gin"
)

// App ...
type App struct {
	appUC      app.Usecase
	appRepo    app.Repository
	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	repo := apprepo.NewRepo()
	uc := appusecase.NewUsecase(repo)
	return &App{
		appUC:   uc,
		appRepo: repo,
	}
}

// Run run application
func (a *App) Run(port string) error {
	defer a.appRepo.Close()
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	apphttp.RegisterHTTPEndpoints(router, a.appUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
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
