package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd/http"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/pkg/database"
	"log/slog"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	db := database.LoadDatabase()

	server := http.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port), db)
	wg.Add(1)
	go func(server *http.APIServer) {
		defer wg.Done()
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

	wg.Wait()
}
