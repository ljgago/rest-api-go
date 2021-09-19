package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"github.com/ljgago/api-rest/internal/book"
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

func initialice(db *sql.DB) book.Controller {
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookController := book.NewController(bookService)
	return bookController
}

func main() {
	var port string
	flag.StringVar(&port, "port", "3000", "server port")

	flag.Parse()

	db := initDB()
	defer db.Close()

	bookController := initialice(db)

	r := chi.NewRouter()
	book.Routes(r, bookController)

	log.Println("Server on port :" + port)
	http.ListenAndServe(":"+port, r)
}
