package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userUsecase.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
