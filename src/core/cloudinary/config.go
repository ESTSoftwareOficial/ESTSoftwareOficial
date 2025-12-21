package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

var CloudinaryInstance *cloudinary.Cloudinary

func InitCloudinary() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar .env para Cloudinary: %v", err)
	}

	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Error: CLOUDINARY_NAME, API_KEY o API_SECRET no están configurados")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatalf("Error al inicializar Cloudinary: %v", err)
	}

	CloudinaryInstance = cld
	log.Println("Cloudinary configurado correctamente")
}
