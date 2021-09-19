package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"github.com/ljgago/api-rest/internal/controller"
	"github.com/ljgago/api-rest/internal/core/port"
	"github.com/ljgago/api-rest/internal/core/service"
	"github.com/ljgago/api-rest/internal/repository"
	"github.com/ljgago/api-rest/internal/routes"
)

type configDB struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func initDB() *sql.DB {
	const (
		host     = "127.0.0.1"
		port     = "5432"
		user     = "postgres"
		password = "postgres"
		dbname   = "book"
	)

	psqlConfig := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConfig)
	if err != nil {
		panic(err)
	}

	return db
}

func initialice(db *sql.DB) port.BookController {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	return bookController
}

func main() {
	var serverPort string
	flag.StringVar(&serverPort, "port", "3000", "server port")

	flag.Parse()

	db := initDB()
	defer db.Close()

	bookController := initialice(db)

	r := chi.NewRouter()
	routes.BookRoutes(r, bookController)

	log.Println("Server on port :" + serverPort)
	http.ListenAndServe(":"+serverPort, r)
}
