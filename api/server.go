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

// sever connection
func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}

// error response helper
func errorsResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
