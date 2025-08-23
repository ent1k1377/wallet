package http

import "github.com/gin-gonic/gin"

func (s *Server) SetRoutes() {
	api := s.engine.Group("/api")
	{
		api.GET("/check", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "all good"})
		})
		api.POST("/send", s.walletHandler.Send)
		api.GET("/transactions", s.walletHandler.GetLast)
		api.GET("/wallet/:address/balance", s.walletHandler.GetBalance)
	}
}
