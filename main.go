package main

import (
	pr "ctrlshiftv/repo/postgres"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("main")
	//repo, err := pr.NewPostgresRepo("postgres://angle:Apasswd@localhost/angle?sslmode=disable")
	repo, err := pr.NewPostgresRepo("postgres://test:Apasswd@localhost/test?sslmode=disable")
	fmt.Println("got repo", repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	os.Exit(0)
}
