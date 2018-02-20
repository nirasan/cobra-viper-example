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

// 設定項目を上書きする環境変数のプレフィックス
var envPrefix string

// 設定構造体
var config Config

func main() {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(c *cobra.Command, args []string) {
			// 読み込んだ設定の内容を表示
			fmt.Printf("config: %#v\n", config)
		},
	}

	// 設定ファイル名をフラグで受け取る
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.toml", "config file name")
	// 設定項目を上書きする環境変数のプレフィックスをフラグで受け取る
	rootCmd.PersistentFlags().StringVar(&envPrefix, "env-prefix", "", "env prefix, enabled if defined")

	// cobra.Command 実行前の初期化処理を定義する
	cobra.OnInitialize(func() {

		// 設定ファイル名を viper に定義する
		viper.SetConfigFile(configFile)

		// 設定項目を上書きする環境変数のプレフィックスの指定
		// 対象項目が "Debug" で指定したプレフィックスが "app" なら、環境変数に "APP_DEBUG" を設定しておくことで値を上書きできる
		viper.SetEnvPrefix(envPrefix)

		// 環境変数を自動で読み込んで項目に対応する値が存在すれば上書きする
		viper.AutomaticEnv()

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
