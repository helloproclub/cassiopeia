package cassiopeia

import (
	"context"
	"time"
)

// move to cassiopeia.go

// Need comment documentation here
type HallOfFame interface {
	// Need comment documentation here
	InsertAchievement(context.Context, Achievement) (Achievement, error)
	// Need comment documentation here
	ListAchievements(context.Context, int, int) ([]Achievement, error)
	// Need comment documentation here
	InsertHero(context.Context, Hero) (Hero, error)
	// Need comment documentation here
	GetHeroByID(context.Context, uint) (Hero, error)
	// Need comment documentation here
	GetHeroByUsername(context.Context, string) (Hero, error)
}

// Need comment documentation here
type Achievement struct {
	// Need comment documentation here
	ID uint64 `json:"-"`
	// Need comment documentation here
	Published time.Time `json:"published"`
	// Need comment documentation here
	Updated time.Time `json:"updated"`
	// Need comment documentation here
	Created time.Time `json:"-"`
	// Need comment documentation here
	Heros []Hero `json:"heros,omitempty"`
	// Need comment documentation here
	Issuer string `json:"issuer"`
	// Need comment documentation here
	Title string `json:"title"`
	// Need comment documentation here
	Attainment string `json:"attainment"`
	// Need comment documentation here
	PostID string `json:"post_id"`
}

// Need comment documentation here
type Hero struct {
	// Need comment documentation here
	ID uint64 `json:"-"`
	// Need comment documentation here
	Updated time.Time `json:"-"`
	// Need comment documentation here
	Created time.Time `json:"-"`
	// Need comment documentation here
	Achievements []Achievement `json:"achievements,omitempty"`
	// Need comment documentation here
	Username string `json:"username"`
	// Need comment documentation here
	Fullname string `json:"fullname"`
	// Need comment documentation here
	Power string `json:"power"`
	// Need comment documentation here
	Bio string `json:"bio"`
	// Need comment documentation here
	ProclubYear uint16 `json:"proclub_year"`
	// Need comment documentation here
	FacebookUsername string `json:"facebook_username"`
	// Need comment documentation here
	TwitterUsername string `json:"twitter_username"`
	// Need comment documentation here
	InstagramUsername string `json:"instagram_username"`
	// Need comment documentation here
	TelegramUsername string `json:"telegram_username"`
	// Need comment documentation here
	LineID string `json:"line_id"`
}

type HeroAchievementRelation struct {
	// Need comment documentation here
	ID uint64
	// Need comment documentation here
	AchievementID uint64
	// Need comment documentation here
	HeroID uint64
}
