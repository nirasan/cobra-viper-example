package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("hello world")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
