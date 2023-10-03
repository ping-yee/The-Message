package repository

import "context"

type Player struct {
	Id    int
	Name  string
	Group string
	G_Id  int
}

type PlayerRepository interface {
	GetPlayerById(ctx context.Context, id int) (*Player, error)
	PlayerDrawCards(ctx context.Context, id int, count int) ([]string, error)
	UpdatePlayerCards(ctx context.Context, id int, cards []string) error
	PlayerPlayCard(ctx context.Context, id int, card string) error
	DeleteAllGamePlayers(ctx context.Context, id int) error
}
