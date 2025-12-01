package solve

import (
	"fmt"
	"os"

	"github.com/T-R0D/aoc2025/v2/internal/solution"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve one part of one day's problem for Advent of Code",
	RunE:  runE,
}

var (
	inputPath string
	day       int
	part      int
)

func init() {
	Cmd.Flags().StringVar(&inputPath, "input", "", "Path to the input file")
	Cmd.MarkFlagFilename("input")
	Cmd.MarkFlagRequired("input")

	Cmd.Flags().IntVar(&day, "day", 0, "The day to solve (1-12)")
	Cmd.MarkFlagRequired("day")

	Cmd.Flags().IntVar(&part, "part", 0, "The part of the day to solve (1,2)")
	Cmd.MarkFlagRequired("part")
}

func runE(cmd *cobra.Command, args []string) error {
	if info, err := os.Stat(inputPath); err != nil {
		return fmt.Errorf("input file path is not valid: %w", err)
	} else if info.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", inputPath)
	}

	if day < 1 || 25 < day {
		return fmt.Errorf("'%d' is not a valid advent day (must be 1-25)", day)
	}

	if part < 1 || 2 < part {
		return fmt.Errorf("'%d' is not a valid problem part (must be 1-2)", part)
	}

	return solution.Run(inputPath, day, part)
}
