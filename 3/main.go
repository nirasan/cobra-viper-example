package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// 設定型
type Config struct {
	ApplicationName string
	Debug           bool
}

// 設定ファイル名
var configFile string

// 設定構造体
var config Config

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(c *cobra.Command, args []string) {
			// 受け取った設定ファイル名と設定の内容を表示
			fmt.Printf("configFile: %s\nconfig: %#v", configFile, config)
		},
	}

	// 設定ファイル名をフラグで受け取る
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.toml", "config file name")

	// cobra.Command 実行前の初期化処理を定義する。
	// rootCmd.Execute > コマンドライン引数の処理 > cobra.OnInitialize > rootCmd.Run という順に実行されるので、
	// フラグでうけとった設定ファイル名を使って設定ファイルを読み込み、コマンド実行時に設定ファイルの内容を利用することができる。
	cobra.OnInitialize(func() {

		// 設定ファイル名を viper に定義する
		viper.SetConfigFile(configFile)

		// 設定ファイルを読み込む
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 設定ファイルの内容を構造体にコピーする
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	// コマンド実行
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
