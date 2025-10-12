package main

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/cmd"
)

func main() {
	if err := cmd.RootCommand.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
