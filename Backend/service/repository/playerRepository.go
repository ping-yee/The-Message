package repository

<<<<<<< HEAD
<<<<<<< HEAD
import (
	"context"
)

type Player struct {
	Id           int `gorm:"primaryKey;auto_increment"`
	Name         string
	GameId       int `gorm:"foreignKey:GameId;references:Id"`
	IdentityCard string
}

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *Player) (*Player, error)
=======
=======
>>>>>>> 5305246 (feat: initial the player componenet.)
import "context"

type Player struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
	G_Id  int    `json:"g_id"`
}

type PlayerRepository interface {
	GetPlayerById(ctx context.Context, id int) (*Player, error)
	PlayerDrawCards(ctx context.Context, id int, count int) ([]string, error)
	UpdatePlayerCards(ctx context.Context, id int, cards []string) error
	PlayerPlayCard(ctx context.Context, id int, card string) error
	DeleteAllGamePlayers(ctx context.Context, id int) error
<<<<<<< HEAD
>>>>>>> e21ca85 (feat: initial the player componenet.)
=======
>>>>>>> 5305246 (feat: initial the player componenet.)
}
