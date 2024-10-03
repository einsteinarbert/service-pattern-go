package interfaces

import "service-pattern-go/models"

type IDbHandler interface {
	Execute(statement string)
	Query(playerName string) ([]models.PlayerModel, error)
}
