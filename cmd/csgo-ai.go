package main

import (
	"log"
	"os"
	"runtime"
	"startup"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file. Please rename the '.env.example' template to '.env'")
	}
	gameid := os.Getenv("GAME_ID")
	port := os.Getenv("TELNET_PORT")
	steamdir := os.Getenv("STEAM_DIRECTORY")
	if steamdir == "" {
		if runtime.GOOS == "windows" {
			steamdir = `C:\Program Files (x86)\Steam\Steam.exe`
		}
		if runtime.GOOS == "linux" {
			steamdir = `steam`
		}
		if runtime.GOOS == "darwin" {
			steamdir = `/Applications/Steam.app/Contents/MacOS/steam.sh`
		}
	}
	startup.StartApp(gameid, steamdir, port)
}
