// Need comment documentation here
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	c "github.com/helloproclub/cassiopeia"
)

// Need comment documentation here
type Handler struct {
	cas *c.Cassiopeia
}

// Need comment documentation here
func NewHandler(c *c.Cassiopeia) *Handler {
	return &Handler{
		cas: c,
	}
}

// Need comment documentation here
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	type HelloWorld struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}

	var (
		request HelloWorld
	)

	if readRequestBody(w, r, &request) != nil {
		return
	}

	response := HelloWorld{Name: "Cassiopeia", Message: "Hello " + request.Name}

	writeResponse(w, r, response)
}

// Need comment documentation here
func (h *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		label     string
		pageToken string
		postList  c.PostList
	)

	pageToken = readRequestQuery(r, "page_token")
	label = readRequestQuery(r, "label")

	if postList, err = h.cas.ListPosts(r.Context(), label, pageToken); err != nil {
		return
	}

	writeResponse(w, r, postList)
}

// Need comment documentation here
func (h *Handler) GetPostByPath(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		post c.Post
	)

	path := strings.TrimPrefix(r.URL.Path, "/posts")

	if post, err = h.cas.GetPostByPath(r.Context(), path); err != nil {
		return
	}

	writeResponse(w, r, post)
}

func readRequestBody(w http.ResponseWriter, r *http.Request, res interface{}) error {
	var (
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, "The request couldn't be encoded", http.StatusBadRequest)
		return err
	}
	return nil
}

func readRequestQuery(r *http.Request, q string) string {
	return r.URL.Query().Get(q)
}

func writeResponse(w http.ResponseWriter, r *http.Request, res interface{}) {
	var (
		err error
	)

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "The response couldn't be proceed", http.StatusInternalServerError)
	}
}
