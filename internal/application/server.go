package application

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const Title string = `
▗▄▄▄▖▗▄▄▖ ▗▄▄▄▖ ▗▄▄▖ ▗▄▄▖ ▗▄▖  ▗▄▄▖     ▗▄▖ ▗▄▄▖▗▄▄▄▖
▐▌   ▐▌ ▐▌▐▌   ▐▌   ▐▌   ▐▌ ▐▌▐▌       ▐▌ ▐▌▐▌ ▐▌ █  
▐▛▀▀▘▐▛▀▚▖▐▛▀▀▘ ▝▀▚▖▐▌   ▐▌ ▐▌ ▝▀▚▖    ▐▛▀▜▌▐▛▀▘  █  
▐▌   ▐▌ ▐▌▐▙▄▄▖▗▄▄▞▘▝▚▄▄▖▝▚▄▞▘▗▄▄▞▘    ▐▌ ▐▌▐▌  ▗▄█▄▖                                                                                       

`

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the directory with files representing the database
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	return &ServerChi{
		serverAddress:  cfg.ServerAddress,
		loaderFilePath: cfg.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the directory with files representing the database
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	log.Print(Title)
	log.Printf("Starting server at port %s\n", a.serverAddress)

	db, err := loadDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// resources
	buildApiV1SellerRoutes(db, rt)
	buildApiV1WarehousesRoutes(db, rt)
	buildApiV1ProductsRoutes(db, rt)
	buildApiV1SectionsRoutes(db, rt)
	buildApiV1EmployeesRoutes(db, rt)
	buildApiV1BuyerRoutes(db, rt)
	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
