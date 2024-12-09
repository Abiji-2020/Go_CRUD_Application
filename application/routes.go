package application

import (
	"net/http"
	"os"

	"github.com/Abiji-2020/go-curd/handler"
	"github.com/Abiji-2020/go-curd/repository/order"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		content, err := os.ReadFile("README.md")
		if err != nil {
			http.Error(w, "Failed to read File", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/markdown")
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	})
	router.Route("/orders", a.loadOrderRoutes)
	a.router = router
}
func (a *App) loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: a.rdb,
		},
	}
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateBYID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
