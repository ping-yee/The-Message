package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID        int            `json:"id" gorm:"primaryKey;auto_increment"`
	Token     string         `json:"token" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt,omitempty" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) repository.GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (p *GameRepository) GetGameById(ctx context.Context, id int) (*repository.Game, error) {
	game := new(repository.Game)

	result := p.db.First(game, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) CreateGame(ctx context.Context, game *repository.Game) (*repository.Game, error) {

	result := p.db.Save(game)

	if result.Error != nil {
		return nil, result.Error
	}

	return game, nil
}

func (p *GameRepository) DeleteGame(ctx context.Context, id int) error {
	game := new(repository.Game)

	result := p.db.Delete(game, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
