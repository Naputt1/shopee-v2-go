package goshopee

import (
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	code := m.Run()
	if len(skippedRoutes) > 0 {
		fmt.Println("\nSkipped routes (missing fixtures):")
		for _, r := range skippedRoutes {
			fmt.Println("  -", r)
		}
	}
	os.Exit(code)
}

func TestCheckResponseError(t *testing.T) {
	setup()
	defer teardown()

	rawErr := `{"error":"error_auth","message":"Invalid access token","request_id":"req123"}`
	httpmock.RegisterResponder("GET", "https://open-api.test.com/api/v2/test",
		httpmock.NewStringResponder(401, rawErr))

	resp, err := client.Client.Get("https://open-api.test.com/api/v2/test")
	if err != nil {
		t.Fatalf("Unexpected error from mock HTTP call: %v", err)
	}
	defer resp.Body.Close()

	apiErr := CheckResponseError(resp)
	if apiErr == nil {
		t.Fatal("Expected error, got nil")
	}

	if !IsShopeeError(apiErr, "error_auth") {
		t.Errorf("Expected IsShopeeError to be true for error_auth, got %v", apiErr)
	}

	re, ok := apiErr.(ResponseError)
	if !ok {
		t.Fatalf("Expected ResponseError, got %T", apiErr)
	}

	if re.ShopeeError != "error_auth" {
		t.Errorf("Expected ShopeeError error_auth, got %s", re.ShopeeError)
	}

	if re.RequestID != "req123" {
		t.Errorf("Expected RequestID req123, got %s", re.RequestID)
	}
}

func TestGenericMeta(t *testing.T) {
	type MyMeta struct {
		ID   int
		Name string
	}

	meta := MyMeta{ID: 123, Name: "Test"}
	app := App{
		APIURL: "https://example.com",
	}

	var capturedMeta MyMeta
	c := NewClient(app,
		WithMeta(meta),
		WithOnTokenRefresh(func(res *RefreshAccessTokenResponse, m MyMeta) {
			capturedMeta = m
		}),
	)

	if c.Meta.ID != 123 {
		t.Errorf("Expected meta ID 123, got %d", c.Meta.ID)
	}

	c.OnTokenRefresh(nil, c.Meta)
	if capturedMeta.ID != 123 {
		t.Errorf("Expected captured meta ID 123, got %d", capturedMeta.ID)
	}
}

func TestNewDefaultClient(t *testing.T) {
	app := App{
		APIURL: "https://example.com",
	}

	c := NewDefaultClient(app,
		WithRetryDefault(3),
		WithMetaDefault("some meta"),
	)

	if c.retries != 3 {
		t.Errorf("Expected retries 3, got %d", c.retries)
	}

	if c.Meta != "some meta" {
		t.Errorf("Expected meta 'some meta', got %v", c.Meta)
	}

	// Reset the packages global client to avoid affecting other tests
	client = c
}
