package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"p2p/proto"
	"p2p/server/blockchain"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ports = []string{"8080", "8081", "8082", "8083"}
var wg = sync.WaitGroup{}

func main() {

	port := flag.String("port", "8080", "server port")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		log.Fatalf("unable to liste  port: %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
		port:       *port,
	})
	fmt.Printf("server run on port: %v", *port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	proto.UnimplementedBlockchainServer
	Blockchain *blockchain.Blockchain
	port       string
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.GetData())
	s.BroadCast(ctx, &proto.BroadcastRequest{
		TransactionHash: block.Hash,
		Count:           0,
	})
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {

		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}

func (s *Server) BroadCast(ctx context.Context, in *proto.BroadcastRequest) (*proto.BroadcastResponse, error) {
	addr := s.port
	fmt.Printf("broadcasting server: %v\n", addr)
	for _, port := range ports {
		if port != addr && in.Count < 3 {
			wg.Add(1)
			go broadCastNode(port, in.TransactionHash, in.Count+1)
		}
	}
	return &proto.BroadcastResponse{
		NodeAddr:        addr,
		TransactionHash: in.TransactionHash,
		Count:           in.Count,
	}, nil
}

func broadCastNode(port string, transactionHash string, count int32) {
	defer wg.Done()

	addr := "localhost:" + port

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	defer conn.Close()

	client := proto.NewBlockchainClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	r, err := client.BroadCast(ctx, &proto.BroadcastRequest{TransactionHash: transactionHash, Count: count})
	if err != nil {
		log.Fatalf("cannot broadcast transaction: %v", err)
	}
	fmt.Printf("nodeaddr: %v, transactionHash: %v, count: %v\n", r.NodeAddr, r.TransactionHash, r.Count)

	wg.Wait()
}
