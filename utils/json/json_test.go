package json

import (
	"src/models"
	"testing"
)

func TestJson(t *testing.T) {
	err := models.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	_ = parseUsers()
	_ = parseSeller()
	_ = parseCommodity()
	_ = parsePlatform()
	_ = parseAdmin()

	_ = parseItem()
	_ = parseFavorite()
	_ = parsePriceChange()
	_ = parseMessage()

}
