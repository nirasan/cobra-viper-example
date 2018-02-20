package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

// 設定型
type Config struct {
	B   bool
	S   string
	I   int
	I8  int8
	I16 int16
	I32 int32
	I64 int64
	U   uint
	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64
	F32 float32
	F64 float64
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

	flags := rootCmd.PersistentFlags()

	// 設定ファイル
	flags.StringVarP(&configFile, "config", "c", "config.toml", "config file name")

	// 設定ファイルの各項目ごとにフラグを定義する
	rt := reflect.TypeOf(Config{})
	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)
		name := sf.Name
		desc := fmt.Sprintf("overwrite %s value from config", name)
		switch sf.Type.Kind() {
		case reflect.Bool:
			flags.Bool(name, false, desc)
		case reflect.String:
			flags.String(name, "", desc)
		case reflect.Int:
			flags.Int(name, 0, desc)
		case reflect.Int8:
			flags.Int8(name, 0, desc)
		case reflect.Int16:
			flags.Int16(name, 0, desc)
		case reflect.Int32:
			flags.Int32(name, 0, desc)
		case reflect.Int64:
			flags.Int64(name, 0, desc)
		case reflect.Uint:
			flags.Uint(name, 0, desc)
		case reflect.Uint8:
			flags.Uint8(name, 0, desc)
		case reflect.Uint16:
			flags.Uint16(name, 0, desc)
		case reflect.Uint32:
			flags.Uint32(name, 0, desc)
		case reflect.Uint64:
			flags.Uint64(name, 0, desc)
		case reflect.Float32:
			flags.Float32(name, 0, desc)
		case reflect.Float64:
			flags.Float64(name, 0, desc)
		}
	}

	//flags.Uint64("U8", 0, "")

	// 設定ファイルの各項目をフラグで上書きする
	viper.BindPFlags(rootCmd.PersistentFlags())

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
