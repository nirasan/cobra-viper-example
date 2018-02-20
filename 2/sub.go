package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// サブコマンドの定義
var subCmd = &cobra.Command{
	Use: "sub",
	Run: func(c *cobra.Command, args []string) {
		if name, err := c.PersistentFlags().GetString("name"); err == nil {
			// 共通フラグ debug をサブコマンドからも利用できる
			if debug {
				log.Println("name:", name)
			} else {
				fmt.Println("name:", name)
			}
		}
	},
}

func init() {
	// サブコマンドのフラグ定義
	subCmd.PersistentFlags().String("name", "john", "sub command string flag test")
	// サブコマンドをルートコマンドに登録
	rootCmd.AddCommand(subCmd)
}
