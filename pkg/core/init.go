package core

import (
	"fmt"
)

func StartServer(config Config) {
	fmt.Println(config)
	fmt.Println("Twitch Client ID: ", config.Auth.ClientID)
}
