package repositories

import (
	"service-pattern-go/interfaces"
	"service-pattern-go/models"

	"github.com/afex/hystrix-go/hystrix"

	_ "fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PlayerRepositoryWithCircuitBreaker struct {
	PlayerRepository interfaces.IPlayerRepository
}

func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) ([]models.PlayerModel, error) {

	output := make(chan []models.PlayerModel, 1) // make channel
	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_player_by_name", func() error {

		player, _ := repository.PlayerRepository.GetPlayerByName(name)

		output <- player
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return []models.PlayerModel{}, err
	}
}

type PlayerRepository struct {
	interfaces.IDbHandler
}

func (repository *PlayerRepository) GetPlayerByName(name string) ([]models.PlayerModel, error) {

	rows, err := repository.Query(name)
	if err != nil {
		return []models.PlayerModel{}, err
	}
	return rows, nil
}
