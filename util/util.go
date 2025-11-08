package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

func LogVerbose(cmd *cobra.Command, message string) {
	verbose, _ := cmd.Flags().GetBool("verbose")
	if verbose {
		fmt.Print(message)
	}
}
