package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "setup",
	Short: "Generate or tear down files for Advent of Code problems",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagDelete {
			return deleteSolutionSkeletons(flagTargetDir)
		}

		return generateSolutionSkeletons(flagTargetDir)
	},
}

var (
	flagTargetDir string
	flagDelete    bool
)

func init() {
	Cmd.Flags().StringVar(
		&flagTargetDir,
		"targetDir",
		"",
		"Directory in which to create the artifacts")
	Cmd.MarkFlagRequired("targetDir")

	Cmd.Flags().BoolVar(
		&flagDelete,
		"delete",
		false,
		"Delete artifacts that were created from a previous run")
}

func generateSolutionSkeletons(targetDir string) error {
	for i := 1; i <= 12; i += 1 {
		if err := createDaySkeleton(targetDir, i); err != nil {
			deleteSolutionSkeletons(targetDir)
			return fmt.Errorf("generating day %d: %w", i, err)
		}
	}
	return nil
}

func deleteSolutionSkeletons(targetDir string) error {
	for i := 1; i <= 12; i += 1 {
		folderName := getPackageName(i)
		deletePath := filepath.Join(targetDir, folderName)
		os.RemoveAll(deletePath)
	}

	return nil
}

func getPackageName(day int) string {
	return fmt.Sprintf("day%02d", day)
}

func createDaySkeleton(targetDir string, day int) error {
	packageName := getPackageName(day)
	dayFolderPath := filepath.Join(targetDir, packageName)

	if err := os.Mkdir(dayFolderPath, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create folder for %s: %w", packageName, err)
	}

	solverFilePath := filepath.Join(dayFolderPath, "solver.go")
	solverFileContents := fmt.Sprintf(solverSkeletonFile, packageName)
	if err := writeFile(solverFilePath, solverFileContents); err != nil {
		return fmt.Errorf("writing solver file: %w", err)
	}

	testFilePath := filepath.Join(dayFolderPath, "solver_test.go")
	testFileContents := fmt.Sprintf(testSkeletonFile, packageName)
	if err := writeFile(testFilePath, testFileContents); err != nil {
		return fmt.Errorf("writing solver file: %w", err)
	}

	return nil
}

func writeFile(path string, contents string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating file '%s': %w", path, err)
	}

	if _, err := file.WriteString(contents); err != nil {
		return fmt.Errorf("writing file '%s': %w", path, err)
	}

	cmd := exec.Command("gofmt", path)
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	return nil
}

const solverSkeletonFile = `package %s

import (
	"fmt"
)

type Solver struct{}

func (this *Solver) SolvePartOne(input string) (string, error) {
	return "", fmt.Errorf("part 1 not implemented")
}

func (this *Solver) SolvePartTwo(input string) (string, error) {
	return "", fmt.Errorf("part 2 not implemented")
}
`

const testSkeletonFile = `package %s

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartOne(tc.input)
			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %%s, wanted %%s", result, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct{
		name     string
		input    string
		expected string
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			solver := Solver{}

			result, err := solver.SolvePartTwo(tc.input)
			if err != nil {
				t.Error("unable to complete solution", err)
			}

			if result != tc.expected {
				t.Errorf("got %%s, wanted %%s", result, tc.expected)
			}
		})
	}
}
`
