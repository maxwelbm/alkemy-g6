package sqlconfig

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/melisource/fury_go-toolkit-secrets/pkg/secrets"
)

func NewConn(env string) (config *mysql.Config) {
	switch env {
	case "production":
		fmt.Println("hi")
		config = prodConfig()
	default:
		config = devConfig()
	}

	return
}

func devConfig() *mysql.Config {
	return &mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_ADDR"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
}

func prodConfig() *mysql.Config {
	client, err := secrets.NewClient()
	if err != nil {
		panic(err)
	}

	usr, ok := client.
		GetSecret("DB_MYSQL_<CLUSTER>_<SCHEMA_NAME>_<SCHEMA_NAME>_WPROD_USER")
	if !ok {
		panic("write db user not found")
	}

	pass, ok := client.
		GetSecret("DB_MYSQL_<CLUSTER>_<SCHEMA_NAME>_<SCHEMA_NAME>_WPROD")
	if !ok {
		panic("write db pass not found")
	}

	return &mysql.Config{
		User:                 usr,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_MYSQL_<CLUSTER>_<SCHEMA_NAME>_<SCHEMA_NAME>_ENDPOINT"),
		DBName:               "<SCHEMA_NAME>",
		Timeout:              100 * time.Millisecond,
		ReadTimeout:          100 * time.Millisecond,
		WriteTimeout:         100 * time.Millisecond,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
}
