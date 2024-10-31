package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd/http"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/pkg/database"
	"log"
	"log/slog"
)

func main() {

	db := database.LoadDatabase()

	server := http.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port), db)
	if err := server.Run(); err != nil {
		slog.Error("server unable to start")
		log.Fatal()
	}

}
