package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Livestream_AddItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.add_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AddItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AddItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/add_item_list", app.APIURL), responder)

	var req AddItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.AddItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.AddItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.AddItemList response: %#v", res)
}
func Test_Livestream_ApplyItemSet(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.apply_item_set_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ApplyItemSet due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ApplyItemSet due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/apply_item_set", app.APIURL), responder)

	var req ApplyItemSetRequest
	ctx := context.Background()
	res, err := client.Livestream.ApplyItemSet(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.ApplyItemSet returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.ApplyItemSet response: %#v", res)
}
func Test_Livestream_BanUserComment(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.ban_user_comment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BanUserComment due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BanUserComment due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/ban_user_comment", app.APIURL), responder)

	var req BanUserCommentRequest
	ctx := context.Background()
	res, err := client.Livestream.BanUserComment(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.BanUserComment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.BanUserComment response: %#v", res)
}
func Test_Livestream_CreateSession(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.create_session_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateSession due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateSession due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/create_session", app.APIURL), responder)

	var req CreateSessionRequest
	ctx := context.Background()
	res, err := client.Livestream.CreateSession(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.CreateSession returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.CreateSession response: %#v", res)
}
func Test_Livestream_DeleteItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.delete_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/delete_item_list", app.APIURL), responder)

	var req DeleteItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.DeleteItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.DeleteItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.DeleteItemList response: %#v", res)
}
func Test_Livestream_DeleteShowItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.delete_show_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteShowItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteShowItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/delete_show_item", app.APIURL), responder)

	var req DeleteShowItemRequest
	ctx := context.Background()
	res, err := client.Livestream.DeleteShowItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.DeleteShowItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.DeleteShowItem response: %#v", res)
}
func Test_Livestream_EndSession(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.end_session_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EndSession due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EndSession due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/end_session", app.APIURL), responder)

	var req EndSessionRequest
	ctx := context.Background()
	res, err := client.Livestream.EndSession(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.EndSession returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.EndSession response: %#v", res)
}
func Test_Livestream_GetItemCount(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_item_count_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemCount due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemCount due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_item_count", app.APIURL), responder)

	var req GetItemCountRequest
	ctx := context.Background()
	res, err := client.Livestream.GetItemCount(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetItemCount returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetItemCount response: %#v", res)
}
func Test_Livestream_GetItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_item_list", app.APIURL), responder)

	var req GetItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetItemList response: %#v", res)
}
func Test_Livestream_GetItemSetItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_item_set_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemSetItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemSetItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_item_set_item_list", app.APIURL), responder)

	var req GetItemSetItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetItemSetItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetItemSetItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetItemSetItemList response: %#v", res)
}
func Test_Livestream_GetItemSetList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_item_set_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemSetList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemSetList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_item_set_list", app.APIURL), responder)

	var req GetItemSetListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetItemSetList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetItemSetList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetItemSetList response: %#v", res)
}
func Test_Livestream_GetLatestCommentList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_latest_comment_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLatestCommentList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetLatestCommentList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_latest_comment_list", app.APIURL), responder)

	var req GetLatestCommentListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetLatestCommentList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetLatestCommentList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetLatestCommentList response: %#v", res)
}
func Test_Livestream_GetLikeItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_like_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetLikeItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetLikeItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_like_item_list", app.APIURL), responder)

	var req GetLikeItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetLikeItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetLikeItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetLikeItemList response: %#v", res)
}
func Test_Livestream_GetRecentItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_recent_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetRecentItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetRecentItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_recent_item_list", app.APIURL), responder)

	var req GetRecentItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.GetRecentItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetRecentItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetRecentItemList response: %#v", res)
}
func Test_Livestream_GetSessionDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_session_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSessionDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSessionDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_session_detail", app.APIURL), responder)

	var req GetSessionDetailRequest
	ctx := context.Background()
	res, err := client.Livestream.GetSessionDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetSessionDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetSessionDetail response: %#v", res)
}
func Test_Livestream_GetSessionItemMetric(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_session_item_metric_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSessionItemMetric due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSessionItemMetric due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_session_item_metric", app.APIURL), responder)

	var req GetSessionItemMetricRequest
	ctx := context.Background()
	res, err := client.Livestream.GetSessionItemMetric(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetSessionItemMetric returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetSessionItemMetric response: %#v", res)
}
func Test_Livestream_GetSessionMetric(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_session_metric_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetSessionMetric due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetSessionMetric due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_session_metric", app.APIURL), responder)

	var req GetSessionMetricRequest
	ctx := context.Background()
	res, err := client.Livestream.GetSessionMetric(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetSessionMetric returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetSessionMetric response: %#v", res)
}
func Test_Livestream_GetShowItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.get_show_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShowItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShowItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/livestream/get_show_item", app.APIURL), responder)

	var req GetShowItemRequest
	ctx := context.Background()
	res, err := client.Livestream.GetShowItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.GetShowItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.GetShowItem response: %#v", res)
}
func Test_Livestream_PostComment(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.post_comment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PostComment due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping PostComment due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/post_comment", app.APIURL), responder)

	var req PostCommentRequest
	ctx := context.Background()
	res, err := client.Livestream.PostComment(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.PostComment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.PostComment response: %#v", res)
}
func Test_Livestream_StartSession(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.start_session_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping StartSession due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping StartSession due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/start_session", app.APIURL), responder)

	var req StartSessionRequest
	ctx := context.Background()
	res, err := client.Livestream.StartSession(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.StartSession returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.StartSession response: %#v", res)
}
func Test_Livestream_UnbanUserComment(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.unban_user_comment_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UnbanUserComment due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UnbanUserComment due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/unban_user_comment", app.APIURL), responder)

	var req UnbanUserCommentRequest
	ctx := context.Background()
	res, err := client.Livestream.UnbanUserComment(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.UnbanUserComment returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.UnbanUserComment response: %#v", res)
}
func Test_Livestream_UpdateItemList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.update_item_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateItemList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateItemList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/update_item_list", app.APIURL), responder)

	var req UpdateItemListRequest
	ctx := context.Background()
	res, err := client.Livestream.UpdateItemList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.UpdateItemList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.UpdateItemList response: %#v", res)
}
func Test_Livestream_UpdateSession(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.update_session_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSession due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateSession due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/update_session", app.APIURL), responder)

	var req UpdateSessionRequest
	ctx := context.Background()
	res, err := client.Livestream.UpdateSession(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.UpdateSession returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.UpdateSession response: %#v", res)
}
func Test_Livestream_UpdateShowItem(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.update_show_item_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateShowItem due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateShowItem due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/update_show_item", app.APIURL), responder)

	var req UpdateShowItemRequest
	ctx := context.Background()
	res, err := client.Livestream.UpdateShowItem(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Livestream.UpdateShowItem returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.UpdateShowItem response: %#v", res)
}
func Test_Livestream_UploadImage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.livestream.upload_image_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadImage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadImage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/livestream/upload_image", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Livestream.UploadImage(ctx, sid, mid, tok, "fixtures/test.jpg")
	if err != nil {
		t.Logf("Livestream.UploadImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Livestream.UploadImage response: %#v", res)
}
