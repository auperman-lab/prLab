package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd/http"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/pkg/database"
	"github.com/auperman-lab/lab2/raft"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {

	nodeID := os.Getenv("NODE_ID")
	portStr := os.Getenv("PORT")
	peersEnv := os.Getenv("PEERS")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("Invalid port number: %s\n", portStr)
		return
	}
	// Split peers from environment variable
	peers := strings.Split(peersEnv, ",")

	db := database.LoadDatabase()

	node := raft.NewNode(nodeID, port, peers)

	server := http.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port), db, node)
	go func(server *http.APIServer) {
		if err := server.Run(); err != nil {
			slog.Error("API server encountered an error", "error", err)
		}
	}(server)

	//wsServer := ws.NewWsServer(fmt.Sprintf(":%s", "2001"))
	//
	//wg.Add(1)
	//go func(wbServer *ws.WsServer) {
	//	defer wg.Done()
	//	wsServer.Run()
	//}(wsServer)

	select {}
}
