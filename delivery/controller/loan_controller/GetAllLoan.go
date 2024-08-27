package loan_controller

import (
	"loan_tracker_api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LoanController) ViewAllLoans(c *gin.Context) {
	status := c.Query("status")
	loans, err := lc.loanUsecase.ViewAllLoans(c.Request.Context(), domain.LoanStatus(status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}
