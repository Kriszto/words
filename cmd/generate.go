package cmd

import (
	"scrmabled-strings/internal/scrmabledstrings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a set of test cases from original tests.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msgf("generate called")
		g := scrmabledstrings.NewGenerator()
		g.ProcessData("ts1")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
