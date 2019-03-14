package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var msg string

func main() {
	var cmd = &cobra.Command{
		Use:     "git ng",
		Short:   "git commit with angular style message",
		Long:    `git commit with angular style message`,
		Example: `  git ng -f 'new feature'`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := run(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&msg, "feat", "f", "", "feat")
	cmd.Flags().StringVarP(&msg, "fix", "x", "", "fix")
	cmd.Flags().StringVarP(&msg, "docs", "d", "", "docs")
	cmd.Flags().StringVarP(&msg, "style", "s", "", "style")
	cmd.Flags().StringVarP(&msg, "refactor", "r", "", "refactor")
	cmd.Flags().StringVarP(&msg, "perf", "p", "", "perf")
	cmd.Flags().StringVarP(&msg, "test", "t", "", "refactor")
	cmd.Flags().StringVarP(&msg, "chore", "c", "", "chore")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
