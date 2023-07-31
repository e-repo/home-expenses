package env

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var envFiles = []string{
	".env",
	"local.env",
}

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	for _, filename := range envFiles {
		if isFileExists(filename) {
			viper.SetConfigFile(filename)
		}
	}

	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("Error reading app.env.local config:", err)
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config: ", err)
		os.Exit(1)
	}
}

func isFileExists(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}
