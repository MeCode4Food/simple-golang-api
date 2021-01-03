package server

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-backend/server/handlers"
	"simple-backend/server/services/db"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
	DB     db.DB
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// InitialiseRoutes export
func (app *App) InitialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/cat/{name}", handlers.GetOneCatHandler(&app.DB)).Methods("GET")
	app.Router.HandleFunc("/cat", handlers.GetAllCatHandler(&app.DB)).Methods("GET")
	app.Router.HandleFunc("/cat", handlers.PutCatHandler(&app.DB)).Methods("PUT")
	app.Router.HandleFunc("/cat", handlers.PostCatHandler(&app.DB)).Methods("POST")
	app.Router.HandleFunc("/cat/{name}", handlers.DeleteCatHandler(&app.DB)).Methods("DELETE")
	app.Router.HandleFunc("/", handlers.PingHandler)
}

// InitialiseDatabase Export
func (app *App) InitialiseDatabase() {
	app.DB = db.CreateCatDatabase()
	app.DB.InitialiseSampleData()
}

// Run export
func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
