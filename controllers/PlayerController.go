package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service-pattern-go/interfaces"
	"service-pattern-go/models/dto"
	"service-pattern-go/viewmodels"

	"github.com/go-chi/chi"
)

type PlayerController struct {
	interfaces.IPlayerService
}

func (controller *PlayerController) AddUser(res http.ResponseWriter, req *http.Request) {
	var player dto.PlayerAddRequest
	// Phân tích dữ liệu JSON từ Body của yêu cầu
	if err := json.NewDecoder(req.Body).Decode(&player); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Request: %+v\n", player) // Using Printf for formatted output
}

func (controller *PlayerController) GetPlayerScore(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{scores})
}
