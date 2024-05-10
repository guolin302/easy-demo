package service

import "github.com/gin-gonic/gin"

func Echo2Domain(context *gin.Context) {
	domains := []string{"test.eastbuy.com", "test-2eastbuy.com"}
	response := DomainResponse{Domains: domains}
	context.JSON(200, response)
}
