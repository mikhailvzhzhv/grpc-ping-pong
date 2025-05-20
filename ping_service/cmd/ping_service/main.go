package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	target := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	conn, err := grpc.NewClient(target)
	if err != nil {
		log.Fatalf("cannot create new client: %v", err)
	}

	client := pingpong.NewPingPongerClient(conn)
	req := pingpong.PingRequest{Message: "ping"}

	res, err := client.PingPong(ctx, &req)
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}

	fmt.Printf("response: %s", res.GetMessage())
}
