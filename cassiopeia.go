// Need comment documentation here
package cassiopeia

import (
	"context"
	"fmt"
	"regexp"
	"github.com/helloproclub/cassiopeia/blogger"
	"github.com/helloproclub/cassiopeia/hall_of_fame"
)

// I guest this is the main package of this service right?
// so, every interface needed in this package should be declared here(e.g blogger and halloffame)
// we also need to create interface that will be implemented by Cassiopeia

// this package shuld be implemented by cassiopeia, the intereface name is my random thought so i dont realy like it.
// every function on Cassiopeia shoud be listed here first
type ProclubMemberHallofFame interface {
}

// Need comment documentation here
type Cassiopeia struct {
	blogger blogger.Blogger
	hof     hall_of_fame.HallOfFame
}

// Need comment documentation here
func NewCassiopeia(b blogger.Blogger, h hall_of_fame.HallOfFame) (*Cassiopeia, error) {
	return &Cassiopeia{
		blogger: b,
		hof:     h,
	}, nil
}

// Need comment documentation here
func (c *Cassiopeia) ListPosts(ctx context.Context, label, pageToken string) (blogger.PostList, error) {
	var (
		err      error
		postList blogger.PostList
	)

	if label != "" {
		// fetch posts with label
		if postList, err = c.blogger.ListPostsByLabel(ctx, label, pageToken); err != nil {
			return blogger.PostList{}, err
		}
	} else {
		// fetch posts without label
		if postList, err = c.blogger.ListPosts(ctx, pageToken); err != nil {
			return blogger.PostList{}, err
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

// Need comment documentation here
func (c *Cassiopeia) GetPostByPath(ctx context.Context, path string) (blogger.Post, error) {
	var (
		err  error
		post blogger.Post
	)

	if post, err = c.blogger.GetPostByPath(ctx, path); err != nil {
		return blogger.Post{}, err
	}

	return post, nil
}
