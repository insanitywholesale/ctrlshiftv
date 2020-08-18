package main

import (
	h "ctrlshiftv/api"
	"ctrlshiftv/paste"
	mockrepo "ctrlshiftv/repo/mock"
	psql "ctrlshiftv/repo/postgres"
	sqlite "ctrlshiftv/repo/sqlite"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	StartService()
	os.Exit(0)
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() paste.PasteRepo {
	switch os.Getenv("URL_DB") {
	case "sqlite":
		sqliteURL := os.Getenv("SQLITE_URL")
		repo, err := sqlite.NewSQLiteRepo(sqliteURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		//postgresURL := os.Getenv("POSTGRES_URL")
		//postgresUser := os.Getenv("POSTGRES_USER")
		//postgresPassword := os.Getenv("POSTGRES_PASSWORD")
		//postgresHost := os.Getenv("POSTGRES_HOST")
		//postgresDB := os.Getenv("POSTGRES_DB")
		// Switch to using the above at some point
		postgresURI := "postgres://tester:Apasswd@localhost/test?sslmode=disable"
		if os.Getenv("POSTGRES_URI") != "" {
			postgresURI = os.Getenv("POSTGRES_URI")
		}
		repo, err := psql.NewPostgresRepo(postgresURI)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	default:
		repo, err := mockrepo.NewMockRepo()
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

func StartService() {
	repo := chooseRepo()
	service := paste.NewPasteService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}
