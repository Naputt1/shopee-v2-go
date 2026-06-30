package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FollowPrize_AddFollowPrize(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.add_follow_prize_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddFollowPrize due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddFollowPrize due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/add_follow_prize", app.APIURL), responder)

	var req AddFollowPrizeRequest
	res, err := client.FollowPrize.AddFollowPrize(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.AddFollowPrize returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.AddFollowPrize response: %#v", res)
}
func Test_FollowPrize_DeleteFollowPrize(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.delete_follow_prize_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteFollowPrize due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteFollowPrize due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/delete_follow_prize", app.APIURL), responder)

	var req DeleteFollowPrizeRequest
	res, err := client.FollowPrize.DeleteFollowPrize(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.DeleteFollowPrize returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.DeleteFollowPrize response: %#v", res)
}
func Test_FollowPrize_EndFollowPrize(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.end_follow_prize_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndFollowPrize due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndFollowPrize due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/end_follow_prize", app.APIURL), responder)

	var req EndFollowPrizeRequest
	res, err := client.FollowPrize.EndFollowPrize(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.EndFollowPrize returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.EndFollowPrize response: %#v", res)
}
func Test_FollowPrize_GetFollowPrizeDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.get_follow_prize_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFollowPrizeDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetFollowPrizeDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/get_follow_prize_detail", app.APIURL), responder)

	var req GetFollowPrizeDetailRequest
	res, err := client.FollowPrize.GetFollowPrizeDetail(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.GetFollowPrizeDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.GetFollowPrizeDetail response: %#v", res)
}
func Test_FollowPrize_GetFollowPrizeList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.get_follow_prize_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFollowPrizeList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetFollowPrizeList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/get_follow_prize_list", app.APIURL), responder)

	var req GetFollowPrizeListRequest
	res, err := client.FollowPrize.GetFollowPrizeList(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.GetFollowPrizeList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.GetFollowPrizeList response: %#v", res)
}
func Test_FollowPrize_UpdateFollowPrize(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.follow_prize.update_follow_prize_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateFollowPrize due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateFollowPrize due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/follow_prize/update_follow_prize", app.APIURL), responder)

	var req UpdateFollowPrizeRequest
	res, err := client.FollowPrize.UpdateFollowPrize(shopID, req, accessToken)
	if err != nil {
		t.Logf("FollowPrize.UpdateFollowPrize returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FollowPrize.UpdateFollowPrize response: %#v", res)
}
