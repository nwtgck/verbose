package main

import (
	"fmt"
	"github.com/nwtgck/verbose/cmd"
	"os"
)

func main () {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
