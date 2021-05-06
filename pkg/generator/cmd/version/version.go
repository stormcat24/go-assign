package version

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stormcat24/go-assign/pkg/cli"
	"github.com/stormcat24/go-assign/pkg/version"
)

type versionPrinter struct {
}

func NewCommand() *cobra.Command {
	v := versionPrinter{}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print go-assign version",
		RunE: cli.WithContext(v.run),
	}
	return cmd
}

func (v *versionPrinter) run(ctx context.Context) error {
	fmt.Println(version.Get().String())
	return nil
}