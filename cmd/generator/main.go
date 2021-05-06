package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/stormcat24/go-assign/pkg/generator/cmd/generate"
	"github.com/stormcat24/go-assign/pkg/generator/cmd/version"
)

func main() {
	root := cobra.Command{}
	root.AddCommand(generate.NewCommand())
	root.AddCommand(version.NewCommand())
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
