package cmd

import (
	"PhoneBook/internal/api/http"
	"PhoneBook/internal/config"
	logger2 "PhoneBook/pkg/logger"
	"PhoneBook/pkg/token"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.run(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run PhoneBook server",
		Run:   run,
	}
}

func (cmd *Server) run(cfg *config.Config, trap chan os.Signal) {
	logger := logger2.NewZap(cfg.Logger)

	token, err := token.New(cfg.Token)
	if err != nil {
		logger.Panic("Error creating token object", zap.Error(err))
	}

	server := http.New(logger, token)
	go server.Server(8080)

	filed := zap.String("signal tap", (<-trap).String())
	logger.Info("Exiting by receiving Signal", filed)
}
