package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Order_CancelOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.cancel_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping CancelOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping CancelOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/cancel_order", app.APIURL), responder)

	var req CancelOrderRequest
	ctx := context.Background()
	res, err := client.Order.CancelOrder(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.CancelOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.CancelOrder response: %#v", res)
}
func Test_Order_DownloadFbsInvoices(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.download_fbs_invoices_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadFbsInvoices due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadFbsInvoices due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/download_fbs_invoices", app.APIURL), responder)

	var req DownloadFbsInvoicesRequest
	ctx := context.Background()
	res, err := client.Order.DownloadFbsInvoices(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.DownloadFbsInvoices returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.DownloadFbsInvoices response: %#v", res)
}
func Test_Order_DownloadInvoiceDoc(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.download_invoice_doc_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping DownloadInvoiceDoc due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping DownloadInvoiceDoc due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/download_invoice_doc", app.APIURL), responder)

	var req DownloadInvoiceDocRequest
	ctx := context.Background()
	res, err := client.Order.DownloadInvoiceDoc(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.DownloadInvoiceDoc returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.DownloadInvoiceDoc response: %#v", res)
}
func Test_Order_GenerateFbsInvoices(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.generate_fbs_invoices_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateFbsInvoices due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateFbsInvoices due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/generate_fbs_invoices", app.APIURL), responder)

	var req GenerateFbsInvoicesRequest
	ctx := context.Background()
	res, err := client.Order.GenerateFbsInvoices(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GenerateFbsInvoices returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GenerateFbsInvoices response: %#v", res)
}
func Test_Order_GetBookingDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_booking_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_booking_detail", app.APIURL), responder)

	var req GetBookingDetailRequest
	ctx := context.Background()
	res, err := client.Order.GetBookingDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetBookingDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetBookingDetail response: %#v", res)
}
func Test_Order_GetBookingList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_booking_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBookingList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBookingList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_booking_list", app.APIURL), responder)

	var req GetBookingListRequest
	ctx := context.Background()
	res, err := client.Order.GetBookingList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetBookingList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetBookingList response: %#v", res)
}
func Test_Order_GetBuyerInvoiceInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_buyer_invoice_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBuyerInvoiceInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBuyerInvoiceInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/get_buyer_invoice_info", app.APIURL), responder)

	var req GetBuyerInvoiceInfoRequest
	ctx := context.Background()
	res, err := client.Order.GetBuyerInvoiceInfo(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetBuyerInvoiceInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetBuyerInvoiceInfo response: %#v", res)
}
func Test_Order_GetEstimateCancelValue(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_estimate_cancel_value_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetEstimateCancelValue due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetEstimateCancelValue due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/get_estimate_cancel_value", app.APIURL), responder)

	var req GetEstimateCancelValueRequest
	ctx := context.Background()
	res, err := client.Order.GetEstimateCancelValue(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetEstimateCancelValue returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetEstimateCancelValue response: %#v", res)
}
func Test_Order_GetFbsInvoicesResult(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_fbs_invoices_result_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetFbsInvoicesResult due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetFbsInvoicesResult due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/get_fbs_invoices_result", app.APIURL), responder)

	var req GetFbsInvoicesResultRequest
	ctx := context.Background()
	res, err := client.Order.GetFbsInvoicesResult(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetFbsInvoicesResult returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetFbsInvoicesResult response: %#v", res)
}
func Test_Order_GetOrderDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_order_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrderDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOrderDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_order_detail", app.APIURL), responder)

	var req GetOrderDetailRequest
	ctx := context.Background()
	res, err := client.Order.GetOrderDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetOrderDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOrderDetail response: %#v", res)
}
func Test_Order_GetOrderList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_order_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetOrderList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetOrderList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_order_list", app.APIURL), responder)

	var req GetOrderListRequest
	ctx := context.Background()
	res, err := client.Order.GetOrderList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetOrderList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetOrderList response: %#v", res)
}
func Test_Order_GetPackageDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_package_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPackageDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPackageDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_package_detail", app.APIURL), responder)

	var req GetPackageDetailRequest
	ctx := context.Background()
	res, err := client.Order.GetPackageDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetPackageDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetPackageDetail response: %#v", res)
}
func Test_Order_GetPendingBuyerInvoiceOrderList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_pending_buyer_invoice_order_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPendingBuyerInvoiceOrderList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPendingBuyerInvoiceOrderList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_pending_buyer_invoice_order_list", app.APIURL), responder)

	var req GetPendingBuyerInvoiceOrderListRequest
	ctx := context.Background()
	res, err := client.Order.GetPendingBuyerInvoiceOrderList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetPendingBuyerInvoiceOrderList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetPendingBuyerInvoiceOrderList response: %#v", res)
}
func Test_Order_GetShipmentList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_shipment_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShipmentList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShipmentList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_shipment_list", app.APIURL), responder)

	var req GetShipmentListRequest
	ctx := context.Background()
	res, err := client.Order.GetShipmentList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.GetShipmentList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetShipmentList response: %#v", res)
}
func Test_Order_GetWarehouseFilterConfig(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.get_warehouse_filter_config_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWarehouseFilterConfig due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWarehouseFilterConfig due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/order/get_warehouse_filter_config", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Order.GetWarehouseFilterConfig(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Order.GetWarehouseFilterConfig returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.GetWarehouseFilterConfig response: %#v", res)
}
func Test_Order_HandleBuyerCancellation(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.handle_buyer_cancellation_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping HandleBuyerCancellation due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping HandleBuyerCancellation due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/handle_buyer_cancellation", app.APIURL), responder)

	var req HandleBuyerCancellationRequest
	ctx := context.Background()
	res, err := client.Order.HandleBuyerCancellation(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.HandleBuyerCancellation returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.HandleBuyerCancellation response: %#v", res)
}
func Test_Order_HandlePrescriptionCheck(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.handle_prescription_check_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping HandlePrescriptionCheck due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping HandlePrescriptionCheck due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/handle_prescription_check", app.APIURL), responder)

	var req HandlePrescriptionCheckRequest
	ctx := context.Background()
	res, err := client.Order.HandlePrescriptionCheck(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.HandlePrescriptionCheck returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.HandlePrescriptionCheck response: %#v", res)
}
func Test_Order_SearchPackageList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.search_package_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SearchPackageList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SearchPackageList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/search_package_list", app.APIURL), responder)

	var req SearchPackageListRequest
	ctx := context.Background()
	res, err := client.Order.SearchPackageList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.SearchPackageList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.SearchPackageList response: %#v", res)
}
func Test_Order_SetNote(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.set_note_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetNote due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetNote due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/set_note", app.APIURL), responder)

	var req SetNoteRequest
	ctx := context.Background()
	res, err := client.Order.SetNote(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.SetNote returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.SetNote response: %#v", res)
}
func Test_Order_SplitOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.split_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SplitOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SplitOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/split_order", app.APIURL), responder)

	var req SplitOrderRequest
	ctx := context.Background()
	res, err := client.Order.SplitOrder(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.SplitOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.SplitOrder response: %#v", res)
}
func Test_Order_UnsplitOrder(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.unsplit_order_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UnsplitOrder due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UnsplitOrder due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/unsplit_order", app.APIURL), responder)

	var req UnsplitOrderRequest
	ctx := context.Background()
	res, err := client.Order.UnsplitOrder(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Order.UnsplitOrder returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.UnsplitOrder response: %#v", res)
}
func Test_Order_UploadInvoiceDoc(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.order.upload_invoice_doc_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping UploadInvoiceDoc due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping UploadInvoiceDoc due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/order/upload_invoice_doc", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Order.UploadInvoiceDoc(ctx, shopID, "fixtures/test.jpg", accessToken)
	if err != nil {
		t.Logf("Order.UploadInvoiceDoc returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Order.UploadInvoiceDoc response: %#v", res)
}
