package main

import (
	"log"
	"os"
	"runtime"
	"startup"

	"github.com/joho/godotenv"
)

func main() {
	goos := runtime.GOOS
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file. Please rename the '.env.example' template to '.env'")
	}
	gameid := os.Getenv("GAME_ID")
	port := os.Getenv("TELNET_PORT")
	steamdir := os.Getenv("STEAM_DIRECTORY")
	osDefaults := map[string]string{
		"windows": `C:\Program Files (x86)\Steam\Steam.exe`,
		"darwin":  `/Applications/Steam.app/Contents/MacOS/steam.sh`,
		"linux":   `steam`,
	}
	if goos == "windows" || goos == "darwin" || goos == "linux" {
		if steamdir == "" {
			startup.StartApp(gameid, osDefaults[goos], port)
		} else {
			startup.StartApp(gameid, steamdir, port)
		}
	} else {
		log.Fatal("Unsupported Operating System")
	}
}
