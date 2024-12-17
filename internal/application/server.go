package application

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	sellerController "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	"github.com/maxwelbm/alkemy-g6/internal/loaders"
	sellerRepository "github.com/maxwelbm/alkemy-g6/internal/repository/seller"
	sellerService "github.com/maxwelbm/alkemy-g6/internal/service/seller"
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

	buildSellersRouter(rt)

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

func buildSellersRouter(router *chi.Mux) {
	// ...
	path := fmt.Sprintf(os.Getenv("DB_PATH"), "sellers.json")
	ld := loaders.NewSellerJSONFile(path)
	prods, err := ld.Load()
	if err != nil {
		log.Fatal(err)
	}

	repo := sellerRepository.NewSellerRepository(prods)

	serv := sellerService.NewSellerService(repo)

	cont := sellerController.NewSellerController(serv)

	router.Route("/seller", func(rt chi.Router) {
		rt.Get("/", cont.FindAll())
	})

}
