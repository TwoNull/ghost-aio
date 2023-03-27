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
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file. Please rename the '.env.example' template to '.env'")
	}
	gameid := os.Getenv("GAME_ID")
	port := os.Getenv("TELNET_PORT")
	steamdir := os.Getenv("STEAM_DIRECTORY")
	osDefaults := map[string][2]string{
		"windows": {`C:\Program Files (x86)\Steam\Steam.exe`, "csgo.exe"},
		"darwin":  {`/Applications/Steam.app/Contents/MacOS/steam_osx`, "csgo_osx64"},
		"linux":   {`steam`, "csgo"},
	}
	if goos == "windows" || goos == "darwin" || goos == "linux" {
		if steamdir == "" {
			startup.Run(gameid, osDefaults[goos][0], osDefaults[goos][1], port)
		} else {
			startup.Run(gameid, steamdir, osDefaults[goos][1], port)
		}
	} else {
		log.Fatal("Unsupported Operating System")
	}
}
