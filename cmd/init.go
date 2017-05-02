package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/a10y/atn/storage"
	"github.com/google/subcommands"
)

type InitCommand struct{}

func (i *InitCommand) Name() string {
	return "init"
}

func (i *InitCommand) Synopsis() string {
	return "Initialize the storage system"
}

func (i *InitCommand) Usage() string {
	return "init"
}

func (i *InitCommand) SetFlags(flags *flag.FlagSet) {
	// No flags
}

// Execute the init command
func (i *InitCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	store := storage.NewDefault()
	if err := store.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "Error encountered when initializing file system: ", err)
		return subcommands.ExitFailure
	} else {
		return subcommands.ExitSuccess
	}
}
