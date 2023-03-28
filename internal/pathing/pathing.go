package pathing

import (
	"github.com/mrazza/gonav"
	"log"
	"os"
	"path/filepath"
)

func InitPathing(mapName string) {
	mesh := loadMap(mapName)
	area := mesh.GetNearestArea(gonav.Vector3{-250, 1000, 20}, true)
	log.Println(area)
}

func loadMap(name string) gonav.NavMesh {
	steamapps := os.Getenv("STEAM_GAMES_PATH")
	mapPath := filepath.Join(filepath.Dir(steamapps), "common", "Counter-Strike Global Offensive", "csgo", "maps", name+".nav")
	f, err := os.Open(mapPath)
	if err != nil {
		log.Fatal("Error Opening .nav File for Current Map.")
	}
	parser := gonav.Parser{Reader: f}
	mesh, err := parser.Parse()
	if err != nil {
		log.Fatal("Error Parsing .nav File for Current Map.")
	}
	return mesh
}
