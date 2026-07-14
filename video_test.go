package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Video_DeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.delete_video_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteVideo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteVideo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/video/delete_video", app.APIURL), responder)

	var req DeleteVideoRequest
	ctx := context.Background()
	res, err := client.Video.DeleteVideo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.DeleteVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.DeleteVideo response: %#v", res)
}
func Test_Video_EditVideoInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.edit_video_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping EditVideoInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping EditVideoInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/video/edit_video_info", app.APIURL), responder)

	var req EditVideoInfoRequest
	ctx := context.Background()
	res, err := client.Video.EditVideoInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.EditVideoInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.EditVideoInfo response: %#v", res)
}
func Test_Video_GetCoverList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_cover_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCoverList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCoverList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_cover_list", app.APIURL), responder)

	var req GetCoverListRequest
	ctx := context.Background()
	res, err := client.Video.GetCoverList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetCoverList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetCoverList response: %#v", res)
}
func Test_Video_GetMetricTrend(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_metric_trend_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMetricTrend due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMetricTrend due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_metric_trend", app.APIURL), responder)

	var req GetMetricTrendRequest
	ctx := context.Background()
	res, err := client.Video.GetMetricTrend(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetMetricTrend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetMetricTrend response: %#v", res)
}
func Test_Video_GetOverviewPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_overview_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOverviewPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOverviewPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_overview_performance", app.APIURL), responder)

	var req GetOverviewPerformanceRequest
	ctx := context.Background()
	res, err := client.Video.GetOverviewPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetOverviewPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetOverviewPerformance response: %#v", res)
}
func Test_Video_GetProdcutPerformanceList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_prodcut_performance_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetProdcutPerformanceList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetProdcutPerformanceList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_prodcut_performance_list", app.APIURL), responder)

	var req GetProdcutPerformanceListRequest
	ctx := context.Background()
	res, err := client.Video.GetProdcutPerformanceList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetProdcutPerformanceList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetProdcutPerformanceList response: %#v", res)
}
func Test_Video_GetUserDemographics(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_user_demographics_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetUserDemographics due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetUserDemographics due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_user_demographics", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Video.GetUserDemographics(ctx, sid, mid, tok)
	if err != nil {
		t.Logf("Video.GetUserDemographics returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetUserDemographics response: %#v", res)
}
func Test_Video_GetVideoDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_detail", app.APIURL), responder)

	var req GetVideoDetailRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoDetail response: %#v", res)
}
func Test_Video_GetVideoDetailAudienceDistribution(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_detail_audience_distribution_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailAudienceDistribution due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailAudienceDistribution due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_detail_audience_distribution", app.APIURL), responder)

	var req GetVideoDetailAudienceDistributionRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoDetailAudienceDistribution(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoDetailAudienceDistribution returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoDetailAudienceDistribution response: %#v", res)
}
func Test_Video_GetVideoDetailMetricTrend(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_detail_metric_trend_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailMetricTrend due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailMetricTrend due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_detail_metric_trend", app.APIURL), responder)

	var req GetVideoDetailMetricTrendRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoDetailMetricTrend(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoDetailMetricTrend returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoDetailMetricTrend response: %#v", res)
}
func Test_Video_GetVideoDetailPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_detail_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_detail_performance", app.APIURL), responder)

	var req GetVideoDetailPerformanceRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoDetailPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoDetailPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoDetailPerformance response: %#v", res)
}
func Test_Video_GetVideoDetailProductPerformance(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_detail_product_performance_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailProductPerformance due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoDetailProductPerformance due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_detail_product_performance", app.APIURL), responder)

	var req GetVideoDetailProductPerformanceRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoDetailProductPerformance(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoDetailProductPerformance returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoDetailProductPerformance response: %#v", res)
}
func Test_Video_GetVideoList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_list", app.APIURL), responder)

	var req GetVideoListRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoList response: %#v", res)
}
func Test_Video_GetVideoPerformanceList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.get_video_performance_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoPerformanceList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoPerformanceList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/video/get_video_performance_list", app.APIURL), responder)

	var req GetVideoPerformanceListRequest
	ctx := context.Background()
	res, err := client.Video.GetVideoPerformanceList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.GetVideoPerformanceList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.GetVideoPerformanceList response: %#v", res)
}
func Test_Video_PostVideo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.video.post_video_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping PostVideo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping PostVideo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/video/post_video", app.APIURL), responder)

	var req PostVideoRequest
	ctx := context.Background()
	res, err := client.Video.PostVideo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Video.PostVideo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Video.PostVideo response: %#v", res)
}
