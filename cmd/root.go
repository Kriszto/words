package cmd

import (
	"fmt"
	"os"
	"scrmabled-strings/internal/scrmabledstrings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

func NewRootCmd() *cobra.Command {
	start := time.Now()
	return &cobra.Command{
		Use: "scrambled-strings",
		PostRun: func(cmd *cobra.Command, args []string) {
			elapsed := time.Since(start)
			log.Info().Str("elapsed time", elapsed.String()).Msgf("File reading and processing finished")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dictFilename, _ := cmd.Flags().GetString("dictionary")
			inputFilename, _ := cmd.Flags().GetString("input")
			logLevel, _ := cmd.Flags().GetInt("log-level")

			setupLogging(logLevel)

			log.Debug().Msgf("dictionary: %s, input: %s", dictFilename, inputFilename)
			ret := process(dictFilename, inputFilename)
			for _, r := range ret {
				fmt.Println(r)
			}
			return nil
		},
	}
}

func setupLogging(logLevel int) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	globalLevel := zerolog.Disabled
	if logLevel == 1 {
		globalLevel = zerolog.InfoLevel
	} else if logLevel == 2 {
		globalLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(globalLevel)
	log.Level(globalLevel)
}

func process(dictFilename, inputFilename string) []string {
	ret := make([]string, 0)
	d := scrmabledstrings.NewDictionary(scrmabledstrings.WithFileName(dictFilename))
	d.BuildWords()

	i := scrmabledstrings.NewInput(scrmabledstrings.WithInputFileName(inputFilename), scrmabledstrings.WithDictionary(d))
	t := i.ReadInput()
	r := i.ProcessInput(t)
	for k, n := range r {
		ret = append(ret, fmt.Sprintf("Case #%d: %d", k, n))
	}

	return ret
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = NewRootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringP("dictionary", "d", "", "Dictionary filename")
	rootCmd.Flags().StringP("input", "i", "", "Input filename")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.Flags().IntP("log-level", "l", 0, "log level, 0 (default): no logging, 1: info level, 2: debug level")
	_ = rootCmd.MarkFlagRequired("input")
	_ = rootCmd.MarkFlagRequired("dictionary")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".scrmabled-strings" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".scrmabled-strings")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
