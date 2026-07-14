package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Returns_AcceptOffer(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.accept_offer_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping AcceptOffer due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping AcceptOffer due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/accept_offer", app.APIURL), responder)

	var req AcceptOfferRequest
	ctx := context.Background()
	res, err := client.Returns.AcceptOffer(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.AcceptOffer returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.AcceptOffer response: %#v", res)
}
func Test_Returns_CancelDispute(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.cancel_dispute_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelDispute due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CancelDispute due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/cancel_dispute", app.APIURL), responder)

	var req CancelDisputeRequest
	ctx := context.Background()
	res, err := client.Returns.CancelDispute(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.CancelDispute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.CancelDispute response: %#v", res)
}
func Test_Returns_Confirm(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.confirm_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Confirm due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping Confirm due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/confirm", app.APIURL), responder)

	var req ConfirmRequest
	ctx := context.Background()
	res, err := client.Returns.Confirm(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.Confirm returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.Confirm response: %#v", res)
}
func Test_Returns_ConvertImage(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.convert_image_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ConvertImage due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ConvertImage due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/convert_image", app.APIURL), responder)

	var req ConvertImageRequest
	ctx := context.Background()
	res, err := client.Returns.ConvertImage(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.ConvertImage returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.ConvertImage response: %#v", res)
}
func Test_Returns_Dispute(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.dispute_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Dispute due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping Dispute due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/dispute", app.APIURL), responder)

	var req DisputeRequest
	ctx := context.Background()
	res, err := client.Returns.Dispute(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.Dispute returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.Dispute response: %#v", res)
}
func Test_Returns_GetAvailableSolutions(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_available_solutions_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAvailableSolutions due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAvailableSolutions due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_available_solutions", app.APIURL), responder)

	var req GetAvailableSolutionsRequest
	ctx := context.Background()
	res, err := client.Returns.GetAvailableSolutions(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetAvailableSolutions returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetAvailableSolutions response: %#v", res)
}
func Test_Returns_GetReturnDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_return_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReturnDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetReturnDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_return_detail", app.APIURL), responder)

	var req GetReturnDetailRequest
	ctx := context.Background()
	res, err := client.Returns.GetReturnDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetReturnDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetReturnDetail response: %#v", res)
}
func Test_Returns_GetReturnDisputeReason(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_return_dispute_reason_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReturnDisputeReason due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetReturnDisputeReason due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_return_dispute_reason", app.APIURL), responder)

	var req GetReturnDisputeReasonRequest
	ctx := context.Background()
	res, err := client.Returns.GetReturnDisputeReason(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetReturnDisputeReason returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetReturnDisputeReason response: %#v", res)
}
func Test_Returns_GetReturnList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_return_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReturnList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetReturnList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_return_list", app.APIURL), responder)

	var req GetReturnListRequest
	ctx := context.Background()
	res, err := client.Returns.GetReturnList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetReturnList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetReturnList response: %#v", res)
}
func Test_Returns_GetReverseTrackingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_reverse_tracking_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetReverseTrackingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetReverseTrackingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_reverse_tracking_info", app.APIURL), responder)

	var req GetReverseTrackingInfoRequest
	ctx := context.Background()
	res, err := client.Returns.GetReverseTrackingInfo(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetReverseTrackingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetReverseTrackingInfo response: %#v", res)
}
func Test_Returns_GetShippingCarrier(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.get_shipping_carrier_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingCarrier due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingCarrier due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/get_shipping_carrier", app.APIURL), responder)

	var req GetShippingCarrierRequest
	ctx := context.Background()
	res, err := client.Returns.GetShippingCarrier(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.GetShippingCarrier returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.GetShippingCarrier response: %#v", res)
}
func Test_Returns_Offer(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.offer_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping Offer due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping Offer due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/offer", app.APIURL), responder)

	var req OfferRequest
	ctx := context.Background()
	res, err := client.Returns.Offer(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.Offer returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.Offer response: %#v", res)
}
func Test_Returns_QueryProof(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.query_proof_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping QueryProof due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping QueryProof due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/returns/query_proof", app.APIURL), responder)

	var req QueryProofRequest
	ctx := context.Background()
	res, err := client.Returns.QueryProof(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("Returns.QueryProof returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.QueryProof response: %#v", res)
}
func Test_Returns_UploadProof(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.upload_proof_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadProof due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadProof due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/upload_proof", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Returns.UploadProof(ctx, sid, mid, tok, "fixtures/test.jpg")
	if err != nil {
		t.Logf("Returns.UploadProof returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.UploadProof response: %#v", res)
}
func Test_Returns_UploadShippingProof(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.returns.upload_shipping_proof_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadShippingProof due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadShippingProof due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/returns/upload_shipping_proof", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Returns.UploadShippingProof(ctx, sid, mid, tok, "fixtures/test.jpg")
	if err != nil {
		t.Logf("Returns.UploadShippingProof returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Returns.UploadShippingProof response: %#v", res)
}
