package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mybank/db/sqlc"
	"github.com/mybank/utils"
)

type CreateUserRequest struct {
	Username     string `json:"username" binding:"required" alphanumeric:"true"`
	Password     string `json:"hash_password" binding:"required" min:"6"`
	Email        string `json:"email" binding:"required" email:"true"`
	FullName     string `json:"full_name" binding:"required" `
}

type CreateUserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}



func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorsResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorsResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		HashPassword: hashedPassword,
		Email:        req.Email,
		FullName:     req.FullName,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusConflict, errorsResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorsResponse(err))
		return
	}

	res := CreateUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, res)
}