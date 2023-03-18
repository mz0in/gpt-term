package cmd

import (
	"fmt"
	"os/exec"

	"github.com/collinvandyck/gpterm/lib/cmdkit"
	"github.com/spf13/cobra"
)

func sqlcCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "sqlc",
		Short:   "Generates sqlc queries",
		PreRunE: Deps().RunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdkit.ChdirProject("db")
			ec := exec.Command(
				"sqlc",
				"generate",
				"--file",
				"sqlc.yaml")
			out, err := ec.CombinedOutput()
			if err != nil {
				fmt.Println(string(out))
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}
}
