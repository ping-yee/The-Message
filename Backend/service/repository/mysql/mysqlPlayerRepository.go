package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Name      string
	Group     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) repository.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (p *PlayerRepository) GetPlayerById(ctx context.Context, id int) (*repository.Player, error) {

	return &repository.Player{}, nil
}

func (p *PlayerRepository) PlayerDrawCards(ctx context.Context, id int, count int) ([]string, error) {

	return []string{"mock1"}, nil
}

func (p *PlayerRepository) UpdatePlayerCards(ctx context.Context, id int, cards []string) error {

	return nil
}

func (p *PlayerRepository) PlayerPlayCard(ctx context.Context, id int, card string) error {

	return nil
}

func (p *PlayerRepository) DeleteAllGamePlayers(ctx context.Context, id int) error {

	return nil

}
