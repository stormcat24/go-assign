package cli

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

type Runner func(ctx context.Context) error

func WithContext(runner Runner ) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(ch)
		return runWithContext(cmd, ch, runner)
	}
}

func runWithContext(cmd *cobra.Command, signalCh <-chan os.Signal, runner Runner) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case s := <-signalCh:
			log.Printf("Stopping because it received a signal: %d", s)
			cancel()
		case <-ctx.Done():
		}
	}()
	return runner(ctx)
}