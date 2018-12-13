package blogger

import (
	"net/http"
	"time"
)

type Blogger struct {
	APIKey   string
	BlogRoot string
	Client   *http.Client
	MaxPost  string
}

type BlogInfo struct {
	Kind        string    `json:"kind"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Published   time.Time `json:"published"`
	Updated     time.Time `json:"updated"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Option struct {
	BlogID  string
	APIKey  string
	MaxPost string
}
