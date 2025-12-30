package bunny

import (
	"log"
	"os"
)

type BunnyConfig struct {
	LibraryID   string
	APIKey      string
	PullZoneURL string
}

var Config BunnyConfig

func InitBunny() {
	Config = BunnyConfig{
		LibraryID:   os.Getenv("BUNNY_LIBRARY_ID"),
		APIKey:      os.Getenv("BUNNY_API_KEY"),
		PullZoneURL: os.Getenv("BUNNY_PULL_ZONE_URL"),
	}

	if Config.LibraryID == "" || Config.APIKey == "" {
		log.Fatal("BUNNY_LIBRARY_ID y BUNNY_API_KEY son requeridos")
	}

	log.Println("Bunny.net inicializado correctamente")
	log.Printf("Library ID: %s", Config.LibraryID)
}
