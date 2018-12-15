package main

import (
	"fmt"
	"net/http"

	"github.com/brendonion/shaky-snake/api"
)

func Start(res http.ResponseWriter, req *http.Request) {
	respond(res, api.StartResponse{
		Color:          "orange",
		Name:           "shaky-snake",
		Taunt:          "Let's shake n' bake!",
		HeadType:       api.HEAD_TONGUE,
		TailType:       api.TAIL_SKINNY,
		SecondaryColor: "red",
		HeadURL:        "https://images-na.ssl-images-amazon.com/images/I/91fVmr46sLL._SL1500_.jpg",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	currentMove := "down"
	_, err := api.NewMoveRequest(req)
	if err != nil {
		fmt.Println("ERROR: ", err)
		respond(res, api.MoveResponse{
			Move: "up",
		})
		return
	}

	respond(res, api.MoveResponse{
		Move: currentMove,
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
