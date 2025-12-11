package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mybank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// router setup
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)


	server.router = router

	return server

}

// server connection
func (server *Server) Start(ServerAddress string) error {
	return server.router.Run(ServerAddress)
}

// error response helper
func errorsResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
