package application

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maxwelbm/alkemy-g6b/internal/loaders"
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

	//endpoints
	rt.Route("api/v1", func(r chi.Router) {
		rt.Route("/employees", func(r chi.Router) {
			r.Get("/", ctrl.GetEmployees)
		})
	})

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

func loadProductRepository() (repo *employees_repository.Employees) {
	// loads products from products.json file
	path := fmt.Sprintf(os.Getenv("DB_PATH"), "employees.json")
	ld := loaders.NewEmployeesJSONFile(path)
	emp, err := ld.Load()
	if err != nil {
		log.Fatal(err)
	}

	repo = employees_repository.NewEmployeesRepository(emp)

	return
}
