package cassiopeia

import (
	"context"
	"time"
)

type Blogger interface {
	ListPosts(context.Context, string) (PostList, error)
	ListPostsByLabel(context.Context, string, string) (PostList, error)
	GetPostByPath(context.Context, string) (Post, error)
}

type PostList struct {
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
	Items         []Post `json:"items"`
}

type Post struct {
	ID        string    `json:"id"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
	URL       string    `json:"url,omitempty"`
	Path      string    `json:"path,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content,omitempty"`
	Summary   string    `json:"summary,omitempty"`
	Images    []Image   `json:"images"`
	Author    Author    `json:"author"`
	Labels    []string  `json:"labels"`
}

type Author struct {
	DisplayName string `json:"displayName"`
	URL         string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}
