package cmd

import (
	"github.com/T-R0D/aoc2024/v2/cmd/setup"
	"github.com/T-R0D/aoc2024/v2/cmd/solve"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "aoc2024",
	Short: "AOC2024 is a fun, Christmas-time programming challenge.",
}

func Execute() error {
	return cmd.Execute()
}

func init() {
	cmd.AddCommand(setup.Cmd)
	cmd.AddCommand(solve.Cmd)
}
