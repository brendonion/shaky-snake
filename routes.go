package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func Start(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.StartResponse{
		Color: "blue",
		// Name:           "shaky-snake",
		// Taunt:          "Let's shake n' bake!",
		// HeadType:       apiEntity.HeadTongue,
		// TailType:       apiEntity.TailSkinny,
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

	// Set current move
	currentMove := apiEntity.Up

	// Set board
	manager := game.InitializeBoard(&decoded)

	// Set your snake
	you := manager.Req.You

	// Get average snake length
	totalLength := 0
	for _, snake := range manager.Req.Board.Snakes {
		totalLength += len(snake.Body)
	}
	averageLength := totalLength / len(manager.Req.Board.Snakes)

	if len(you.Body) <= averageLength || you.Health <= manager.Req.Board.Width {
		println("GO TO FOOD")
		// Find closest food
		closestFood := manager.Req.Board.Food[0]
		for i, food := range manager.Req.Board.Food {
			prevDistance := lib.Distance(manager.OurHead, closestFood)
			currentDistance := lib.Distance(manager.OurHead, food)
			if prevDistance > currentDistance {
				closestFood = manager.Req.Board.Food[i]
			}
		}
		pathToFood, err := manager.FindPath(manager.OurHead, closestFood)
		if err != nil {
			println("ERROR - No path to food!")
		}
		currentMove = lib.DirectionFromCoords(pathToFood[0], pathToFood[1])

	} else if you.Health < 95 && len(you.Body) > 2 {
		println("GO TO TAIL")
		pathToTail, err := manager.FindPath(manager.OurHead, you.Body[len(you.Body)-1])
		if err != nil {
			println("ERROR - No path to tail!")
		}
		currentMove = lib.DirectionFromCoords(pathToTail[0], pathToTail[1])

	} else if len(you.Body) == 2 && manager.OurHead != you.Body[1] {
		fmt.Println("GO RIGHT")
		direction := lib.DirectionFromCoords(manager.OurHead, you.Body[1])
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

	} else {
		fmt.Println("GO RANDOM")
		validTiles := manager.GameBoard.GetValidTiles(manager.OurHead)
		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		move := 0
		if len(validTiles) > 1 {
			move = random.Intn(len(validTiles) - 1)
		}
		currentMove = lib.DirectionFromCoords(manager.OurHead, validTiles[move])
	}

	fmt.Println("CURRENT MOVE: %v", currentMove)

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
