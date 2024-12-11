package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	Config config
}

func New(Config config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: Config.RedisAddress,
		}),
		Config: Config,
	}
	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.ServerPort),
		Handler: a.router,
	}
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}
	fmt.Printf("Server is running on port: %d", a.Config.ServerPort)
	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()

		if err != nil {
			ch <- fmt.Errorf("failed to listen to server: %w", err)
		}
		close(ch)
	}()
	if err := a.callHealthCheck(); err != nil {
		return fmt.Errorf("health check failed %w", err)
	}

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil
}

func (a *App) callHealthCheck() error {
    client := http.Client{Timeout: 2 * time.Second}
    url := fmt.Sprintf("http://localhost:%d/healthCheck", a.Config.ServerPort)

    for i := 0; i < 3; i++ {
        resp, err := client.Get(url)
        if err == nil && resp.StatusCode == http.StatusOK {
            fmt.Println("Health Check Passed")
            return nil
        }
        if resp != nil {
            resp.Body.Close()
        }
        time.Sleep(1 * time.Second) // Wait before retrying
    }
    return fmt.Errorf("health check failed after retries")
}
