package generate

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stormcat24/go-assign/pkg/cli"
)

type generator struct {
}

func NewCommand() *cobra.Command{
	v := generator{}
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate an inherited struct with go-assign tag.",
		RunE: cli.WithContext(v.run),
	}
	return cmd
}

func (g *generator) run(ctx context.Context) error {
	fmt.Println("This is generate")
	return nil
}