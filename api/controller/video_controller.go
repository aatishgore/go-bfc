package controller

import (
	"net/http"
	"stefano/api/models"
	"stefano/api/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Videos struct {
	Service service.VideoService
}

func (videos Videos) GetVideos(ctxt *gin.Context) {
	log.Info("Start GetVideos")
	var request models.VideosRequest

	if err := ctxt.ShouldBindQuery(&request); err != nil {
		ctxt.AbortWithStatusJSON(400, err)

		return
	}

	response := videos.Service.GetVideos(request.Type)
	ctxt.JSON(http.StatusOK, response)

	log.Info("End GetVideos")
	return
}
