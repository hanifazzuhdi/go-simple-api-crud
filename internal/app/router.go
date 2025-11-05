package app

import (
	"net/http"
	"simple-api/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(nationalityHandler *handler.NationalityHandler, customerHandler *handler.CustomerHandler, familyListHandler *handler.FamilyListHandler) *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/nationalities", nationalityHandler.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.GetById).Methods(http.MethodGet)
	api.HandleFunc("/nationalities", nationalityHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.Update).Methods(http.MethodPut)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.Delete).Methods(http.MethodDelete)

	api.HandleFunc("/customers", customerHandler.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/customers/{id}", customerHandler.GetById).Methods(http.MethodGet)
	api.HandleFunc("/customers", customerHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/customers/{id}/families/sync", customerHandler.SyncCustomerFamilies).Methods(http.MethodPost)
	api.HandleFunc("/customers/{id}", customerHandler.Update).Methods(http.MethodPut)
	api.HandleFunc("/customers/{id}", customerHandler.Delete).Methods(http.MethodDelete)

	api.HandleFunc("/families", familyListHandler.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/families/{id}", familyListHandler.GetById).Methods(http.MethodGet)
	api.HandleFunc("/families", familyListHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/families/{id}", familyListHandler.Update).Methods(http.MethodPut)
	api.HandleFunc("/families/{id}", familyListHandler.Delete).Methods(http.MethodDelete)

	router.Use(mux.CORSMethodMiddleware(router))

	return router
}
