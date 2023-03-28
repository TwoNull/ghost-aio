package main

import (
	"log"
	"os"
	"runtime"

	"github.com/0xdarktwo/ghost-aio/internal/startup"

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
	width := os.Getenv("WIDTH")
	height := os.Getenv("HEIGHT")
	if width == "" {
		width = "1280"
	}
	if height == "" {
		height = "720"
	}
	osDefaults := map[string][2]string{
		"windows": {`C:\Program Files (x86)\Steam\Steam.exe`, "csgo.exe"},
		"darwin":  {`/Applications/Steam.app/Contents/MacOS/steam_osx`, "csgo_osx64"},
		"linux":   {`steam`, "csgo"},
	}
	if goos == "windows" || goos == "darwin" || goos == "linux" {
		if steamdir == "" {
			startup.Run(gameid, osDefaults[goos][0], osDefaults[goos][1], port, width, height)
		} else {
			startup.Run(gameid, steamdir, osDefaults[goos][1], port, width, height)
		}
	} else {
		log.Fatal("Unsupported Operating System")
	}
}
