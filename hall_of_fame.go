package cassiopeia

import (
	"context"
	"time"
)

type HallOfFame interface {
	InsertAchievement(context.Context, Achievement) (Achievement, error)
	ListAchievements(context.Context, int, int) ([]Achievement, error)
	InsertHero(context.Context, Hero) (Hero, error)
	GetHeroByID(context.Context, uint) (Hero, error)
	GetHeroByUsername(context.Context, string) (Hero, error)
}

type Achievement struct {
	ID         uint64    `json:"-"`
	Published  time.Time `json:"published"`
	Updated    time.Time `json:"updated"`
	Created    time.Time `json:"-"`
	Heros      []Hero    `json:"heros,omitempty"`
	Issuer     string    `json:"issuer"`
	Title      string    `json:"title"`
	Attainment string    `json:"attainment"`
	PostID     string    `json:"post_id"`
}

type Hero struct {
	ID                uint64        `json:"-"`
	Updated           time.Time     `json:"-"`
	Created           time.Time     `json:"-"`
	Achievements      []Achievement `json:"achievements,omitempty"`
	Username          string        `json:"username"`
	Fullname          string        `json:"fullname"`
	Power             string        `json:"power"`
	Bio               string        `json:"bio"`
	ProclubYear       uint16        `json:"proclub_year"`
	FacebookUsername  string        `json:"facebook_username"`
	TwitterUsername   string        `json:"twitter_username"`
	InstagramUsername string        `json:"instagram_username"`
	TelegramUsername  string        `json:"telegram_username"`
	LineID            string        `json:"line_id"`
}

type HeroAchievementRelation struct {
	ID            uint64
	AchievementID uint64
	HeroID        uint64
}
