package main

import (
	"context"
	protos "ctrlshiftv/proto/shorten"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func mmain() {
	sAddress := "localhost:4040"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := protos.NewShortenRequestClient(conn)
	shortLink, err := client.GetShortURL(context.Background(), &protos.LongLink{
		Link: "http://example.com",
	})
	fmt.Println("shortlink", shortLink)
	if err != nil {
		log.Fatalf("Failed to get short link code: %v", e)
	}
}
