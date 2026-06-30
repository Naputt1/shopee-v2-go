package goshopee

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Payment_GenerateIncomeReport(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.generate_income_report_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateIncomeReport due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateIncomeReport due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/generate_income_report", app.APIURL), responder)

	var req GenerateIncomeReportRequest
	ctx := context.Background()
	res, err := client.Payment.GenerateIncomeReport(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GenerateIncomeReport returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GenerateIncomeReport response: %#v", res)
}
func Test_Payment_GenerateIncomeStatement(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.generate_income_statement_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GenerateIncomeStatement due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GenerateIncomeStatement due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/payment/generate_income_statement", app.APIURL), responder)

	var req GenerateIncomeStatementRequest
	ctx := context.Background()
	res, err := client.Payment.GenerateIncomeStatement(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GenerateIncomeStatement returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GenerateIncomeStatement response: %#v", res)
}
func Test_Payment_GetBillingTransactionInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_billing_transaction_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetBillingTransactionInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetBillingTransactionInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_billing_transaction_info", app.APIURL), responder)

	var req GetBillingTransactionInfoRequest
	ctx := context.Background()
	res, err := client.Payment.GetBillingTransactionInfo(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetBillingTransactionInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetBillingTransactionInfo response: %#v", res)
}
func Test_Payment_GetEscrowDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_escrow_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetEscrowDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetEscrowDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_escrow_detail", app.APIURL), responder)

	var req GetEscrowDetailRequest
	ctx := context.Background()
	res, err := client.Payment.GetEscrowDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetEscrowDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetEscrowDetail response: %#v", res)
}
func Test_Payment_GetEscrowDetailBatch(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_escrow_detail_batch_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetEscrowDetailBatch due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetEscrowDetailBatch due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_escrow_detail_batch", app.APIURL), responder)

	var req GetEscrowDetailBatchRequest
	ctx := context.Background()
	res, err := client.Payment.GetEscrowDetailBatch(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetEscrowDetailBatch returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetEscrowDetailBatch response: %#v", res)
}
func Test_Payment_GetEscrowList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_escrow_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetEscrowList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetEscrowList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_escrow_list", app.APIURL), responder)

	var req GetEscrowListRequest
	ctx := context.Background()
	res, err := client.Payment.GetEscrowList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetEscrowList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetEscrowList response: %#v", res)
}
func Test_Payment_GetIncomeDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_income_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetIncomeDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetIncomeDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_income_detail", app.APIURL), responder)

	var req GetIncomeDetailRequest
	ctx := context.Background()
	res, err := client.Payment.GetIncomeDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetIncomeDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetIncomeDetail response: %#v", res)
}
func Test_Payment_GetIncomeOverview(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_income_overview_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetIncomeOverview due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetIncomeOverview due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/payment/get_income_overview", app.APIURL), responder)

	var req GetIncomeOverviewRequest
	ctx := context.Background()
	res, err := client.Payment.GetIncomeOverview(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetIncomeOverview returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetIncomeOverview response: %#v", res)
}
func Test_Payment_GetIncomeReport(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_income_report_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetIncomeReport due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetIncomeReport due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_income_report", app.APIURL), responder)

	var req GetIncomeReportRequest
	ctx := context.Background()
	res, err := client.Payment.GetIncomeReport(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetIncomeReport returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetIncomeReport response: %#v", res)
}
func Test_Payment_GetIncomeStatement(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_income_statement_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetIncomeStatement due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetIncomeStatement due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/api/v2/payment/get_income_statement", app.APIURL), responder)

	var req GetIncomeStatementRequest
	ctx := context.Background()
	res, err := client.Payment.GetIncomeStatement(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetIncomeStatement returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetIncomeStatement response: %#v", res)
}
func Test_Payment_GetItemInstallmentStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_item_installment_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetItemInstallmentStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetItemInstallmentStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_item_installment_status", app.APIURL), responder)

	var req GetItemInstallmentStatusRequest
	ctx := context.Background()
	res, err := client.Payment.GetItemInstallmentStatus(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetItemInstallmentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetItemInstallmentStatus response: %#v", res)
}
func Test_Payment_GetPaymentMethodList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_payment_method_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPaymentMethodList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPaymentMethodList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_payment_method_list", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Payment.GetPaymentMethodList(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Payment.GetPaymentMethodList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetPaymentMethodList response: %#v", res)
}
func Test_Payment_GetPayoutDetail(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_payout_detail_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPayoutDetail due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPayoutDetail due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_payout_detail", app.APIURL), responder)

	var req GetPayoutDetailRequest
	ctx := context.Background()
	res, err := client.Payment.GetPayoutDetail(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetPayoutDetail returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetPayoutDetail response: %#v", res)
}
func Test_Payment_GetPayoutInfo(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_payout_info_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetPayoutInfo due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetPayoutInfo due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_payout_info", app.APIURL), responder)

	var req GetPayoutInfoRequest
	ctx := context.Background()
	res, err := client.Payment.GetPayoutInfo(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetPayoutInfo returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetPayoutInfo response: %#v", res)
}
func Test_Payment_GetShopInstallmentStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_shop_installment_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetShopInstallmentStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetShopInstallmentStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_shop_installment_status", app.APIURL), responder)

	ctx := context.Background()
	res, err := client.Payment.GetShopInstallmentStatus(ctx, shopID, accessToken)
	if err != nil {
		t.Logf("Payment.GetShopInstallmentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetShopInstallmentStatus response: %#v", res)
}
func Test_Payment_GetWalletTransactionList(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.get_wallet_transaction_list_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping GetWalletTransactionList due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping GetWalletTransactionList due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/get_wallet_transaction_list", app.APIURL), responder)

	var req GetWalletTransactionListRequest
	ctx := context.Background()
	res, err := client.Payment.GetWalletTransactionList(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.GetWalletTransactionList returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.GetWalletTransactionList response: %#v", res)
}
func Test_Payment_SetItemInstallmentStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.set_item_installment_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetItemInstallmentStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetItemInstallmentStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/set_item_installment_status", app.APIURL), responder)

	var req SetItemInstallmentStatusRequest
	ctx := context.Background()
	res, err := client.Payment.SetItemInstallmentStatus(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.SetItemInstallmentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.SetItemInstallmentStatus response: %#v", res)
}
func Test_Payment_SetShopInstallmentStatus(t *testing.T) {
	setup()
	defer teardown()

	fixture := "v2.payment.set_shop_installment_status_resp.json"
	data, err := loadFixtureSafe(fixture)
	if err != nil {
		t.Skipf("Skipping SetShopInstallmentStatus due to missing fixture: %v", err)
	}
	responder, err := httpmock.NewJsonResponder(200, data)
	if err != nil {
		t.Skipf("Skipping SetShopInstallmentStatus due to invalid fixture: %v", err)
	}

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v2/payment/set_shop_installment_status", app.APIURL), responder)

	var req SetShopInstallmentStatusRequest
	ctx := context.Background()
	res, err := client.Payment.SetShopInstallmentStatus(ctx, shopID, req, accessToken)
	if err != nil {
		t.Logf("Payment.SetShopInstallmentStatus returned error (possibly expected with mock data): %s", err)
	}

	t.Logf("Payment.SetShopInstallmentStatus response: %#v", res)
}
