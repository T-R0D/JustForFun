package cmd

import (
	"github.com/T-R0D/aoc2025/v2/cmd/setup"
	"github.com/T-R0D/aoc2025/v2/cmd/solve"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "aoc2025",
	Short: "AOC2025 is a fun, Christmas-time programming challenge.",
}

func Execute() error {
	return cmd.Execute()
}

func init() {
	cmd.AddCommand(setup.Cmd)
	cmd.AddCommand(solve.Cmd)
}
