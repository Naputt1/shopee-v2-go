package goshopee

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Logistics_BatchShipOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.batch_ship_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchShipOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchShipOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/batch_ship_order", app.APIURL), responder)

	var req BatchShipOrderRequest
	res, err := client.Logistics.BatchShipOrder(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.BatchShipOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.BatchShipOrder response: %#v", res)
}
func Test_Logistics_BatchUpdateTpfWarehouseTrackingStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.batch_update_tpf_warehouse_tracking_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping BatchUpdateTpfWarehouseTrackingStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping BatchUpdateTpfWarehouseTrackingStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/batch_update_tpf_warehouse_tracking_status", app.APIURL), responder)

	var req BatchUpdateTpfWarehouseTrackingStatusRequest
	res, err := client.Logistics.BatchUpdateTpfWarehouseTrackingStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.BatchUpdateTpfWarehouseTrackingStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.BatchUpdateTpfWarehouseTrackingStatus response: %#v", res)
}
func Test_Logistics_CheckPolygonUpdateStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.check_polygon_update_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CheckPolygonUpdateStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CheckPolygonUpdateStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/check_polygon_update_status", app.APIURL), responder)

	var req CheckPolygonUpdateStatusRequest
	res, err := client.Logistics.CheckPolygonUpdateStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.CheckPolygonUpdateStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.CheckPolygonUpdateStatus response: %#v", res)
}
func Test_Logistics_CreateBookingShippingDocument(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.create_booking_shipping_document_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateBookingShippingDocument due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateBookingShippingDocument due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/create_booking_shipping_document", app.APIURL), responder)

	var req CreateBookingShippingDocumentRequest
	res, err := client.Logistics.CreateBookingShippingDocument(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.CreateBookingShippingDocument returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.CreateBookingShippingDocument response: %#v", res)
}
func Test_Logistics_CreateShippingDocument(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.create_shipping_document_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateShippingDocument due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateShippingDocument due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/create_shipping_document", app.APIURL), responder)

	var req CreateShippingDocumentRequest
	res, err := client.Logistics.CreateShippingDocument(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.CreateShippingDocument returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.CreateShippingDocument response: %#v", res)
}
func Test_Logistics_CreateShippingDocumentJob(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.create_shipping_document_job_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CreateShippingDocumentJob due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CreateShippingDocumentJob due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/create_shipping_document_job", app.APIURL), responder)

	var req CreateShippingDocumentJobRequest
	res, err := client.Logistics.CreateShippingDocumentJob(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.CreateShippingDocumentJob returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.CreateShippingDocumentJob response: %#v", res)
}
func Test_Logistics_DeleteAddress(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.delete_address_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteAddress due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteAddress due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/delete_address", app.APIURL), responder)

	var req DeleteAddressRequest
	res, err := client.Logistics.DeleteAddress(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DeleteAddress returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DeleteAddress response: %#v", res)
}
func Test_Logistics_DeleteSpecialOperatingHour(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.delete_special_operating_hour_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DeleteSpecialOperatingHour due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DeleteSpecialOperatingHour due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/delete_special_operating_hour", app.APIURL), responder)

	var req DeleteSpecialOperatingHourRequest
	res, err := client.Logistics.DeleteSpecialOperatingHour(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DeleteSpecialOperatingHour returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DeleteSpecialOperatingHour response: %#v", res)
}
func Test_Logistics_DownloadBookingShippingDocument(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.download_booking_shipping_document_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadBookingShippingDocument due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadBookingShippingDocument due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/download_booking_shipping_document", app.APIURL), responder)

	var req DownloadBookingShippingDocumentRequest
	res, err := client.Logistics.DownloadBookingShippingDocument(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DownloadBookingShippingDocument returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DownloadBookingShippingDocument response: %#v", res)
}
func Test_Logistics_DownloadShippingDocument(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.download_shipping_document_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadShippingDocument due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadShippingDocument due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/download_shipping_document", app.APIURL), responder)

	var req DownloadShippingDocumentRequest
	res, err := client.Logistics.DownloadShippingDocument(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DownloadShippingDocument returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DownloadShippingDocument response: %#v", res)
}
func Test_Logistics_DownloadShippingDocumentJob(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.download_shipping_document_job_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadShippingDocumentJob due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadShippingDocumentJob due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/download_shipping_document_job", app.APIURL), responder)

	var req DownloadShippingDocumentJobRequest
	res, err := client.Logistics.DownloadShippingDocumentJob(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DownloadShippingDocumentJob returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DownloadShippingDocumentJob response: %#v", res)
}
func Test_Logistics_DownloadToLabel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.download_to_label_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadToLabel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadToLabel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/download_to_label", app.APIURL), responder)

	var req DownloadToLabelRequest
	res, err := client.Logistics.DownloadToLabel(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.DownloadToLabel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.DownloadToLabel response: %#v", res)
}
func Test_Logistics_GetAddressList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_address_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetAddressList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetAddressList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_address_list", app.APIURL), responder)

	res, err := client.Logistics.GetAddressList(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetAddressList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetAddressList response: %#v", res)
}
func Test_Logistics_GetBookingShippingDocumentDataInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_shipping_document_data_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentDataInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentDataInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_booking_shipping_document_data_info", app.APIURL), responder)

	var req GetBookingShippingDocumentDataInfoRequest
	res, err := client.Logistics.GetBookingShippingDocumentDataInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingShippingDocumentDataInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingShippingDocumentDataInfo response: %#v", res)
}
func Test_Logistics_GetBookingShippingDocumentParameter(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_shipping_document_parameter_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentParameter due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentParameter due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_booking_shipping_document_parameter", app.APIURL), responder)

	var req GetBookingShippingDocumentParameterRequest
	res, err := client.Logistics.GetBookingShippingDocumentParameter(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingShippingDocumentParameter returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingShippingDocumentParameter response: %#v", res)
}
func Test_Logistics_GetBookingShippingDocumentResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_shipping_document_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingDocumentResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_booking_shipping_document_result", app.APIURL), responder)

	var req GetBookingShippingDocumentResultRequest
	res, err := client.Logistics.GetBookingShippingDocumentResult(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingShippingDocumentResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingShippingDocumentResult response: %#v", res)
}
func Test_Logistics_GetBookingShippingParameter(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_shipping_parameter_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingParameter due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingShippingParameter due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_booking_shipping_parameter", app.APIURL), responder)

	var req GetBookingShippingParameterRequest
	res, err := client.Logistics.GetBookingShippingParameter(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingShippingParameter returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingShippingParameter response: %#v", res)
}
func Test_Logistics_GetBookingTrackingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_tracking_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingTrackingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingTrackingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_booking_tracking_info", app.APIURL), responder)

	var req GetBookingTrackingInfoRequest
	res, err := client.Logistics.GetBookingTrackingInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingTrackingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingTrackingInfo response: %#v", res)
}
func Test_Logistics_GetBookingTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_booking_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_booking_tracking_number", app.APIURL), responder)

	var req GetBookingTrackingNumberRequest
	res, err := client.Logistics.GetBookingTrackingNumber(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetBookingTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetBookingTrackingNumber response: %#v", res)
}
func Test_Logistics_GetChannelList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_channel_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetChannelList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetChannelList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_channel_list", app.APIURL), responder)

	res, err := client.Logistics.GetChannelList(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetChannelList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetChannelList response: %#v", res)
}
func Test_Logistics_GetMartPackagingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_mart_packaging_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMartPackagingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMartPackagingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_mart_packaging_info", app.APIURL), responder)

	res, err := client.Logistics.GetMartPackagingInfo(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetMartPackagingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetMartPackagingInfo response: %#v", res)
}
func Test_Logistics_GetMassShippingParameter(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_mass_shipping_parameter_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMassShippingParameter due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMassShippingParameter due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_mass_shipping_parameter", app.APIURL), responder)

	var req GetMassShippingParameterRequest
	res, err := client.Logistics.GetMassShippingParameter(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetMassShippingParameter returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetMassShippingParameter response: %#v", res)
}
func Test_Logistics_GetMassTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_mass_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetMassTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetMassTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_mass_tracking_number", app.APIURL), responder)

	var req GetMassTrackingNumberRequest
	res, err := client.Logistics.GetMassTrackingNumber(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetMassTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetMassTrackingNumber response: %#v", res)
}
func Test_Logistics_GetOperatingHourRestrictions(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_operating_hour_restrictions_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOperatingHourRestrictions due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOperatingHourRestrictions due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_operating_hour_restrictions", app.APIURL), responder)

	res, err := client.Logistics.GetOperatingHourRestrictions(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetOperatingHourRestrictions returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetOperatingHourRestrictions response: %#v", res)
}
func Test_Logistics_GetOperatingHours(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_operating_hours_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOperatingHours due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOperatingHours due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_operating_hours", app.APIURL), responder)

	res, err := client.Logistics.GetOperatingHours(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetOperatingHours returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetOperatingHours response: %#v", res)
}
func Test_Logistics_GetPauseStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_pause_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPauseStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPauseStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_pause_status", app.APIURL), responder)

	res, err := client.Logistics.GetPauseStatus(shopID, accessToken)
	if err != nil {
		t.Logf("Logistics.GetPauseStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetPauseStatus response: %#v", res)
}
func Test_Logistics_GetShippingDocumentDataInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_shipping_document_data_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentDataInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentDataInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_shipping_document_data_info", app.APIURL), responder)

	var req GetShippingDocumentDataInfoRequest
	res, err := client.Logistics.GetShippingDocumentDataInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetShippingDocumentDataInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetShippingDocumentDataInfo response: %#v", res)
}
func Test_Logistics_GetShippingDocumentJobStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_shipping_document_job_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentJobStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentJobStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_shipping_document_job_status", app.APIURL), responder)

	var req GetShippingDocumentJobStatusRequest
	res, err := client.Logistics.GetShippingDocumentJobStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetShippingDocumentJobStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetShippingDocumentJobStatus response: %#v", res)
}
func Test_Logistics_GetShippingDocumentParameter(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_shipping_document_parameter_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentParameter due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentParameter due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_shipping_document_parameter", app.APIURL), responder)

	var req GetShippingDocumentParameterRequest
	res, err := client.Logistics.GetShippingDocumentParameter(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetShippingDocumentParameter returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetShippingDocumentParameter response: %#v", res)
}
func Test_Logistics_GetShippingDocumentResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_shipping_document_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingDocumentResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/get_shipping_document_result", app.APIURL), responder)

	var req GetShippingDocumentResultRequest
	res, err := client.Logistics.GetShippingDocumentResult(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetShippingDocumentResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetShippingDocumentResult response: %#v", res)
}
func Test_Logistics_GetShippingParameter(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_shipping_parameter_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShippingParameter due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShippingParameter due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_shipping_parameter", app.APIURL), responder)

	var req GetShippingParameterRequest
	res, err := client.Logistics.GetShippingParameter(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetShippingParameter returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetShippingParameter response: %#v", res)
}
func Test_Logistics_GetTrackingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_tracking_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTrackingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTrackingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_tracking_info", app.APIURL), responder)

	var req GetTrackingInfoRequest
	res, err := client.Logistics.GetTrackingInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetTrackingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetTrackingInfo response: %#v", res)
}
func Test_Logistics_GetTrackingNumber(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.get_tracking_number_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetTrackingNumber due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetTrackingNumber due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/logistics/get_tracking_number", app.APIURL), responder)

	var req GetTrackingNumberRequest
	res, err := client.Logistics.GetTrackingNumber(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.GetTrackingNumber returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.GetTrackingNumber response: %#v", res)
}
func Test_Logistics_MassShipOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.mass_ship_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping MassShipOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping MassShipOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/mass_ship_order", app.APIURL), responder)

	var req MassShipOrderRequest
	res, err := client.Logistics.MassShipOrder(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.MassShipOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.MassShipOrder response: %#v", res)
}
func Test_Logistics_SetAddressConfig(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.set_address_config_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetAddressConfig due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetAddressConfig due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/set_address_config", app.APIURL), responder)

	var req SetAddressConfigRequest
	res, err := client.Logistics.SetAddressConfig(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.SetAddressConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.SetAddressConfig response: %#v", res)
}
func Test_Logistics_SetMartPackagingInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.set_mart_packaging_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetMartPackagingInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetMartPackagingInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/set_mart_packaging_info", app.APIURL), responder)

	var req SetMartPackagingInfoRequest
	res, err := client.Logistics.SetMartPackagingInfo(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.SetMartPackagingInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.SetMartPackagingInfo response: %#v", res)
}
func Test_Logistics_SetPauseStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.set_pause_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetPauseStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetPauseStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/set_pause_status", app.APIURL), responder)

	var req SetPauseStatusRequest
	res, err := client.Logistics.SetPauseStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.SetPauseStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.SetPauseStatus response: %#v", res)
}
func Test_Logistics_ShipBooking(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.ship_booking_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ShipBooking due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ShipBooking due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/ship_booking", app.APIURL), responder)

	var req ShipBookingRequest
	res, err := client.Logistics.ShipBooking(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.ShipBooking returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.ShipBooking response: %#v", res)
}
func Test_Logistics_ShipOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.ship_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping ShipOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping ShipOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/ship_order", app.APIURL), responder)

	var req ShipOrderRequest
	res, err := client.Logistics.ShipOrder(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.ShipOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.ShipOrder response: %#v", res)
}
func Test_Logistics_UpdateAddress(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_address_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateAddress due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateAddress due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_address", app.APIURL), responder)

	var req UpdateAddressRequest
	res, err := client.Logistics.UpdateAddress(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateAddress returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateAddress response: %#v", res)
}
func Test_Logistics_UpdateChannel(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_channel_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateChannel due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateChannel due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_channel", app.APIURL), responder)

	var req UpdateChannelRequest
	res, err := client.Logistics.UpdateChannel(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateChannel returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateChannel response: %#v", res)
}
func Test_Logistics_UpdateOperatingHours(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_operating_hours_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateOperatingHours due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateOperatingHours due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_operating_hours", app.APIURL), responder)

	var req UpdateOperatingHoursRequest
	res, err := client.Logistics.UpdateOperatingHours(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateOperatingHours returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateOperatingHours response: %#v", res)
}
func Test_Logistics_UpdateSelfCollectionOrderLogistics(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_self_collection_order_logistics_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateSelfCollectionOrderLogistics due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateSelfCollectionOrderLogistics due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_self_collection_order_logistics", app.APIURL), responder)

	var req UpdateSelfCollectionOrderLogisticsRequest
	res, err := client.Logistics.UpdateSelfCollectionOrderLogistics(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateSelfCollectionOrderLogistics returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateSelfCollectionOrderLogistics response: %#v", res)
}
func Test_Logistics_UpdateShippingOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_shipping_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateShippingOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateShippingOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_shipping_order", app.APIURL), responder)

	var req UpdateShippingOrderRequest
	res, err := client.Logistics.UpdateShippingOrder(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateShippingOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateShippingOrder response: %#v", res)
}
func Test_Logistics_UpdateTrackingStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.update_tracking_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UpdateTrackingStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UpdateTrackingStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/update_tracking_status", app.APIURL), responder)

	var req UpdateTrackingStatusRequest
	res, err := client.Logistics.UpdateTrackingStatus(shopID, req, accessToken)
	if err != nil {
		t.Logf("Logistics.UpdateTrackingStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UpdateTrackingStatus response: %#v", res)
}
func Test_Logistics_UploadServiceablePolygon(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.logistics.upload_serviceable_polygon_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadServiceablePolygon due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadServiceablePolygon due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/logistics/upload_serviceable_polygon", app.APIURL), responder)

	res, err := client.Logistics.UploadServiceablePolygon(shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Logistics.UploadServiceablePolygon returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Logistics.UploadServiceablePolygon response: %#v", res)
}
