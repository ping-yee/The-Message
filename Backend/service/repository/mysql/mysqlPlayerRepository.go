package mysql

import (
	"context"
	domain "github.com/Game-as-a-Service/The-Message/service/repository"
	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	"gorm.io/gorm"
	"time"
)

type Player struct {
	gorm.Model
<<<<<<< HEAD
	Id           int `gorm:"primaryKey;auto_increment"`
	Name         string
	GameId       int `gorm:"foreignKey:GameId;references:Id"`
	IdentityCard string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoCreateTime"`
	DeletedAt    gorm.DeletedAt
=======
	Id        int            `gorm:"primaryKey;auto_increment"`
	Name      string         `gorm:"type:varchar(255);NOT NULL"`
	Group     string         `gorm:"type:varchar(15);NOT NULL"`
	G_id      string         `gorm:"type:varchar(15);NOT NULL"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
>>>>>>> f5444a2 (fix: fix the soft delete and cannot auto increame time problem.)
}

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) repository.PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (p *PlayerRepository) CreatePlayer(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	err := p.db.Table("players").Create(player).Error
	return player, err
}
