package main

import (
	"fmt"
	"net/http"

	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func Start(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.StartResponse{
		Color: "orange",
		// Name:           "shaky-snake",
		// Taunt:          "Let's shake n' bake!",
		// HeadType:       api.HEAD_TONGUE,
		// TailType:       api.TAIL_SKINNY,
		// SecondaryColor: "red",
		// HeadURL:        "https://images-na.ssl-images-amazon.com/images/I/91fVmr46sLL._SL1500_.jpg",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	// Decode request
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		println("Bad move request: %v", err)
	}

	// Init current move
	currentMove := apiEntity.Up

	// Init board
	manager := game.InitializeBoard(&decoded)

	// Init your snake
	you := manager.Req.You

	// TODO: Find the closest food coords and a good time to go after that food
	// Find a path to nearest food if low on health
	if you.Health <= 50 {
		pathToFood, err := manager.FindPath(manager.OurHead, manager.Req.Board.Food[0])
		if err != nil {
			println("ERROR - No path to food!")
		}
		// Get direction from coords to food
		currentMove = lib.DirectionFromCoords(pathToFood[0], pathToFood[1])
		println("GO TO FOOD")
	} else if you.Health < 95 && len(you.Body) > 2 {
		pathToTail, err := manager.FindPath(manager.OurHead, you.Body[len(you.Body)-1])
		if err != nil {
			println("ERROR - No path to tail!")
		}
		// Get direction from coords to tail
		currentMove = lib.DirectionFromCoords(pathToTail[0], pathToTail[1])
		println("GO TO TAIL")
	} else if len(you.Body) == 2 && you.Body[0] != you.Body[1] {
		direction := lib.DirectionFromCoords(you.Body[0], you.Body[1])
		switch direction {
		case apiEntity.Up:
			currentMove = apiEntity.Left
		case apiEntity.Down:
			currentMove = apiEntity.Right
		case apiEntity.Left:
			currentMove = apiEntity.Down
		case apiEntity.Right:
			currentMove = apiEntity.Up
		}
		fmt.Println("TURN RIGHT")
	} else {
		currentMove = lib.DirectionFromCoords(you.Body[0], manager.GameBoard.GetValidTiles(you.Body[0])[0])
		fmt.Println("GO RANDOM %v", currentMove)
	}

	lib.Respond(res, api.MoveResponse{
		Move: currentMove,
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}

func Ping(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}
