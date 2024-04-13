package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sayCmd)
}

var sayCmd = &cobra.Command{
	Use:   "say [string to say]",
	Short: "Prints provided text to the console",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello:", strings.Join(args, " "))
	},
}
