package mysql

import (
	"context"
	"time"

	"github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id        int `gorm:"primaryKey;auto_increment"`
	Token     string
	Players   []Player
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
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

func (g *GameRepository) GetGameWithPlayers(ctx context.Context, id int) (*repository.Game, error) {
	var game repository.Game
	if err := g.db.Preload("Players").First(&game, id).Error; err != nil {
		return nil, err
	}
	return &game, nil
}
