package service

import "github.com/gin-gonic/gin"

type DomainResponse struct {
	Domains []string `json:"domains"`
}

func EchoDomain(context *gin.Context) {
	domains := []string{"www.eastbuy.com", "eastbuy.com"}
	response := DomainResponse{Domains: domains}
	context.JSON(200, response)
}
