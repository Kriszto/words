package cmd

import (
	"fmt"
	"os"
	"scrmabled-strings/internal/scrmabledstrings"

	"github.com/Ak-Army/xlog"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scrmabled-strings",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		dictFilename, _ := cmd.Flags().GetString("dictionary")
		inputFilename, _ := cmd.Flags().GetString("input")
		verbose, _ := cmd.Flags().GetBool("verbose")
		color.Green("d: %s, i: %s", dictFilename, inputFilename)

		d := scrmabledstrings.NewDictionary(scrmabledstrings.WithFileName(dictFilename))
		d.BuildWords()
		if verbose {
			xlog.Debug(d)
		}
		i := scrmabledstrings.NewInput(inputFilename, scrmabledstrings.WithDictionary(d))
		i.ProcessFile()
		//x := d.BuildAllScrambledWord()
		//xlog.Debug(x)
		//scrmabledstrings.GoPerm()
		n, l := d.Result()
		fmt.Printf("Case #1: %d (%d)\n", n, l)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringP("dictionary", "d", "", "Dictionary filename")
	rootCmd.Flags().StringP("input", "i", "", "Input filename")
	rootCmd.Flags().StringP("verbose", "v", "", "Verbose output")
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
