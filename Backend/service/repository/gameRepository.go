package repository

import "context"

type Game struct {
	Id    int
	Token string
}

type GameRepository interface {
	GetGameById(ctx context.Context, id int) (*Game, error)
	CreateGame(ctx context.Context, game *Game) (*Game, error)
	DeleteGame(ctx context.Context, id int) error
}
