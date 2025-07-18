package handler

import (
	"net/http"

	"github.com/buzyka/about_go/06_gin_project/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func CreateFullName(ctx *gin.Context) {

	var user entity.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return full name
	ctx.JSON(http.StatusOK, gin.H{"full_name": user.FullName()})
}
