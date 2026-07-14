package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_FirstMile_BindCourierDeliveryFirstMileTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.bind_courier_delivery_first_mile_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BindCourierDeliveryFirstMileTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BindCourierDeliveryFirstMileTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/bind_courier_delivery_first_mile_tracking_number", app.APIURL), responder)

	var req BindCourierDeliveryFirstMileTrackingNumberRequest
	ctx := context.Background()
	res, err := client.FirstMile.BindCourierDeliveryFirstMileTrackingNumber(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.BindCourierDeliveryFirstMileTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.BindCourierDeliveryFirstMileTrackingNumber response: %#v", res)
}
func Test_FirstMile_BindFirstMileTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.bind_first_mile_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BindFirstMileTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BindFirstMileTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/bind_first_mile_tracking_number", app.APIURL), responder)

	var req BindFirstMileTrackingNumberRequest
	ctx := context.Background()
	res, err := client.FirstMile.BindFirstMileTrackingNumber(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.BindFirstMileTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.BindFirstMileTrackingNumber response: %#v", res)
}
func Test_FirstMile_GenerateAndBindFirstMileTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.generate_and_bind_first_mile_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateAndBindFirstMileTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateAndBindFirstMileTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/generate_and_bind_first_mile_tracking_number", app.APIURL), responder)

	var req GenerateAndBindFirstMileTrackingNumberRequest
	ctx := context.Background()
	res, err := client.FirstMile.GenerateAndBindFirstMileTrackingNumber(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GenerateAndBindFirstMileTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GenerateAndBindFirstMileTrackingNumber response: %#v", res)
}
func Test_FirstMile_GenerateFirstMileTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.generate_first_mile_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateFirstMileTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateFirstMileTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/generate_first_mile_tracking_number", app.APIURL), responder)

	var req GenerateFirstMileTrackingNumberRequest
	ctx := context.Background()
	res, err := client.FirstMile.GenerateFirstMileTrackingNumber(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GenerateFirstMileTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GenerateFirstMileTrackingNumber response: %#v", res)
}
func Test_FirstMile_GetChannelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_channel_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChannelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetChannelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/get_channel_list", app.APIURL), responder)

	var req GetChannelListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetChannelList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetChannelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetChannelList response: %#v", res)
}
func Test_FirstMile_GetCourierDeliveryChannelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_courier_delivery_channel_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryChannelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryChannelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_courier_delivery_channel_list", app.APIURL), responder)

	var req GetCourierDeliveryChannelListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetCourierDeliveryChannelList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetCourierDeliveryChannelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetCourierDeliveryChannelList response: %#v", res)
}
func Test_FirstMile_GetCourierDeliveryDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_courier_delivery_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_courier_delivery_detail", app.APIURL), responder)

	var req GetCourierDeliveryDetailRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetCourierDeliveryDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetCourierDeliveryDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetCourierDeliveryDetail response: %#v", res)
}
func Test_FirstMile_GetCourierDeliveryTrackingNumberList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_courier_delivery_tracking_number_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryTrackingNumberList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryTrackingNumberList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/get_courier_delivery_tracking_number_list", app.APIURL), responder)

	var req GetCourierDeliveryTrackingNumberListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetCourierDeliveryTrackingNumberList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetCourierDeliveryTrackingNumberList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetCourierDeliveryTrackingNumberList response: %#v", res)
}
func Test_FirstMile_GetCourierDeliveryWaybill(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_courier_delivery_waybill_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryWaybill due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetCourierDeliveryWaybill due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/get_courier_delivery_waybill", app.APIURL), responder)

	var req GetCourierDeliveryWaybillRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetCourierDeliveryWaybill(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetCourierDeliveryWaybill returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetCourierDeliveryWaybill response: %#v", res)
}
func Test_FirstMile_GetDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_detail", app.APIURL), responder)

	var req GetDetailRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetDetail(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetDetail response: %#v", res)
}
func Test_FirstMile_GetTrackingNumberList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_tracking_number_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTrackingNumberList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTrackingNumberList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_tracking_number_list", app.APIURL), responder)

	var req GetTrackingNumberListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetTrackingNumberList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetTrackingNumberList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetTrackingNumberList response: %#v", res)
}
func Test_FirstMile_GetTransitWarehouseList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_transit_warehouse_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTransitWarehouseList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTransitWarehouseList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_transit_warehouse_list", app.APIURL), responder)

	var req GetTransitWarehouseListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetTransitWarehouseList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetTransitWarehouseList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetTransitWarehouseList response: %#v", res)
}
func Test_FirstMile_GetUnbindOrderList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_unbind_order_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetUnbindOrderList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetUnbindOrderList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/first_mile/get_unbind_order_list", app.APIURL), responder)

	var req GetUnbindOrderListRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetUnbindOrderList(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetUnbindOrderList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetUnbindOrderList response: %#v", res)
}
func Test_FirstMile_GetWaybill(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.get_waybill_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWaybill due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWaybill due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/get_waybill", app.APIURL), responder)

	var req GetWaybillRequest
	ctx := context.Background()
	res, err := client.FirstMile.GetWaybill(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.GetWaybill returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.GetWaybill response: %#v", res)
}
func Test_FirstMile_UnbindFirstMileTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.unbind_first_mile_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UnbindFirstMileTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UnbindFirstMileTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/unbind_first_mile_tracking_number", app.APIURL), responder)

	var req UnbindFirstMileTrackingNumberRequest
	ctx := context.Background()
	res, err := client.FirstMile.UnbindFirstMileTrackingNumber(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.UnbindFirstMileTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.UnbindFirstMileTrackingNumber response: %#v", res)
}
func Test_FirstMile_UnbindFirstMileTrackingNumberAll(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.first_mile.unbind_first_mile_tracking_number_all_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UnbindFirstMileTrackingNumberAll due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UnbindFirstMileTrackingNumberAll due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/first_mile/unbind_first_mile_tracking_number_all", app.APIURL), responder)

	var req UnbindFirstMileTrackingNumberAllRequest
	ctx := context.Background()
	res, err := client.FirstMile.UnbindFirstMileTrackingNumberAll(ctx, sid, mid, tok, req)
	if err != nil {
		t.Logf("FirstMile.UnbindFirstMileTrackingNumberAll returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("FirstMile.UnbindFirstMileTrackingNumberAll response: %#v", res)
}
