package goshopee

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAutoTokenRefresh(t *testing.T) {
	setup()
	defer teardown()

	rawErr := `{"error":"invalid_acceess_token","message":"Invalid access token","request_id":"req123"}`

	refreshRes := RefreshAccessTokenResponse{
		AccessToken:  "new_access_token",
		RefreshToken: "new_refresh_token",
		ExpireIn:     3600,
	}
	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/auth/access_token/get", app.APIURL),
		httpmock.NewJsonResponderOrPanic(200, refreshRes))

	successRes := GetShopInfoResponse{
		ShopName: "Test Shop",
	}

	callCount := 0
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/shop/get_shop_info", app.APIURL),
		func(req *http.Request) (*http.Response, error) {
			callCount++
			if callCount == 1 {
				return httpmock.NewStringResponse(200, rawErr), nil
			}
			if req.URL.Query().Get("access_token") != "new_access_token" {
				return httpmock.NewStringResponse(400, `{"error":"wrong_token"}`), nil
			}
			return httpmock.NewJsonResponse(200, successRes)
		})

	refreshed := false
	client.RefreshToken = "old_refresh_token"
	client.OnTokenRefresh = func(res *RefreshAccessTokenResponse, meta interface{}) {
		refreshed = true
		if res.AccessToken != "new_access_token" {
			t.Errorf("Expected new access token, got %s", res.AccessToken)
		}
	}

	res, err := client.Shop.GetShopInfo(context.Background(), shopID, 0, accessToken)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.ShopName != "Test Shop" {
		t.Errorf("Expected shop name Test Shop, got %s", res.ShopName)
	}

	if !refreshed {
		t.Error("Expected OnTokenRefresh to be called")
	}

	if callCount != 2 {
		t.Errorf("Expected 2 calls to GetShopInfo, got %d", callCount)
	}
}
