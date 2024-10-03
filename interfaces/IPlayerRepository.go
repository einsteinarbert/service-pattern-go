package interfaces

import (
	"service-pattern-go/models"
)

type IPlayerRepository interface {
	GetPlayerByName(name string) (models.PlayerModel, error)
}
