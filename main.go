package main

import (
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
		postgresURI := "postgres://tester:Apasswd@localhost?sslmode=disable"
		if os.Getenv("POSTGRES_URI") != "" {
			postgresURI = os.Getenv("POSTGRES_URI")
		}
		if os.Getenv("POSTGRES_URI") == "" {
			postgresUser := "tester"
			postgresUser = os.Getenv("POSTGRES_USER")
			postgresPassword := "passwd"
			postgresPassword = os.Getenv("POSTGRES_PASSWORD")
			postgresHost := "localhost"
			postgresHost = os.Getenv("POSTGRES_HOST")
			postgresDB := "test"
			postgresDB = os.Getenv("POSTGRES_DB")
			// Switch to using the above at some point
			postgresConnStr := fmt.Sprintf("postgres://%d:%d@%d/%d?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresDB)
			postgresURI = postgresConnStr
			log.Println("postgresURI:", postgresURI)
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
	// TODO: change to optionally use env var
	sAddress := os.Getenv("URLSHORT_ADDR")
	if sAddress == "" {
		sAddress = "localhost:4040"
	}
	conn, err := grpc.Dial(
		sAddress,
		grpc.WithInsecure(),
	)
	// Can help with connection is closing errors sometimes
	//grpc.FailOnNonTempDialError(true),
	//gr	grpc.WithBlock(),
	//gr)
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}
	// makes passing the connection fail so it's commented out for now
	//defer conn.Close()

	client := protos.NewShortenRequestClient(conn)
	paste.SaveClient(client)
	paste.SaveConn(conn)
}

func startHTTP() {
	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "8080"
	}
	port = ":" + port

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
		fmt.Printf("Listening on port %s\n", port)
		errs <- http.ListenAndServe(port, r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s\n", <-errs)
}

func main() {
	go startGRPC()
	defer startHTTP()
}
