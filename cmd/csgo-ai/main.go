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
	steamdir := os.Getenv("STEAM_PATH")
	steamapps := os.Getenv("STEAMAPPS_PATH")
	port := os.Getenv("TELNET_PORT")
	width := os.Getenv("WIDTH")
	height := os.Getenv("HEIGHT")
	homedir, err := os.UserHomeDir()
	osDefaults := map[string][3]string{
		"windows": {`C:\Program Files (x86)\Steam\Steam.exe`, `C:\Program Files (x86)\Steam\steamapps\common`, "csgo.exe"},
		"darwin":  {`/Applications/Steam.app/Contents/MacOS/steam_osx`, homedir + `/Library/Application Support/Steam/steamapps/common`, "csgo_osx64"},
		"linux":   {`steam`, homedir + `/.steam/steam/SteamApps/common/`, "csgo"},
	}
	os.Setenv("PROCESS_NAME", osDefaults[goos][2])
	if goos == "windows" || goos == "darwin" || goos == "linux" {
		if width == "" {
			width = "1280"
		}
		if height == "" {
			height = "720"
		}
		if steamdir == "" {
			os.Setenv("STEAM_PATH", osDefaults[goos][0])
		}
		if steamapps == "" {
			os.Setenv("STEAMAPPS_PATH", osDefaults[goos][1])
		}
		startup.Run(gameid, port, width, height)
	} else {
		log.Fatal("Unsupported Operating System")
	}
}
