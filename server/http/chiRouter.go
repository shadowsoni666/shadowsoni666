package server

import (
	"github.com/go-chi/chi/v5"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	auth "shadowsoni666/server/http/auth"
	handlers "shadowsoni666/server/http/handlers"
)

func registerChiMiddlewares(r *chi.Mux, l *zerolog.Logger) {
	// Request logger has middleware.Recoverer and RequestID baked into it.
	r.Use(httplog.RequestLogger(*l),
		middleware.Heartbeat("/ping"),
		middleware.RedirectSlashes,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Bearer", "Set-Cookie", "X-Powered-By", "X-Content-Type-Options"},
			ExposedHeaders:   []string{"Link", "Access-Control-Expose-Headers"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
}

func registerChiHandlers(r *chi.Mux, s handlers.NoteService, t auth.TokenManager, tokenDuration time.Duration, l *zerolog.Logger) {
	r.Post("/register", handlers.RegisterUser(s))
	r.Post("/login", handlers.LoginUser(s, t, tokenDuration))
	r.Post("/logout", handlers.LogoutUser())
	r.Route("/notes", func(r chi.Router) {
		r.Use(auth.AuthMiddleware(t, l))
		r.Post("/create", handlers.CreateNote(s))
		r.Get("/", handlers.GetAllNotesFromUser(s))
		r.Put("/{id}", handlers.UpdateNote(s))
		r.Delete("/{id}", handlers.DeleteNote(s))
	})
}

func NewChiRouter(s handlers.NoteService, symmetricKey string, tokenDuration time.Duration, l *zerolog.Logger) (*chi.Mux, error) {
	pm, err := auth.NewPasetoManager(symmetricKey)
	if err != nil {
		l.Err(err).Msgf("could not create a new PasetoCreator. %v", err)
		return nil, err
	}

	r := chi.NewRouter()
	registerChiMiddlewares(r, l)
	registerChiHandlers(r, s, pm, tokenDuration, l)

	return r, nil
}
