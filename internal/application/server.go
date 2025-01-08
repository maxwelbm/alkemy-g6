package application

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/application/resources"
)

const Title string = `
▗▄▄▄▖▗▄▄▖ ▗▄▄▄▖ ▗▄▄▖ ▗▄▄▖ ▗▄▖  ▗▄▄▖     ▗▄▖ ▗▄▄▖▗▄▄▄▖
▐▌   ▐▌ ▐▌▐▌   ▐▌   ▐▌   ▐▌ ▐▌▐▌       ▐▌ ▐▌▐▌ ▐▌ █  
▐▛▀▀▘▐▛▀▚▖▐▛▀▀▘ ▝▀▚▖▐▌   ▐▌ ▐▌ ▝▀▚▖    ▐▛▀▜▌▐▛▀▘  █  
▐▌   ▐▌ ▐▌▐▙▄▄▖▗▄▄▞▘▝▚▄▄▖▝▚▄▞▘▗▄▄▞▘    ▐▌ ▐▌▐▌  ▗▄█▄▖                                                                                       

`

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// Db is the database configuration.
	DB *mysql.Config
	// ServerAddress is the address where the server will be listening
	Addr string
	// LoaderFilePath is the path to the directory with files representing the database
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	return &ServerChi{
		cfgDB:          cfg.DB,
		Addr:           cfg.Addr,
		loaderFilePath: cfg.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// Db is the database configuration.
	cfgDB *mysql.Config
	// serverAddress is the address where the server will be listening
	Addr string
	// db is the database connection.
	db *sql.DB
	// loaderFilePath is the path to the directory with files representing the database
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	log.Print(Title)
	log.Printf("Starting server at port %s\n", a.Addr)

	// deprecated: loads json files for old data access
	jsonDB, err := loadDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	a.db, err = sql.Open("mysql", a.cfgDB.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return
	}
	defer a.db.Close()

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// resources
	buildApiV1SellerRoutes(jsonDB, rt)
	buildApiV1WarehousesRoutes(jsonDB, rt)
	buildApiV1ProductsRoutes(jsonDB, rt)
	buildApiV1SectionsRoutes(jsonDB, rt)
	buildApiV1EmployeesRoutes(jsonDB, rt)
	resources.InitLocalities(a.db, rt)
	// run server
	err = http.ListenAndServe(a.Addr, rt)
	return
}
