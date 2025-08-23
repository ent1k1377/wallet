package http

func (s *Server) SetRoutes() {
	api := s.engine.Group("/api")
	{
		api.POST("/send", s.walletHandler.Send)
		api.GET("/transactions", s.walletHandler.GetLast)
		api.GET("/wallet/:address/balance", s.walletHandler.GetBalance)
	}
}
