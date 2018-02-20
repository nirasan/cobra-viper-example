package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	name1 string
)

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(c *cobra.Command, args []string) {
			// セットされた値を変数から取得する
			fmt.Println("name1:", name1)

			// フラグ名で値を取得する
			if name2, err := c.PersistentFlags().GetString("name2"); err == nil {
				fmt.Println("name2:", name2)
			}
		},
	}

	// フラグの値を変数にセットする場合
	// 第1引数: 変数のポインタ
	// 第2引数: フラグ名
	// 第3引数: デフォルト値
	// 第4引数: 説明
	rootCmd.PersistentFlags().StringVar(&name1, "name1", "name1", "your name1")

	// フラグの値をフラグ名で参照する場合
	// 第1引数: フラグ名
	// 第2引数: 短縮フラグ名（末尾が "P" の関数では短縮フラグを指定できる）
	// 第3引数: デフォルト値
	// 第4引数: 説明
	rootCmd.PersistentFlags().StringP("name2", "n", "name2", "your name2")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
