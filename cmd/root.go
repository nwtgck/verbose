package cmd

import (
	"github.com/nwtgck/verbose"
	"github.com/spf13/cobra"
	"os"
)

var ampN int
var usesDecode bool

func init() {
	cobra.OnInitialize()
	RootCmd.Flags().IntVarP(&ampN,  "n", "n", 10, "The number of byte replication")
	RootCmd.Flags().BoolVarP(&usesDecode,  "decode", "d", false, "Decode or not")
}

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "verbose",
	Long:  "Verbose input",
	RunE: func(cmd *cobra.Command, args []string) error {
		if usesDecode {
			return verbose.Decode(ampN, os.Stdin, os.Stdout)
		} else {
			return verbose.Encode(ampN, os.Stdin, os.Stdout)
		}
	},
}
