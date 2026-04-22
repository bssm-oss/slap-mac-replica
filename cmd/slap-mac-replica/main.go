package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bssm-oss/slap-mac-replica/internal/app"
	"github.com/bssm-oss/slap-mac-replica/internal/config"
)

var version = "dev"

func main() {
	cfg, err := config.Parse(os.Args[1:])
	if err != nil {
		switch {
		case errors.Is(err, config.ErrHelpRequested):
			fmt.Print(config.Usage(version))
			return
		case errors.Is(err, config.ErrVersionRequested):
			fmt.Println(version)
			return
		default:
			fmt.Fprintf(os.Stderr, "slap-mac-replica: %v\n\n%s", err, config.Usage(version))
			os.Exit(2)
		}
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.Execute(ctx, cfg, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "slap-mac-replica: %v\n", err)
		os.Exit(1)
	}
}
