package cmd

import (
	"github.com/nwtgck/verbose"
	"github.com/spf13/cobra"
	"os"
)

var ampN int

func init() {
	cobra.OnInitialize()
	RootCmd.Flags().IntVarP(&ampN,  "n", "n", 10, "The number of byte replication")
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "verbose",
	Long:  "Verbose input",
	RunE: func(cmd *cobra.Command, args []string) error {
		return verbose.Encode(ampN, os.Stdin, os.Stdout)
	},
}
