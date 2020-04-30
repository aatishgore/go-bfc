package service

import (
	"fmt"
	"stefano/api/models"

	"github.com/sirupsen/logrus"
)

type TransferService interface {
	GetStatus() models.Cards
}

type transferService struct {
}

func NewTransferService() TransferService {
	return &transferService{}
}

var temp = 0

//NOTE Response is hardcoded for now as requested by Adnan
func (transferService *transferService) GetStatus() models.Cards {
	logrus.Info("GetStatus Service Start")
	Cards := models.Cards{}

	if temp < 100 {
		temp = temp + 20
	} else {
		temp = 0
	}
	percent := fmt.Sprintf("%d%%", temp)
	Cards.Cards = []models.Card{
		{CardName: "card1", Percentage: percent},
		{CardName: "card2", Percentage: percent},
		{CardName: "card3", Percentage: percent},
	}
	return Cards
}
