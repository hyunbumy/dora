package direction

import (
	"testing"
)

func TestGetClient(t *testing.T) {
	client, err := getClient()
	t.Log(*client, err)
}