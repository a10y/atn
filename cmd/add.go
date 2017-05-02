// Definition of the "add" command for atn.
// Add a message to the store.
package cmd

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/a10y/atn/storage"
	"github.com/google/subcommands"
)

type AddCommand struct{}

func (a *AddCommand) Name() string {
	return "add"
}

func (a *AddCommand) Synopsis() string {
	return "Add a new message to the store. Reads from stdin."
}

func (a *AddCommand) Usage() string {
	return "echo \"hello\" | atn add"
}

func (a *AddCommand) SetFlags(flags *flag.FlagSet) {
	// No flags
}

func (i *AddCommand) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	// Create the storage system
	store := storage.NewDefault()
	// Read the data from stdin, stream it into a file
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "add:", err)
		return subcommands.ExitFailure
	}
	uuid, err := store.AddMessage(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "add:", err)
		return subcommands.ExitFailure
	}
	fmt.Println(uuid)
	return subcommands.ExitSuccess
}
