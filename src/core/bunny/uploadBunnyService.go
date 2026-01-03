package bunny

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BunnyUploadResponse struct {
	VideoID   string `json:"guid"`
	LibraryID int    `json:"videoLibraryId"`
	Title     string `json:"title"`
	Status    int    `json:"status"`
}

// UploadVideo - Sube un video a Bunny.net
func UploadVideo(fileBytes []byte, fileName string, title string) (string, error) {
	// 1. Crear el video en Bunny.net
	videoID, err := createVideoInBunny(title)
	if err != nil {
		return "", fmt.Errorf("error al crear video en Bunny: %v", err)
	}

	// 2. Subir el archivo del video
	err = uploadVideoFile(videoID, fileBytes)
	if err != nil {
		// Si falla la subida, intentar eliminar el video creado
		_ = DeleteVideo(videoID)
		return "", fmt.Errorf("error al subir archivo de video: %v", err)
	}

	return videoID, nil
}

// createVideoInBunny - Crea un video vacío en Bunny.net
func createVideoInBunny(title string) (string, error) {
	url := fmt.Sprintf("https://video.bunnycdn.com/library/%s/videos", Config.LibraryID)

	payload := map[string]interface{}{
		"title": title,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// IMPORTANTE: Bunny.net usa "AccessKey" en el header
	req.Header.Set("AccessKey", Config.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("error al crear video: status %d, body: %s", resp.StatusCode, string(body))
	}

	var result BunnyUploadResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error al decodificar respuesta: %v, body: %s", err, string(body))
	}

	return result.VideoID, nil
}

// uploadVideoFile - Sube el archivo del video a Bunny.net
func uploadVideoFile(videoID string, fileBytes []byte) error {
	url := fmt.Sprintf("https://video.bunnycdn.com/library/%s/videos/%s", Config.LibraryID, videoID)

	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileBytes))
	if err != nil {
		return err
	}

	// IMPORTANTE: Bunny.net usa "AccessKey" en el header
	req.Header.Set("AccessKey", Config.APIKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{Timeout: 300 * time.Second} // 5 minutos timeout
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("error al subir video: status %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// DeleteVideo - Elimina un video de Bunny.net
func DeleteVideo(videoID string) error {
	url := fmt.Sprintf("https://video.bunnycdn.com/library/%s/videos/%s", Config.LibraryID, videoID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("AccessKey", Config.APIKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error al eliminar video: status %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetVideoEmbedURL - Genera la URL de embed del video
func GetVideoEmbedURL(videoID string) string {
	return fmt.Sprintf("%s/embed/%s/%s", Config.PullZoneURL, Config.LibraryID, videoID)
}
