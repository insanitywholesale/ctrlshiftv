package main

import (
	"context"
	h "ctrlshiftv/api"
	"ctrlshiftv/paste"
	protos "ctrlshiftv/proto/shorten"
	mockrepo "ctrlshiftv/repo/mock"
	psql "ctrlshiftv/repo/postgres"
	sqlite "ctrlshiftv/repo/sqlite"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func httpPort() string {
	port := "8080"
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
		postgresURI := "postgres://tester:Apasswd@localhost?sslmode=disable"
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

func startGRPC() {
	sAddress := "localhost:4040"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := protos.NewShortenRequestClient(conn)
	shortLink, err := client.GetShortURL(context.Background(), &protos.LongLink{
		Link: "http://distro.watch",
	})
	fmt.Println("shortlink", shortLink)
	if err != nil {
		log.Fatalf("Failed to get short link code: %v", e)
	}

	// exit so the following don't run
	//os.Exit(0)
}

func startHTTP() {
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
		fmt.Printf("Listening on port %s\n", httpPort())
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s\n", <-errs)
}

func main() {
	startHTTP()
}
