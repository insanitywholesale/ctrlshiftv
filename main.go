package main

import (
	mockrepo "ctrlshiftv/repo/mock"
	//psql "ctrlshiftv/repo/postgres"
	//sqlite "ctrlshiftv/repo/postgres"
	h "ctrlshiftv/api"
	"ctrlshiftv/paste"
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
	//sample of connecting to postgres -- useless since psql implementation not done
	//repo, err := pr.NewPostgresRepo("postgres://tester:Apasswd@localhost/test?sslmode=disable")
	//repo, err := mockrepo.NewMockRepo()
	//p, err := repo.Find("1234")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("code 1234's content:\n------------------------\n	%s\n------------------------\n", p.Content)
	//fmt.Println("got repo", repo)
	//if err != nil {
	//	log.Fatal(err)
	//}
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

func StartService() {
	repo, err := mockrepo.NewMockRepo()
	if err != nil {
		log.Fatal(err)
	}

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
