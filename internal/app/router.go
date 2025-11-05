package app

import (
	"net/http"
	"simple-api/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(nationalityHandler *handler.NationalityHandler, customerHandler *handler.CustomerHandler, familyListHandler *handler.FamilyListHandler) *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/nationalities", nationalityHandler.GetAll).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.GetById).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/nationalities", nationalityHandler.Create).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.Update).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/nationalities/{id}", nationalityHandler.Delete).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/customers", customerHandler.GetAll).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/customers/{id}", customerHandler.GetById).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/customers", customerHandler.Create).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/customers/{id}/families/sync", customerHandler.SyncCustomerFamilies).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/customers/{id}", customerHandler.Update).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/customers/{id}", customerHandler.Delete).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/families", familyListHandler.GetAll).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/families/{id}", familyListHandler.GetById).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/families", familyListHandler.Create).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/families/{id}", familyListHandler.Update).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/families/{id}", familyListHandler.Delete).Methods(http.MethodDelete, http.MethodOptions)

	router.Use(handleCorsMiddleware)

	return router
}

func handleCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-requested-with")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
