package main

import (
	"context"
	"flag"
	"log"
	"p2p/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.BlockchainClient

func main() {
	wg := sync.WaitGroup{}

	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	defer conn.Close()

	client = proto.NewBlockchainClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	if *addFlag {

		addBlock(ctx)
	}
	if *listFlag {
		listBlock(ctx)
	}

	wg.Wait()

}

func addBlock(ctx context.Context) {
	block, err := client.AddBlock(ctx, &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("unable to add block: %v", err)
	}
	log.Printf("new block hash: %s", block.Hash)

}

func listBlock(ctx context.Context) {
	list, err := client.GetBlockchain(ctx, &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to add list: %v", err)
	}
	log.Println("Blocks:")
	for _, b := range list.Blocks {

		log.Printf("hash: %s, prev block hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}

}
