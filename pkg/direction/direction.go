package direction

import (
	"os"

	"googlemaps.github.io/maps"
)

func getClient() (*maps.Client, error) {
	apiKey := os.Getenv("API_KEY")

	return maps.NewClient(maps.WithAPIKey(apiKey))
}