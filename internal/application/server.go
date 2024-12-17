package application

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
	fmt.Print(Title)
	fmt.Printf("Starting server at port %s\n", a.serverAddress)
	
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	buildApiV1WarehousesRoutes(rt)

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}

const Title string = `
▗▄▄▄▖▗▄▄▖ ▗▄▄▄▖ ▗▄▄▖ ▗▄▄▖ ▗▄▖  ▗▄▄▖     ▗▄▖ ▗▄▄▖▗▄▄▄▖
▐▌   ▐▌ ▐▌▐▌   ▐▌   ▐▌   ▐▌ ▐▌▐▌       ▐▌ ▐▌▐▌ ▐▌ █  
▐▛▀▀▘▐▛▀▚▖▐▛▀▀▘ ▝▀▚▖▐▌   ▐▌ ▐▌ ▝▀▚▖    ▐▛▀▜▌▐▛▀▘  █  
▐▌   ▐▌ ▐▌▐▙▄▄▖▗▄▄▞▘▝▚▄▄▖▝▚▄▞▘▗▄▄▞▘    ▐▌ ▐▌▐▌  ▗▄█▄▖                                                                                       

`

// func buildWarehousesRouter(rt *chi.Mux) (err error) {
//     path := fmt.Sprintf("%s%s", os.Getenv("DB_PATH"), "warehouses.json")
//     ld := loaders.NewWarehouseJSONFile(path)
//     warehouses, err := ld.Load()
//     if err != nil {
//         log.Fatal(err)
// 		return
//     }

//     repo := warehouse_repository.NewWarehouse(warehouses)
// 	service := service.NewWarehouseDefault(repo)
// 	controller := controllers.NewWarehouseDefault(service)

// 	rt.Route("/api/v1/warehouses", func(r chi.Router ) {
// 		r.Get("/", controller.GetAll())
// 	})

//     return
// }