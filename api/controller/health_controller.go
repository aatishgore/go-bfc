package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func (healthcontroller HealthController) GetHealth(ctx *gin.Context) {
	log.Info("Start GetHealth")
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
	log.Info("End GetHealth")
}
