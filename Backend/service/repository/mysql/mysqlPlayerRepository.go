package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Id        int            `json:"id" gorm:"primaryKey;auto_increment"`
	Name      string         `json:"name" gorm:"type:varchar(255);NOT NULL"`
	Group     string         `json:"group" gorm:"type:varchar(15);NOT NULL"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
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
