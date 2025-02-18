package application

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g6/internal/application/reqlogger"
	"github.com/maxwelbm/alkemy-g6/internal/application/resources"
	"github.com/maxwelbm/alkemy-g6/internal/controllers"
	_ "github.com/maxwelbm/alkemy-g6/swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger"
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

	a.db, err = sql.Open("mysql", a.cfgDB.FormatDSN())
	if err != nil {
		return err
	}
	// - db: ping
	err = a.db.Ping()
	if err != nil {
		return err
	}
	defer a.db.Close()

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	rt.Use(reqlogger.LoggerMDW(a.db))

	// swagger
	rt.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json")),
	)
	// resources
	rt.Get("/ping", controllers.PingHandler)
	resources.InitInboundOrders(a.db, rt)
	resources.InitEmployees(a.db, rt)
	resources.InitBuyers(a.db, rt)
	resources.InitSellers(a.db, rt)
	resources.InitCarries(a.db, rt)
	resources.InitLocalities(a.db, rt)
	resources.InitProducts(a.db, rt)
	resources.InitWarehouses(a.db, rt)
	resources.InitSections(a.db, rt)
	resources.InitProductRecords(a.db, rt)
	resources.InitPurchaseOrders(a.db, rt)
	resources.InitProductBatches(a.db, rt)

	// run server
	err = http.ListenAndServe(a.Addr, rt)

	return err
}
