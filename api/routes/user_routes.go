package routes

import (
	"github.com/gorilla/mux"
	handlers "github.com/lahiruudayakumara/bet_settlement_engine/api/handler"
	"net/http"
)

func RegisterUserRoutes(mux *mux.Router) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.UserHandler(w, r)
		case http.MethodGet:
			handlers.GetAllUsersHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
