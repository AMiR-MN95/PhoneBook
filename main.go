package main

import (
	"PhoneBook/cmd"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const description = "PhoneBook Application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),
	)
	if err := root.Execute(); err != nil {
		log.Fatal("Failed to execute root command: \n%v", err)
	}
}
