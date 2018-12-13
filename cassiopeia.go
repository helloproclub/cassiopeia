package cassiopeia

import (
	"context"
	"fmt"
	"regexp"
)

type Cassiopeia struct {
	blogger Blogger
	hof     HallOfFame
}

func NewCassiopeia(b Blogger, h HallOfFame) (*Cassiopeia, error) {
	return &Cassiopeia{
		blogger: b,
		hof:     h,
	}, nil
}

func (c *Cassiopeia) ListPosts(ctx context.Context, label, pageToken string) (PostList, error) {
	var (
		err      error
		postList PostList
	)

	if label != "" {
		// fetch posts with label
		if postList, err = c.blogger.ListPostsByLabel(ctx, label, pageToken); err != nil {
			return PostList{}, err
		}
	} else {
		// fetch posts without label
		if postList, err = c.blogger.ListPosts(ctx, pageToken); err != nil {
			return PostList{}, err
		}
	}

	rx := regexp.MustCompile(`\/[0-9]{4}\/[0-9]{2}\/[a-z-]*.html`)
	for i, _ := range postList.Items {
		postList.Items[i].Path = fmt.Sprintf("/posts%v", rx.FindString(postList.Items[i].URL))
		postList.Items[i].URL = ""
		postList.Items[i].Content = ""
	}

	return postList, nil
}

func (c *Cassiopeia) GetPostByPath(ctx context.Context, path string) (Post, error) {
	var (
		err  error
		post Post
	)

	if post, err = c.blogger.GetPostByPath(ctx, path); err != nil {
		return Post{}, err
	}

	return post, nil
}
