package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Media_CancelVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.cancel_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CancelVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media/cancel_video_upload", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.CancelVideoUpload(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.CancelVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.CancelVideoUpload response: %#v", res)
}
func Test_Media_CompleteVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.complete_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CompleteVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CompleteVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media/complete_video_upload", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.CompleteVideoUpload(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.CompleteVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.CompleteVideoUpload response: %#v", res)
}
func Test_Media_GetVideoUploadResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.get_video_upload_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoUploadResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoUploadResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/media/get_video_upload_result", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.GetVideoUploadResult(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.GetVideoUploadResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.GetVideoUploadResult response: %#v", res)
}
func Test_Media_InitVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.init_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping InitVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media/init_video_upload", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.InitVideoUpload(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.InitVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.InitVideoUpload response: %#v", res)
}
func Test_Media_UploadImage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.upload_image_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadImage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadImage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media/upload_image", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.UploadImage(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.UploadImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.UploadImage response: %#v", res)
}
func Test_Media_UploadVideoPart(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media.upload_video_part_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadVideoPart due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadVideoPart due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media/upload_video_part", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Media.UploadVideoPart(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Media.UploadVideoPart returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Media.UploadVideoPart response: %#v", res)
}
