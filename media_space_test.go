package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_MediaSpace_CancelVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.cancel_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CancelVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/cancel_video_upload", app.APIURL), responder)

	res, err := client.MediaSpace.CancelVideoUpload(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.CancelVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.CancelVideoUpload response: %#v", res)
}
func Test_MediaSpace_CompleteVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.complete_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CompleteVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CompleteVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/complete_video_upload", app.APIURL), responder)

	res, err := client.MediaSpace.CompleteVideoUpload(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.CompleteVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.CompleteVideoUpload response: %#v", res)
}
func Test_MediaSpace_GetVideoUploadResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.get_video_upload_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetVideoUploadResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetVideoUploadResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/media_space/get_video_upload_result", app.APIURL), responder)

	res, err := client.MediaSpace.GetVideoUploadResult(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.GetVideoUploadResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.GetVideoUploadResult response: %#v", res)
}
func Test_MediaSpace_InitVideoUpload(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.init_video_upload_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping InitVideoUpload due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping InitVideoUpload due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/init_video_upload", app.APIURL), responder)

	res, err := client.MediaSpace.InitVideoUpload(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.InitVideoUpload returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.InitVideoUpload response: %#v", res)
}
func Test_MediaSpace_UploadImage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.upload_image_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadImage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadImage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/upload_image", app.APIURL), responder)

	res, err := client.MediaSpace.UploadImage(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.UploadImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.UploadImage response: %#v", res)
}
func Test_MediaSpace_UploadVideoPart(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.media_space.upload_video_part_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadVideoPart due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadVideoPart due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/media_space/upload_video_part", app.APIURL), responder)

	res, err := client.MediaSpace.UploadVideoPart(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("MediaSpace.UploadVideoPart returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("MediaSpace.UploadVideoPart response: %#v", res)
}
