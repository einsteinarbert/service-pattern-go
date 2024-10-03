package services

import (
	"service-pattern-go/interfaces"
	"service-pattern-go/models"
)

type PlayerService struct {
	interfaces.IPlayerRepository
}

func (service *PlayerService) GetScores(player1Name string) ([]models.PlayerModel, error) {
	player1, err := service.GetPlayerByName(player1Name)
	if err != nil {
		//Handle error
	}

	return player1, nil;
}
