package main

import (
    "log"
    context "golang.org/x/net/context"
    "os"
    pb "github.com/toukii/grpc/hello"
    grpc "google.golang.org/grpc"
)


func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewHelloServiceClient(conn)

    // Contact the server and print out its response.
    name := "toukii"
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Reply)
}