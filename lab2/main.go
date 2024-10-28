package main

import (
	"fmt"
	"github.com/auperman-lab/lab2/cmd"
	"github.com/auperman-lab/lab2/internal/configs"
	"log/slog"
)

func main() {

	server := cmd.NewAPIServer(fmt.Sprintf(":%s", configs.Env.Port))
	if err := server.Run(); err != nil {
		slog.Error("server unable to start")
	}

}
