package pathing

import (
	"image"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andygrunwald/vdf"
	_ "github.com/lukegb/dds"
	"github.com/mrazza/gonav"
)

var dds image.Image
var mesh gonav.NavMesh
var scale float64
var posX int
var posY int

func InitPathing(mapName string) {
	steamapps := os.Getenv("STEAMAPPS_PATH")
	mesh = loadMesh(steamapps, mapName)
	dds, posX, posY, scale = loadRadar(steamapps, mapName)
	log.Println(posX)
	log.Println(posY)
	log.Println(scale)
}

func loadMesh(steamapps string, mapName string) gonav.NavMesh {
	mapPath := filepath.Join(filepath.Dir(steamapps), "common", "Counter-Strike Global Offensive", "csgo", "maps", mapName+".nav")
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

func loadRadar(steamapps string, mapName string) (image.Image, int, int, float64) {
	var posX int
	var posY int
	var scale float64
	ddsPath := filepath.Join(filepath.Dir(steamapps), "common", "Counter-Strike Global Offensive", "csgo", "resource", "overviews", mapName+"_radar.dds")
	txtPath := filepath.Join(filepath.Dir(steamapps), "common", "Counter-Strike Global Offensive", "csgo", "resource", "overviews", mapName+".txt")
	dds, err := os.Open(ddsPath)
	if err != nil {
		log.Fatal("Error Opening Radar dds File for Current Map")
	}
	ddsImage, _, err := image.Decode(dds)
	if err != nil {
		log.Fatal("Error Decoding Radar dds File for Current Map")
	}
	txt, err := os.Open(txtPath)
	if err != nil {
		log.Fatal("Error Opening Radar vdf File for Current Map")
	}
	p := vdf.NewParser(txt)
	txtMap, err := p.Parse()
	if err != nil {
		log.Fatal("Error Parsing Radar vdf File for Current Map")
	}
	posX, _ = strconv.Atoi(txtMap[mapName].(map[string]string)["pos_x"])
	posY, _ = strconv.Atoi(txtMap[mapName].(map[string]string)["pos_y"])
	scale, _ = strconv.ParseFloat(txtMap[mapName].(map[string]string)["scale"], 64)
	if &posX == nil || &posY == nil || &scale == nil {
		log.Fatal("Error Collecting Radar vdf Values for Current Map")
	}
	return ddsImage, posX, posY, scale
}
