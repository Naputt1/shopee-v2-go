package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Auth_GetAuthURL(t *testing.T) {
	setup()
	defer teardown()

	authURL, err := client.Auth.GetAuthURL()
	if err != nil {
		t.Fatalf("Auth.GetAuthURL error: %s", err)
	}
	t.Logf("auth url: %s", authURL)
}

func Test_Auth_GetCancelAuthURL(t *testing.T) {
	setup()
	defer teardown()

	authURL, err := client.Auth.GetCancelAuthURL()
	if err != nil {
		t.Fatalf("Auth.GetCancelAuthURL error: %s", err)
	}
	t.Logf("cancel auth url: %s", authURL)
}

func Test_Auth_GetAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/auth/token/get", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("v2.public.get_access_token_resp.json")))

	res, err := client.Auth.GetAccessToken(context.Background(), 123456, 0, "testcode")
	if err != nil {
		t.Errorf("Auth.GetToken error: %s", err)
	}

	t.Logf("return tok: %#v", res)

	if res.AccessToken != "6b5a46716e474d6f6e59777659459849" {
		t.Errorf("Token.AccessToken returned %+v, expected %+v", res.AccessToken, "6b5a46716e474d6f6e59777659459849")
	}
}

func Test_Auth_RefreshAccessToken(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/auth/access_token/get", app.APIURL),
		httpmock.NewBytesResponder(200, loadFixture("v2.public.refresh_access_token_resp.json")))

	res, err := client.Auth.RefreshAccessToken(context.Background(), 123456, 0, "testcode")
	if err != nil {
		t.Errorf("Auth.GetToken error: %s", err)
	}

	t.Logf("return tok: %#v", res)

	if res.AccessToken != "71594a4c54537649697341796363674a" {
		t.Errorf("Token.AccessToken returned %+v, expected %+v", res.AccessToken, "71594a4c54537649697341796363674a")
	}
}
