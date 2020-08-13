package main

import (
	mockrepo "ctrlshiftv/repo/mock"
	//psql "ctrlshiftv/repo/postgres"
	//sqlite "ctrlshiftv/repo/postgres"
	"fmt"
	"log"
	"os"
)

func main() {
	//sample of connecting to postgres -- useless since psql implementation not done
	//repo, err := pr.NewPostgresRepo("postgres://tester:Apasswd@localhost/test?sslmode=disable")
	repo, err := mockrepo.NewMockRepo()
	p, err := repo.Find("1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("code 1234's content:\n------------------------\n	%s\n------------------------\n", p.Content)
	fmt.Println("got repo", repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	os.Exit(0)
}
