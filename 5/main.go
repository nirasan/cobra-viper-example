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
			fmt.Printf("config: %#v\n", config)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.toml", "config file name")

	// 設定ファイルの ApplicationName 項目をフラグで上書きする
	rootCmd.PersistentFlags().String("name", "", "application name")
	viper.BindPFlag("ApplicationName", rootCmd.PersistentFlags().Lookup("name"))

	cobra.OnInitialize(func() {
		viper.SetConfigFile(configFile)
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
