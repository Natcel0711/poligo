package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Natcel0711/poligo/config"
	"github.com/Natcel0711/poligo/internal/storage"
	"github.com/Natcel0711/poligo/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	db, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		panic(err.Error())
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy!"))
	})
	userStore := user.NewUserStorage(db)
	userController := user.NewUserController(userStore)
	user.AddUserRoutes(r, userController)
	fmt.Printf("listening on %s", "0.0.0.0:"+env.PORT)
	http.ListenAndServe("0.0.0.0:"+env.PORT, r)
}
