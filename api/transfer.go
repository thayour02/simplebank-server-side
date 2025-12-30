package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/mybank/db/sqlc"
)

type CreateTransferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) CreateTransfers(ctx *gin.Context) {
	var req CreateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorsResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}
	if !server.validAmount(ctx, req.FromAccountID, req.Amount) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(500, errorsResponse(err))
		return
	}
	ctx.JSON(200, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(404, errorsResponse(err))
			return false
		}
		ctx.JSON(500, errorsResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %d currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(400, errorsResponse(err))
		return false
	}

	return true
}

func (server *Server) validAmount(ctx *gin.Context, accountID int64, amount int64) bool {

	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(404, errorsResponse(err))
			return false
		}

		ctx.JSON(500, errorsResponse(err))
		return false
	}

	if account.Balance < amount {
		err := fmt.Errorf("account %d insufficient funds: %d < %d",
			account.ID, account.Balance, amount)
		ctx.JSON(400, errorsResponse(err))
		return false
	}
	return true
}
