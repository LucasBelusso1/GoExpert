package main

import (
	"log"
	"net/http"

	"github.com/LucasBelusso1/GoExpert/7-APIS/configs"
	_ "github.com/LucasBelusso1/GoExpert/7-APIS/docs"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/entity"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/infra/database"
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lucas belusso
// @contact.url    http://www.lucasbelusso.com.br
// @contact.email  lucasbelusso@lucasbelusso.com.br

// @license.name   Lucas Belusso License
// @license.url    http://www.lucasbelusso.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("TokenAuth", cfg.TokenAuth))
	r.Use(middleware.WithValue("JWTExpiresIn", cfg.JWTExpiresIn))
	// r.Use(LogRequest) // Middleware created manually

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.Delete)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
