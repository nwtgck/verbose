package main

import (
	"bytes"
	"fmt"
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
		return verbose(ampN)
	},
}

func verbose(ampN int) error {
	// One byte slice
	bs := make([]byte, 1)
	for {
		// Read one byte
		n, err := os.Stdin.Read(bs)
		// Quit if EOF
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		// Write amplified bytes
		os.Stdout.Write(bytes.Repeat(bs, ampN))
	}
	return nil
}

func main () {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
