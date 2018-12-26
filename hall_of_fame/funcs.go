// Need comment documentation here
package hall_of_fame

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// HallOfFamer hold functionallity to hall of fame
type HallOfFamer struct {
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
func NewHallOfFamer(opt Option) (HallOfFame, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset))
	if err != nil {
		return &HallOfFamer{}, errors.New(fmt.Sprintf("error when connecting to DB: %v", err))
	}

	db.DB().SetConnMaxLifetime(time.Minute)
	db.DB().SetMaxIdleConns(0)

	return &HallOfFamer{db: db}, nil
}

// Need comment documentation here
func (h *HallOfFamer) InsertAchievement(ctx context.Context, achievement Achievement) (Achievement, error) {
	return Achievement{}, nil
}

// Need comment documentation here
func (h *HallOfFamer) ListAchievements(ctx context.Context, limit, offset int) ([]Achievement, error) {
	return []Achievement{}, nil
}

// Need comment documentation here
func (h *HallOfFamer) InsertHero(ctx context.Context, hero Hero) (Hero, error) {
	return Hero{}, nil
}

// Need comment documentation here
func (h *HallOfFamer) GetHeroByID(ctx context.Context, id uint) (Hero, error) {
	return Hero{}, nil
}

// Need comment documentation here
func (h *HallOfFamer) GetHeroByUsername(ctx context.Context, username string) (Hero, error) {
	return Hero{}, nil
}
