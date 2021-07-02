package controllers

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	models "store/api/models"
	services "store/api/services"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

var config models.Config

// InitConfig get variable from config file .env
func InitConfig(apiPortConfig string) {
	config.APIPort = apiPortConfig
}

// Initialize connect to the database and wire up routes
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = sql.Open("postgres", DBURI)
	if err != nil {
		log.Fatal("Cannot connect to database : ", err)
	} else {
		log.Println("We are connected to the database ", DbName)
	}
	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Get("/", home)
	a.Get("/view-inventory", a.handleRequest(services.ViewInventory))
	a.Put("/update-inventory", a.handleRequest(services.UpdateInventory))
	a.Post("/checkout-order", a.handleRequest(services.CheckoutOrder))
	a.Post("/checkout-order-withLock", a.handleRequest(services.CheckoutOrderWithLock))

}

func (a *App) RunServer() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("\nServer starting on port %s", config.APIPort)
	log.Fatal(http.ListenAndServe((fmt.Sprintf(":%s", config.APIPort)), handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, this API works.\n"))
}

// reference --> https://github.com/mingrammer/go-todo-rest-api-example/blob/master/app/app.go

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// RequestService represents generalized form of services
type RequestService func(db *sql.DB, config models.Config, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, config, w, r)
	}
}
