package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"webProject/restful"
	"webProject/service"
)

func ShowPlat(c *gin.Context) {
	fmt.Println("[common][showPlat] showPlat data")
	showValues, err := service.CommonShowValues()
	restful.WithDataset(showValues, err, c)
}