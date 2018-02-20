package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// ルートコマンドの定義
var rootCmd = &cobra.Command{
	Use: "app",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("debug:", debug)
	},
}

// 共通フラグ用の変数
var debug bool

// ファイル読み込みのタイミングでフラグを定義する
func init() {
	// コマンド共通のフラグを定義
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug enable flag")
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
