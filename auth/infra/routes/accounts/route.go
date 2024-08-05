package accounts

import (
	"auth/infra/core/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountCreationModel struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateAccount(context *gin.Context) {
	var json AccountCreationModel

	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := entities.CreateAccount(
		json.Name, 
		json.Email, 
		json.Password,
	)

	context.JSON(http.StatusOK, gin.H{"account": account})
}	