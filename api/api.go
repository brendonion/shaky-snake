package api

import (
	"encoding/json"
	"net/http"
)

const (
	HEAD_BENDR     HeadType = "bendr"
	HEAD_DEAD               = "dead"
	HEAD_FANG               = "fang"
	HEAD_PIXEL              = "pixel"
	HEAD_REGULAR            = "regular"
	HEAD_SAFE               = "safe"
	HEAD_SAND_WORM          = "sand-worm"
	HEAD_SHADES             = "shades"
	HEAD_SMILE              = "smile"
	HEAD_TONGUE             = "tongue"
)

const (
	TAIL_BLOCK_BUM    = "block-bum"
	TAIL_CURLED       = "curled"
	TAIL_FAT_RATTLE   = "fat-rattle"
	TAIL_FRECKLED     = "freckled"
	TAIL_PIXEL        = "pixel"
	TAIL_REGULAR      = "regular"
	TAIL_ROUND_BUM    = "round-bum"
	TAIL_SKINNY       = "skinny"
	TAIL_SMALL_RATTLE = "small-rattle"
)

func NewStartRequest(req *http.Request) (*StartRequest, error) {
	decoded := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}

func (list *PointList) UnmarshalJSON(data []byte) error {
	var obj struct {
		Data []Point `json:"data"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	*list = obj.Data
	return nil
}

func (list *SnakeList) UnmarshalJSON(data []byte) error {
	var obj struct {
		Data []Snake `json:"data"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	*list = obj.Data
	return nil
}

func (snake Snake) Head() Point { return snake.Body[0] }

func (snake Snake) Tail() Point { return snake.Body[len(snake.Body)-1] }

type HeadType string

type TailType string

type StartRequest struct {
	GameID int `json:"game_id"`
}

type StartResponse struct {
	Color          string   `json:"color,omitempty"`
	Name           string   `json:"name,omitempty"`
	HeadURL        string   `json:"head_url,omitempty"`
	Taunt          string   `json:"taunt,omitempty"`
	HeadType       HeadType `json:"head_type,omitempty"`
	TailType       TailType `json:"tail_type,omitempty"`
	SecondaryColor string   `json:"secondary_color,omitempty"`
}

type MoveRequest struct {
	Food   PointList `json:"food"`
	Height int       `json:"height"`
	ID     int       `json:"id"`
	Snakes SnakeList `json:"snakes"`
	Turn   int       `json:"turn"`
	Width  int       `json:"width"`
	You    Snake     `json:"you"`
}

type MoveResponse struct {
	Move string `json:"move"`
}

type Snake struct {
	Body   PointList `json:"body"`
	Health int       `json:"health"`
	ID     string    `json:"id"`
	Length int       `json:"length"`
	Name   string    `json:"name"`
	Taunt  string    `json:"taunt"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// PointList parses List<Point> into []Point
type PointList []Point

// SnakeList parses List<Snake> into []Snake
type SnakeList []Snake
