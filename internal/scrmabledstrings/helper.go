package scrmabledstrings

import (
	"os"

	"github.com/rs/zerolog/log"
)

// openFile opens a file by path
func openFile(filename string) *os.File {
	log.Info().Str("filename", filename).Msg("Opening file")
	file, err := os.Open(filename)
	handleFileError(filename, err)

	return file
}

// createFile creates a file by path
func createFile(filename string) *os.File {
	log.Info().Str("filename", filename).Msg("Creating file")
	file, err := os.Open(filename)
	handleFileError(filename, err)

	return file
}

func handleFileError(filename string, err error) {
	if err != nil {
		log.Fatal().Err(err).Str("filename", filename).Msg("Failed to create file")
		os.Exit(2)
	}
}
