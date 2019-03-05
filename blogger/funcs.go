package blogger

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

)

func NewBlogger(opt Option) (*Blogger, error) {
	var (
		blogInfo   BlogInfo
		blogger    *Blogger
		err        error
		maxPostInt int
		params     map[string]string
	)

	if maxPostInt, err = strconv.Atoi(opt.MaxPost); err != nil {
		log.Printf("Error on parsing maxPost, then set maxPostInt to 9: %v", err)
		maxPostInt = 9
	}

	blogger = &Blogger{
		Client:   &http.Client{Timeout: time.Second},
		BlogRoot: fmt.Sprintf("https://www.googleapis.com/blogger/v3/blogs/%v", opt.BlogID),
		APIKey:   opt.APIKey,
		MaxPost:  strconv.Itoa(maxPostInt),
	}

	params = map[string]string{
		"fields": "kind, id, name, description, published, updated",
	}

	if blogger.getResource(epBlogGet, params, &blogInfo) != nil {
		return &Blogger{}, err
	}

	if blogInfo.Kind != "blogger#blog" {
		return &Blogger{}, errors.New("blog is not found")
	}

	return blogger, nil
}

func (b *Blogger) ListPosts(ctx context.Context, pageToken string) (PostList, error) {
	params := map[string]string{
		"fields":     "kind,nextPageToken,items(id, published, updated, url, title, content, author/displayName, author/url)",
		"orderBy":    "published",
		"maxResults": b.MaxPost,
	}

	if len(pageToken) != 0 {
		params["pageToken"] = pageToken
	}

	var postList PostList
	if err := b.getResource(epPostsList, params, &postList); err != nil {
		return PostList{}, err
	}

	return postList, nil
}

func (b *Blogger) ListPostsByLabel(ctx context.Context, label, pageToken string) (PostList, error) {
	params := map[string]string{
		"q":          fmt.Sprintf("label=\"%v\"", label),
		"fields":     "kind,nextPageToken,items(id, published, updated, url, title, content, author/displayName, author/url)",
		"orderBy":    "published",
		"maxResults": b.MaxPost,
	}

	if len(pageToken) != 0 {
		params["pageToken"] = pageToken
	}

	var postList PostList
	if err := b.getResource(epPostsSearch, params, &postList); err != nil {
		return PostList{}, err
	}

	return postList, nil
}

func (b *Blogger) GetPostByPath(ctx context.Context, postPath string) (Post, error) {
	params := map[string]string{
		"path":       postPath,
		"maxComents": "0",
		"fields":     "id, published, updated, url, title, content, author/displayName, author/url",
	}

	var post Post
	if err := b.getResource(epPostsGetByPath, params, &post); err != nil {
		return Post{}, err
	}

	if post.ID == "" {
		return Post{}, errors.New("Not found")
	}

	post.URL = ""

	return post, nil
}

func (b *Blogger) getResource(endpoint string, params map[string]string, result interface{}) error {
	// prepare url
	var urlBuilder strings.Builder
	urlBuilder.WriteString(b.BlogRoot)
	if len(endpoint) > 0 {
		urlBuilder.WriteString(fmt.Sprintf("%v", endpoint))
	}

	// preparing request
	request, err := http.NewRequest("GET", urlBuilder.String(), nil)

	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	q.Add("key", b.APIKey)
	request.URL.RawQuery = q.Encode()

	request.Header.Add("Accept-Encoding", "gzip")

	response, err := b.Client.Do(request)
	if err != nil {
		return errors.New(fmt.Sprintf("Error doing request to blogger: %v", err))
	}

	defer response.Body.Close()

	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		defer reader.Close()
	default:
		reader = response.Body
	}

	err = json.NewDecoder(reader).Decode(&result)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decoding response from blogger: %v", err))

	}

	return nil
}
