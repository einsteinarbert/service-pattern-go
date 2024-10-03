package main

import (
	"service-pattern-go/controllers"
	"service-pattern-go/infrastructures"
	"service-pattern-go/repositories"
	"service-pattern-go/services"
	"sync"
)

type IServiceContainer interface {
	InjectPlayerController() controllers.PlayerController
}

type kernel struct{}

func (k *kernel) InjectPlayerController() controllers.PlayerController {
	handler := infrastructures.NewDbHandler("/var/tmp/tennis.db")

	playerRepository := &repositories.PlayerRepository{IDbHandler: handler}
	playerService := &services.PlayerService{IPlayerRepository: &repositories.PlayerRepositoryWithCircuitBreaker{playerRepository}}
	playerController := controllers.PlayerController{IPlayerService: playerService}

	return playerController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
