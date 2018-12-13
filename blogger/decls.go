package blogger

import "fmt"

const (
	epBlogGet        = ""
	epPostsList      = "/posts"
	epPostsSearch    = "/posts/search"
	epPostsGetByPath = "/posts/bypath"
)

var (
	epPostsGet = func(postID string) string {
		return fmt.Sprintf("/posts/%v", postID)
	}
)
