package loan_controller

import (
	"loan_tracker_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LoanController) ApplyForLoan(c *gin.Context) {
	var req domain.LoanApplication
	userID := c.MustGet("user_id").(primitive.ObjectID) // Assuming user_id is stored as ObjectID in the context

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := lc.loanUsecase.ApplyForLoan(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
