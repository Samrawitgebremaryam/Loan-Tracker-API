package loan_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LoanController) ViewLoanStatus(c *gin.Context) {
	loanID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
		return
	}

	userID := c.MustGet("user_id").(primitive.ObjectID)

	resp, err := lc.loanUsecase.ViewLoanStatus(c.Request.Context(), userID, loanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
