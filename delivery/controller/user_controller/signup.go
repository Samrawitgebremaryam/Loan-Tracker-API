package user_controller

import (
	"bytes"
	"fmt"
	"io"
	"loan_tracker_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) SignUp(c *gin.Context) {
	var req domain.SignupRequest

	body, _ := io.ReadAll(c.Request.Body)
	fmt.Println("Request Body: ", string(body))

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error in binding request signup is : ", err)
		return
	}

	resp, err := uc.userUsecase.SignUp(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
