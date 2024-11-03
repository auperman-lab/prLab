package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd/ws"
	"github.com/auperman-lab/lab2/internal/configs"
)

func main() {

	//db := database.LoadDatabase()
	//
	//server := http.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port), db)
	//if err := server.Run(); err != nil {
	//	slog.Error("server unable to start")
	//	log.Fatal()
	//}

	wsServer := ws.NewWsServer(fmt.Sprintf(":%s", configs.Env.Port))

	wsServer.Run()

}
