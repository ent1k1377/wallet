package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ent1k1377/wallet/internal/config"
	"github.com/ent1k1377/wallet/internal/database/postgres"
	"github.com/ent1k1377/wallet/internal/database/postgres/repository"
	"github.com/ent1k1377/wallet/internal/service"
	myhttp "github.com/ent1k1377/wallet/internal/transport/http"
	"github.com/ent1k1377/wallet/internal/transport/http/handler"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var (
	shutdownTimeout = time.Second * 5
)

type App struct {
	closer Closer
	server *myhttp.Server
	db     *postgres.DB
}

func New() *App {
	fmt.Println("create app")
	cfg := config.MustLoadConfig()
	
	db := postgres.NewDB(cfg.DatabaseConfig)
	walletRepository := repository.NewWallet(db.GetPool())
	
	walletService := service.NewWallet(walletRepository)
	walletHandler := handler.NewWallet(walletService)
	
	server := myhttp.NewServer(walletHandler, cfg.ServerConfig)
	return &App{
		server: server,
		db:     db,
	}
}

func (a *App) Run() {
	fmt.Println("starting app")
	a.closer.Add(a.server.Close)
	a.closer.Add(a.db.Close)
	
	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	ctx, cancel := context.WithCancel(sigCtx)
	defer cancel()
	
	go func() {
		fmt.Println("starting server")
		if err := a.server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			cancel()
		}
	}()
	
	<-ctx.Done()
	
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	
	if err := a.closer.Close(shutdownCtx); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Server shutdown gracefully")
	}
}
