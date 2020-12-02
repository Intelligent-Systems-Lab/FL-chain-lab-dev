package main

import (
	"os"

	"github.com/FL-chain-lab-dev/BCFL-demo/bc-app/cmd/bc-appd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
