package hall_of_fame

import (
	"context"
	"errors"
	"fmt"
	"time"

	c "github.com/helloproclub/cassiopeia"
	"github.com/jinzhu/gorm"
)

// HallOfFame hold functionallity to hall of fame
type HallOfFame struct {
	db *gorm.DB
}

// Option holds all necessary options for database.
type Option struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

// NewMySQL returns a pointer of MySQL instance and error.
func NewHallOfFame(opt Option) (*HallOfFame, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset))
	if err != nil {
		return &HallOfFame{}, errors.New(fmt.Sprintf("error when connecting to DB: %v", err))
	}

	db.DB().SetConnMaxLifetime(time.Minute)
	db.DB().SetMaxIdleConns(0)

	return &HallOfFame{db: db}, nil
}

func (h *HallOfFame) InsertAchievement(ctx context.Context, achievement c.Achievement) (c.Achievement, error) {
	return c.Achievement{}, nil
}

func (h *HallOfFame) ListAchievements(ctx context.Context, limit, offset int) ([]c.Achievement, error) {
	return []c.Achievement{}, nil
}
func (h *HallOfFame) InsertHero(ctx context.Context, hero c.Hero) (c.Hero, error) {
	return c.Hero{}, nil
}

func (h *HallOfFame) GetHeroByID(ctx context.Context, id uint) (c.Hero, error) {
	return c.Hero{}, nil
}

func (h *HallOfFame) GetHeroByUsername(ctx context.Context, username string) (c.Hero, error) {
	return c.Hero{}, nil
}
