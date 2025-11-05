package app

import (
	"context"
	"net/http"
	"simple-api/internal/handler"
	"simple-api/internal/repository/postgres"
	"simple-api/internal/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type App struct {
	db     *pgx.Conn
	router *mux.Router
}

func New(ctx context.Context, dbUrl string) *App {
	dbConn := GetDbConnection(ctx, dbUrl)

	// repositories
	nationalityRepository := postgres.NewNationalityRepository(dbConn)
	customerRepository := postgres.NewCustomerRepository(dbConn)
	familyListRepository := postgres.NewFamilyRepository(dbConn)

	// services
	nationalityService := service.NewNationalityService(nationalityRepository)
	customerService := service.NewCustomerService(customerRepository, familyListRepository)
	familyListService := service.NewFamilyListService(familyListRepository)

	// handlers
	nationalityHandler := handler.NewNationalityHandler(nationalityService)
	customerHandler := handler.NewCustomerHandler(customerService)
	familyListHandler := handler.NewFamilyListHandler(familyListService)

	router := NewRouter(nationalityHandler, customerHandler, familyListHandler)

	return &App{
		db:     dbConn,
		router: router,
	}
}

func (app *App) Run(port string) error {
	server := http.Server{
		Addr:    ":" + port,
		Handler: app.router,
	}

	return server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.db.Close(ctx)
}
