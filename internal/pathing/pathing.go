package pathing

import (
	"github.com/0xdarktwo/ghost-aio/internal/telnet"
	"github.com/0xdarktwo/ghost-aio/internal/window"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/andygrunwald/vdf"
	"github.com/mrazza/gonav"
	_ "github.com/robroyd/dds"
)

var dds image.Image
var ddsw int
var ddsh int
var mesh gonav.NavMesh
var rscale float64
var tlx int
var tly int

func InitPathing(mapName string) {
	steamapps := os.Getenv("STEAMAPPS_PATH")
	mesh = loadMesh(steamapps, mapName)
	dds, ddsw, ddsh, tlx, tly, rscale = loadRadar(steamapps, mapName)
}

func GetInGameRadar(conn *net.Conn) {
	telnet.Write(conn, "name")
	time.Sleep(100 * time.Millisecond)
	radar := window.GetWindow(8, 45, 243, 243)
	resizedImg := resize.Resize(uint(ddsw), uint(ddsh), radar, resize.Bicubic)
	f, err := os.Create("CSGORadar0xDarkTwo.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	jpeg.Encode(f, resizedImg, nil)
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

func loadRadar(steamapps string, mapName string) (image.Image, int, int, int, int, float64) {
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
	bounds := ddsImage.Bounds()
	txt, err := os.Open(txtPath)
	if err != nil {
		log.Fatal("Error Opening Radar vdf File for Current Map")
	}
	p := vdf.NewParser(txt)
	txtMap, err := p.Parse()
	if err != nil {
		log.Fatal("Error Parsing Radar vdf File for Current Map")
	}
	posX, _ = strconv.Atoi(txtMap[mapName].(map[string]interface{})["pos_x"].(string))
	posY, _ = strconv.Atoi(txtMap[mapName].(map[string]interface{})["pos_y"].(string))
	scale, _ = strconv.ParseFloat(txtMap[mapName].(map[string]interface{})["scale"].(string), 64)
	if &posX == nil || &posY == nil || &scale == nil {
		log.Fatal("Error Collecting Radar vdf Values for Current Map")
	}
	return ddsImage, bounds.Max.X, bounds.Max.Y, posX, posY, scale
}
