package http

import (
	"context"
	"fmt"
	"github.com/ent1k1377/wallet/internal/config"
	"github.com/ent1k1377/wallet/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer    *http.Server
	engine        *gin.Engine
	walletHandler *handler.Wallet
}

func NewServer(walletHandler *handler.Wallet, cfg *config.ServerConfig) *Server {
	fmt.Println("create server")
	engine := gin.Default()

	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: engine,
		},
		engine:        engine,
		walletHandler: walletHandler,
	}
}

func (s *Server) Run() error {
	s.SetRoutes()

	return s.httpServer.ListenAndServe()
}

func (s *Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
