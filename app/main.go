package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	c "github.com/helloproclub/cassiopeia"
	b "github.com/helloproclub/cassiopeia/blogger"
	h "github.com/helloproclub/cassiopeia/hall_of_fame"
	"github.com/helloproclub/cassiopeia/handler"
	"github.com/subosito/gotenv"
)

func main() {
	// declaring variables in alphabetical order
	var (
		blo *b.Blogger
		cas *c.Cassiopeia
		err error
		hof *h.HallOfFame
	)

	// load environment variables
	gotenv.Load(
		os.Getenv("GOPATH") + "/src/github.com/helloproclub/cassiopeia/.env",
	)

	// create new blogger
	if blo, err = b.NewBlogger(
		b.Option{
			BlogID:  os.Getenv("BLOGGER_BLOG_ID"),
			APIKey:  os.Getenv("BLOGGER_API_KEY"),
			MaxPost: os.Getenv("BLOGGER_MAX_POST"),
		},
	); err != nil {
		fmt.Println(err)
	}

	// create new cassiopeia
	if cas, err = c.NewCassiopeia(blo, hof); err != nil {
		fmt.Println(err)
	}

	// create new handler
	hc := handler.NewHandler(cas)

	// create new router
	router := mux.NewRouter()
	router.Use(loggerMiddleware)

	// root
	router.HandleFunc("/", hc.Ping)

	// posts routers
	router.HandleFunc("/posts", hc.ListPosts).Methods("GET")
	router.HandleFunc("/posts/{year}/{month}/{title}", hc.GetPostByPath).Methods("GET")

	// listen and serve
	appHostname := os.Getenv("APP_HOSTNAME")
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8000"
	}

	go func() {
		log.Println("🏃 🏃 🏃 cassiopeia will be ready in no time 🏃 🏃 🏃")
	}()

	http.ListenAndServe(
		fmt.Sprintf("%v:%v", appHostname, appPort),
		router,
	)
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
