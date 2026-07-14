package goshopee

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/jarcoal/httpmock"
)

var (
	client        *DefaultClient
	app           App
	shopID        uint64 = 123456
	merchantID    uint64 = 789012
	accessToken          = "test_access_token"
	sid           uint64 = shopID
	mid           uint64 = merchantID
	tok                  = accessToken
	skippedMu     sync.Mutex
	skippedRoutes []string
)

func setup() {
	httpmock.Activate()
	app = App{
		PartnerID:   123456,
		PartnerKey:  "test_partner_key",
		RedirectURL: "https://example.com/callback",
		APIURL:      "https://open-api.test.com",
	}
	client = NewDefaultClient(app)

}

func teardown() {
	httpmock.DeactivateAndReset()
}

func loadFixture(filename string) []byte {
	f, err := os.ReadFile("fixtures/" + filename)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %v", filename))
	}
	return f
}

func loadFixtureSafe(path string) (interface{}, error) {
	f, err := os.Open("fixtures/" + path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
