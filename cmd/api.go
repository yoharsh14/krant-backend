package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"yoharsh14/krant-backend/internal/business/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Dependencies struct {
	userHandler *user.Handler
}

func (app *application) mount(ctx context.Context) http.Handler{
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID) // for rate limitting
	r.Use(middleware.RealIP)    // for rate limtting and analytics
	r.Use(middleware.Logger)    // logging
	r.Use(middleware.Recoverer) // recover from crashes

	// set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all goood"))
	})
	userReposotory := user.NewRepository(app.db.Database("krant"))
	userService :=user.NewService(userReposotory)
	userHandler :=user.NewHandler(userService)
	r.Post("/user/create",userHandler.CreateUser)

	return r
}

func (app *application) run(h http.Handler) error{
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 40,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Minute * 2,
	}
	log.Printf("Server started at port %s", app.config.addr)
	return srv.ListenAndServe()
}

type application struct {
	config config
	db     *mongo.Client
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
