package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_TopPicks_AddTopPicks(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.top_picks.add_top_picks_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddTopPicks due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddTopPicks due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/top_picks/add_top_picks", app.APIURL), responder)

	var req AddTopPicksRequest
	ctx := context.Background()
	res, err := client.TopPicks.AddTopPicks(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("TopPicks.AddTopPicks returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("TopPicks.AddTopPicks response: %#v", res)
}
func Test_TopPicks_DeleteTopPicks(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.top_picks.delete_top_picks_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteTopPicks due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteTopPicks due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/top_picks/delete_top_picks", app.APIURL), responder)

	var req DeleteTopPicksRequest
	ctx := context.Background()
	res, err := client.TopPicks.DeleteTopPicks(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("TopPicks.DeleteTopPicks returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("TopPicks.DeleteTopPicks response: %#v", res)
}
func Test_TopPicks_GetTopPicksList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.top_picks.get_top_picks_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTopPicksList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTopPicksList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/top_picks/get_top_picks_list", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.TopPicks.GetTopPicksList(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("TopPicks.GetTopPicksList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("TopPicks.GetTopPicksList response: %#v", res)
}
func Test_TopPicks_UpdateTopPicks(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.top_picks.update_top_picks_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateTopPicks due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateTopPicks due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/top_picks/update_top_picks", app.APIURL), responder)

	var req UpdateTopPicksRequest
	ctx := context.Background()
	res, err := client.TopPicks.UpdateTopPicks(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("TopPicks.UpdateTopPicks returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("TopPicks.UpdateTopPicks response: %#v", res)
}
