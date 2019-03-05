package blogger

import (
	"context"
	"time"
)

// move the interface to cassiopeia.go

// Need comment documentation here
type Blog interface {
	// Need comment documentation here
	ListPosts(context.Context, string) (PostList, error)
	// Need comment documentation here
	ListPostsByLabel(context.Context, string, string) (PostList, error)
	// Need comment documentation here
	GetPostByPath(context.Context, string) (Post, error)
}

// Need comment documentation here
type PostList struct {
	// Need comment documentation here
	Kind string `json:"kind"`
	// Need comment documentation here
	NextPageToken string `json:"nextPageToken"`
	// Need comment documentation here
	Items []Post `json:"items"`
}

// Need comment documentation here
type Post struct {
	// Need comment documentation here
	ID string `json:"id"`
	// Need comment documentation here
	Published time.Time `json:"published"`
	// Need comment documentation here
	Updated time.Time `json:"updated"`
	// Need comment documentation here
	URL string `json:"url,omitempty"`
	// Need comment documentation here
	Path string `json:"path,omitempty"`
	// Need comment documentation here
	Title string `json:"title"`
	// Need comment documentation here
	Content string `json:"content,omitempty"`
	// Need comment documentation here
	Summary string `json:"summary,omitempty"`
	// Need comment documentation here
	Images []Image `json:"images"`
	// Need comment documentation here
	Author Author `json:"author"`
	// Need comment documentation here
	Labels []string `json:"labels"`
}

// Need comment documentation here
type Author struct {
	// Need comment documentation here
	DisplayName string `json:"displayName"`
	// Need comment documentation here
	URL string `json:"url"`
}

// Need comment documentation here
type Image struct {
	// Need comment documentation here
	URL string `json:"url"`
}
