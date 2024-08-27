package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := uc.userUsecase.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
