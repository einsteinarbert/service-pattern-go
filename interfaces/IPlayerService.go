package interfaces

import "service-pattern-go/models"

type IPlayerService interface {
	GetScores(player1Name string) ([]models.PlayerModel, error)
}
