package controller

import (
	"net/http"
	"stefano/api/service"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	Service service.TransferService
}

func (transferController TransferController) GetTransferStatus(ctxt *gin.Context) {
	log.Info("Start TransferStatus")
	response := transferController.Service.GetStatus()
	ctxt.JSON(http.StatusOK, response)
	log.Info("End TransferStatus")
}
