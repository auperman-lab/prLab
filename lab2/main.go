package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/pkg/database"
	"log/slog"
)

func main() {

	db := database.LoadDatabase()

	server := cmd.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port), db)
	if err := server.Run(); err != nil {
		slog.Error("server unable to start")
	}

}
