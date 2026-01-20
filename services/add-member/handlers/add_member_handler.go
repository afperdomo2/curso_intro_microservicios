package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {
	// Dummy implementation not connected to DB
	c.JSON(http.StatusOK, gin.H{
		"message": "Member received successfully (Not saved to DB)",
	})
}
