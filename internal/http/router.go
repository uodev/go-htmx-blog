package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	ID             int
	Title          string
	Excerpt        string
	ExcerptShort   string // Go tarafında kısaltılmış versiyon
	PublishedAt    time.Time
	PublishedAtStr string // Go tarafında formatlanmış tarih
}

type Router struct {
	render func(http.ResponseWriter, string, any)
}

func New(render func(http.ResponseWriter, string, any)) http.Handler {
	r := &Router{render: render}
	mux := chi.NewRouter()

	mux.Get("/home", r.Home)
	mux.Get("/posts", r.PostList)

	return mux
}

func (r *Router) Home(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"Title": "Anasayfa",
		"User":  "Uygar",
		"Year":  time.Now().Year(),
	}
	r.render(w, "home", data)
}

func (r *Router) PostList(w http.ResponseWriter, _ *http.Request) {
	posts := []Post{
		{
			ID:             1,
			Title:          "Go Templates ile Başlarken",
			Excerpt:        "Template mimarisi, base ve partial pratikleri…",
			ExcerptShort:   "Template mimarisi, base ve partial pratikleri…", // ister burada kısalt
			PublishedAt:    time.Now().AddDate(0, 0, -2),
			PublishedAtStr: time.Now().AddDate(0, 0, -2).Format("02 Jan 2006"),
		},
		{
			ID:             2,
			Title:          "HTMX’a Giriş",
			Excerpt:        "Partial render ve progressive enhancement…",
			ExcerptShort:   "Partial render ve progressive enhancement…",
			PublishedAt:    time.Now().AddDate(0, 0, -1),
			PublishedAtStr: time.Now().AddDate(0, 0, -1).Format("02 Jan 2006"),
		},
	}

	data := map[string]any{
		"Title": "Yazılar",
		"Year":  time.Now().Year(),
		"Posts": posts,
	}
	r.render(w, "posts", data)
}
