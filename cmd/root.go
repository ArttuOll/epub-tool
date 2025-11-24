package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "epub-tool",
	Args:  cobra.MinimumNArgs(1),
	Short: "Removes annoying hard-coded CSS styles from epub files.",
	Long: `epub-tool removes from epub files CSS styles that harm the reading experience.
	
Currently this means: 
	1. Removing font size declarations that make it impossible to adjust the font size in your ebook reader
	2. Removing text color declarations that cause the text color to remain black when the reader switches the background to black in dark mode`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if dryRun, _ := cmd.Flags().GetBool("dryRun"); dryRun {
			cmd.Flags().Set("verbose", "true")
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return CleanupE(cmd, args[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.epub-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	rootCmd.Flags().StringP("outputFileName", "o", "", "name of the cleaned output file")
	rootCmd.Flags().BoolP("dryRun", "d", false, "print changes that would be made, but don't write them to disk")
	rootCmd.Flags().BoolP("removeBackgroundColors", "b", false, "additionally remove all background-color declarations. This might help if some graphic elements aren't visible in dark mode.")
}
