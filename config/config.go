package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	// Token is the bots discord token.
	Token string
	// BotPrefix is what all commands from the user should start with
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

// ReadConfig reads the configuration file for the bot. This should only be read once by the init func
func ReadConfig() error {
	fmt.Println("Reading from config file...")

	// Read from the config file
	file, err := ioutil.ReadFile(".\\config\\config.json")
	if err != nil {
		fmt.Println("Failed to read from config file:", err.Error())
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("Failed to unmarshal payload:", err.Error())
		return err
	}
	fmt.Printf("Token: %v\nBot Prefix: %v", config.Token, config.BotPrefix)

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
