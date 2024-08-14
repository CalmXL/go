package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func getEnv(env string) any {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {
	env := getEnv("SALE_SYSTEM_ENV")
	fmt.Println(env)
}
