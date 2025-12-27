package cloudinary

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

var CloudinaryInstance *cloudinary.Cloudinary

func InitCloudinary() {
	fmt.Println("=== INICIANDO CLOUDINARY ===")

	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar .env: %v", err)
	}

	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	fmt.Printf("CLOUDINARY_NAME: '%s'\n", cloudName)
	fmt.Printf("API_KEY: '%s'\n", apiKey)
	fmt.Printf("API_SECRET (primeros 5 chars): '%s...'\n", apiSecret[:5])

	if cloudName == "" {
		log.Fatal("ERROR: CLOUDINARY_NAME vacío")
	}
	if apiKey == "" {
		log.Fatal("ERROR: API_KEY vacío")
	}
	if apiSecret == "" {
		log.Fatal("ERROR: API_SECRET vacío")
	}

	fmt.Println("Creando instancia de Cloudinary...")
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("ERROR al crear instancia: %v", err)
	}

	CloudinaryInstance = cld
	fmt.Println("=== CLOUDINARY CONFIGURADO CORRECTAMENTE ===")
}
