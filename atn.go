package main

import (
	"context"
	"flag"
	"os"

	"github.com/a10y/atn/cmd"
	"github.com/google/subcommands"
)

// Commands

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&cmd.InitCommand{}, "")
	subcommands.Register(&cmd.AddCommand{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
