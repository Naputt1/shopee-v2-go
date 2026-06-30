package goshopee

import (
	"context"
)

type PaymentService interface {
	// GenerateIncomeReport Trigger income report generation.
	// Path: /api/v2/payment/generate_income_report
	// https://open.shopee.com/documents/v2/v2.payment.generate_income_report?module=97&type=1
	GenerateIncomeReport(ctx context.Context, sid uint64, req GenerateIncomeReportRequest, tok string) (*GenerateIncomeReportResponse, error)
	// GenerateIncomeStatement Trigger income statement generation.
	// Path: /api/v2/payment/generate_income_statement
	// https://open.shopee.com/documents/v2/v2.payment.generate_income_statement?module=97&type=1
	GenerateIncomeStatement(ctx context.Context, sid uint64, opt GenerateIncomeStatementRequest, tok string) (*GenerateIncomeStatementResponse, error)
	// GetBillingTransactionInfo This API is applicable for Cross Border (CB) sellers only to get the detailed payout transaction data, both released and to-be released transaction can be found in here
	// Path: /api/v2/payment/get_billing_transaction_info
	// https://open.shopee.com/documents/v2/v2.payment.get_billing_transaction_info?module=97&type=1
	GetBillingTransactionInfo(ctx context.Context, sid uint64, req GetBillingTransactionInfoRequest, tok string) (*GetBillingTransactionInfoResponse, error)
	// GetEscrowDetail {"content":"<p>Use this API to fetch the accounting detail of order.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to fetch the accounting detail of order."}]}]}
	// Path: /api/v2/payment/get_escrow_detail
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=97&type=1
	GetEscrowDetail(ctx context.Context, sid uint64, req GetEscrowDetailRequest, tok string) (*GetEscrowDetailResponse, error)
	// GetEscrowDetailBatch {"content":"<p>Use this API to fetch the details of order income by batch.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to fetch the details of order income by batch."}]}]}
	// Path: /api/v2/payment/get_escrow_detail_batch
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail_batch?module=97&type=1
	GetEscrowDetailBatch(ctx context.Context, sid uint64, req GetEscrowDetailBatchRequest, tok string) (*GetEscrowDetailBatchResponse, error)
	// GetEscrowList Use this API to fetch the accounting list of order.
	// Path: /api/v2/payment/get_escrow_list
	// https://open.shopee.com/documents/v2/v2.payment.get_escrow_list?module=97&type=1
	GetEscrowList(ctx context.Context, sid uint64, req GetEscrowListRequest, tok string) (*GetEscrowListResponse, error)
	// GetIncomeDetail Retrieves detailed order-level income information across various income statuses for a specified time period. This API enables partners to display granular transaction-level income data consistent with Seller Center’s “Income Details” view, segmented by income status and payout stage.
	//
	// The API dynamically adapts data fields based on the seller’s shop type (Local or Cross Border) and the selected income status (e.g., Pending, To Release, Released).
	// Path: /api/v2/payment/get_income_detail
	// https://open.shopee.com/documents/v2/v2.payment.get_income_detail?module=97&type=1
	GetIncomeDetail(ctx context.Context, sid uint64, req GetIncomeDetailRequest, tok string) (*GetIncomeDetailResponse, error)
	// GetIncomeOverview Retrieves a consolidated snapshot of the seller’s income amounts categorized by income status for a specified shop. This API provides a holistic overview similar to Seller Center’s “Income Overview” section, allowing external systems to reflect the same current payout view.
	//
	// Data is dynamically determined based on the shop type (Local or Cross Border) and the income status requested. Historical income results are not retrievable, providing consistent information as Seller Centre.
	// Path: /api/v2/payment/get_income_overview
	// https://open.shopee.com/documents/v2/v2.payment.get_income_overview?module=97&type=1
	GetIncomeOverview(ctx context.Context, sid uint64, opt GetIncomeOverviewRequest, tok string) (*GetIncomeOverviewResponse, error)
	// GetIncomeReport To query income report status and provide file link if the income report is ready to be downloaded.
	// Path: /api/v2/payment/get_income_report
	// https://open.shopee.com/documents/v2/v2.payment.get_income_report?module=97&type=1
	GetIncomeReport(ctx context.Context, sid uint64, req GetIncomeReportRequest, tok string) (*GetIncomeReportResponse, error)
	// GetIncomeStatement To query income statement status and provide file link if the income statement is ready to be downloaded.
	// Path: /api/v2/payment/get_income_statement
	// https://open.shopee.com/documents/v2/v2.payment.get_income_statement?module=97&type=1
	GetIncomeStatement(ctx context.Context, sid uint64, opt GetIncomeStatementRequest, tok string) (*GetIncomeStatementResponse, error)
	// GetItemInstallmentStatus Get item installment tenures.Only for TH、TW.
	// Path: /api/v2/payment/get_item_installment_status
	// https://open.shopee.com/documents/v2/v2.payment.get_item_installment_status?module=97&type=1
	GetItemInstallmentStatus(ctx context.Context, sid uint64, req GetItemInstallmentStatusRequest, tok string) (*GetItemInstallmentStatusResponse, error)
	// GetPaymentMethodList Obtain payment method (no authentication required)
	// Path: /api/v2/payment/get_payment_method_list
	// https://open.shopee.com/documents/v2/v2.payment.get_payment_method_list?module=97&type=1
	GetPaymentMethodList(ctx context.Context, sid uint64, tok string) (*GetPaymentMethodListResponse, error)
	// GetPayoutDetail This API is applicable for Cross Border (CB) sellers only to get the shop's payout data, such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
	// Path: /api/v2/payment/get_payout_detail
	// https://open.shopee.com/documents/v2/v2.payment.get_payout_detail?module=97&type=1
	GetPayoutDetail(ctx context.Context, sid uint64, req GetPayoutDetailRequest, tok string) (*GetPayoutDetailResponse, error)
	// GetPayoutInfo This is a new API which applicable for Cross Border (CB) sellers only to get the shop's payout data, will be used for the original API v2.get_payout_details replacement, we provide data such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
	// Path: /api/v2/payment/get_payout_info
	// https://open.shopee.com/documents/v2/v2.payment.get_payout_info?module=97&type=1
	GetPayoutInfo(ctx context.Context, sid uint64, req GetPayoutInfoRequest, tok string) (*GetPayoutInfoResponse, error)
	// GetShopInstallmentStatus Get the installment state of shop.
	// Path: /api/v2/payment/get_shop_installment_status
	// https://open.shopee.com/documents/v2/v2.payment.get_shop_installment_status?module=97&type=1
	GetShopInstallmentStatus(ctx context.Context, sid uint64, tok string) (*GetShopInstallmentStatusResponse, error)
	// GetWalletTransactionList {"content":"<p>Use this API to get the transaction records of wallet. Only applicable for local shops</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to get the transaction records of wallet. Only applicable for local shops"}]}]}
	// Path: /api/v2/payment/get_wallet_transaction_list
	// https://open.shopee.com/documents/v2/v2.payment.get_wallet_transaction_list?module=97&type=1
	GetWalletTransactionList(ctx context.Context, sid uint64, req GetWalletTransactionListRequest, tok string) (*GetWalletTransactionListResponse, error)
	// SetItemInstallmentStatus Set item installment.Only for TH、TW.
	// Path: /api/v2/payment/set_item_installment_status
	// https://open.shopee.com/documents/v2/v2.payment.set_item_installment_status?module=97&type=1
	SetItemInstallmentStatus(ctx context.Context, sid uint64, req SetItemInstallmentStatusRequest, tok string) (*SetItemInstallmentStatusResponse, error)
	// SetShopInstallmentStatus Sets the staging capability of shop level.
	// Path: /api/v2/payment/set_shop_installment_status
	// https://open.shopee.com/documents/v2/v2.payment.set_shop_installment_status?module=97&type=1
	SetShopInstallmentStatus(ctx context.Context, sid uint64, req SetShopInstallmentStatusRequest, tok string) (*SetShopInstallmentStatusResponse, error)
}

type PaymentServiceOp[T any] struct {
	client *Client[T]
}

// GenerateIncomeReport Trigger income report generation.
// Path: /api/v2/payment/generate_income_report
// https://open.shopee.com/documents/v2/v2.payment.generate_income_report?module=97&type=1
func (s *PaymentServiceOp[T]) GenerateIncomeReport(ctx context.Context, sid uint64, req GenerateIncomeReportRequest, tok string) (*GenerateIncomeReportResponse, error) {
	path := "/payment/generate_income_report"
	resp := new(GenerateIncomeReportResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GenerateIncomeStatement Trigger income statement generation.
// Path: /api/v2/payment/generate_income_statement
// https://open.shopee.com/documents/v2/v2.payment.generate_income_statement?module=97&type=1
func (s *PaymentServiceOp[T]) GenerateIncomeStatement(ctx context.Context, sid uint64, opt GenerateIncomeStatementRequest, tok string) (*GenerateIncomeStatementResponse, error) {
	path := "/payment/generate_income_statement"
	resp := new(GenerateIncomeStatementResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetBillingTransactionInfo This API is applicable for Cross Border (CB) sellers only to get the detailed payout transaction data, both released and to-be released transaction can be found in here
// Path: /api/v2/payment/get_billing_transaction_info
// https://open.shopee.com/documents/v2/v2.payment.get_billing_transaction_info?module=97&type=1
func (s *PaymentServiceOp[T]) GetBillingTransactionInfo(ctx context.Context, sid uint64, req GetBillingTransactionInfoRequest, tok string) (*GetBillingTransactionInfoResponse, error) {
	path := "/payment/get_billing_transaction_info"
	resp := new(GetBillingTransactionInfoResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetEscrowDetail {"content":"<p>Use this API to fetch the accounting detail of order.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to fetch the accounting detail of order."}]}]}
// Path: /api/v2/payment/get_escrow_detail
// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail?module=97&type=1
func (s *PaymentServiceOp[T]) GetEscrowDetail(ctx context.Context, sid uint64, req GetEscrowDetailRequest, tok string) (*GetEscrowDetailResponse, error) {
	path := "/payment/get_escrow_detail"
	resp := new(GetEscrowDetailResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetEscrowDetailBatch {"content":"<p>Use this API to fetch the details of order income by batch.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to fetch the details of order income by batch."}]}]}
// Path: /api/v2/payment/get_escrow_detail_batch
// https://open.shopee.com/documents/v2/v2.payment.get_escrow_detail_batch?module=97&type=1
func (s *PaymentServiceOp[T]) GetEscrowDetailBatch(ctx context.Context, sid uint64, req GetEscrowDetailBatchRequest, tok string) (*GetEscrowDetailBatchResponse, error) {
	path := "/payment/get_escrow_detail_batch"
	resp := new(GetEscrowDetailBatchResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetEscrowList Use this API to fetch the accounting list of order.
// Path: /api/v2/payment/get_escrow_list
// https://open.shopee.com/documents/v2/v2.payment.get_escrow_list?module=97&type=1
func (s *PaymentServiceOp[T]) GetEscrowList(ctx context.Context, sid uint64, req GetEscrowListRequest, tok string) (*GetEscrowListResponse, error) {
	path := "/payment/get_escrow_list"
	resp := new(GetEscrowListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetIncomeDetail Retrieves detailed order-level income information across various income statuses for a specified time period. This API enables partners to display granular transaction-level income data consistent with Seller Center’s “Income Details” view, segmented by income status and payout stage.
//
// The API dynamically adapts data fields based on the seller’s shop type (Local or Cross Border) and the selected income status (e.g., Pending, To Release, Released).
// Path: /api/v2/payment/get_income_detail
// https://open.shopee.com/documents/v2/v2.payment.get_income_detail?module=97&type=1
func (s *PaymentServiceOp[T]) GetIncomeDetail(ctx context.Context, sid uint64, req GetIncomeDetailRequest, tok string) (*GetIncomeDetailResponse, error) {
	path := "/payment/get_income_detail"
	resp := new(GetIncomeDetailResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetIncomeOverview Retrieves a consolidated snapshot of the seller’s income amounts categorized by income status for a specified shop. This API provides a holistic overview similar to Seller Center’s “Income Overview” section, allowing external systems to reflect the same current payout view.
//
// Data is dynamically determined based on the shop type (Local or Cross Border) and the income status requested. Historical income results are not retrievable, providing consistent information as Seller Centre.
// Path: /api/v2/payment/get_income_overview
// https://open.shopee.com/documents/v2/v2.payment.get_income_overview?module=97&type=1
func (s *PaymentServiceOp[T]) GetIncomeOverview(ctx context.Context, sid uint64, opt GetIncomeOverviewRequest, tok string) (*GetIncomeOverviewResponse, error) {
	path := "/payment/get_income_overview"
	resp := new(GetIncomeOverviewResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetIncomeReport To query income report status and provide file link if the income report is ready to be downloaded.
// Path: /api/v2/payment/get_income_report
// https://open.shopee.com/documents/v2/v2.payment.get_income_report?module=97&type=1
func (s *PaymentServiceOp[T]) GetIncomeReport(ctx context.Context, sid uint64, req GetIncomeReportRequest, tok string) (*GetIncomeReportResponse, error) {
	path := "/payment/get_income_report"
	resp := new(GetIncomeReportResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetIncomeStatement To query income statement status and provide file link if the income statement is ready to be downloaded.
// Path: /api/v2/payment/get_income_statement
// https://open.shopee.com/documents/v2/v2.payment.get_income_statement?module=97&type=1
func (s *PaymentServiceOp[T]) GetIncomeStatement(ctx context.Context, sid uint64, opt GetIncomeStatementRequest, tok string) (*GetIncomeStatementResponse, error) {
	path := "/payment/get_income_statement"
	resp := new(GetIncomeStatementResponse)
	err := s.client.WithShop(sid, tok).Get(ctx, path, resp, opt)
	return resp, err
}

// GetItemInstallmentStatus Get item installment tenures.Only for TH、TW.
// Path: /api/v2/payment/get_item_installment_status
// https://open.shopee.com/documents/v2/v2.payment.get_item_installment_status?module=97&type=1
func (s *PaymentServiceOp[T]) GetItemInstallmentStatus(ctx context.Context, sid uint64, req GetItemInstallmentStatusRequest, tok string) (*GetItemInstallmentStatusResponse, error) {
	path := "/payment/get_item_installment_status"
	resp := new(GetItemInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetPaymentMethodList Obtain payment method (no authentication required)
// Path: /api/v2/payment/get_payment_method_list
// https://open.shopee.com/documents/v2/v2.payment.get_payment_method_list?module=97&type=1
func (s *PaymentServiceOp[T]) GetPaymentMethodList(ctx context.Context, sid uint64, tok string) (*GetPaymentMethodListResponse, error) {
	path := "/payment/get_payment_method_list"
	resp := new(GetPaymentMethodListResponse)
	err := s.client.Post(ctx, path, nil, resp)
	return resp, err
}

// GetPayoutDetail This API is applicable for Cross Border (CB) sellers only to get the shop's payout data, such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
// Path: /api/v2/payment/get_payout_detail
// https://open.shopee.com/documents/v2/v2.payment.get_payout_detail?module=97&type=1
func (s *PaymentServiceOp[T]) GetPayoutDetail(ctx context.Context, sid uint64, req GetPayoutDetailRequest, tok string) (*GetPayoutDetailResponse, error) {
	path := "/payment/get_payout_detail"
	resp := new(GetPayoutDetailResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetPayoutInfo This is a new API which applicable for Cross Border (CB) sellers only to get the shop's payout data, will be used for the original API v2.get_payout_details replacement, we provide data such as the payout amount, currency, FX rate, the payout's associated order income and adjustment records etc.
// Path: /api/v2/payment/get_payout_info
// https://open.shopee.com/documents/v2/v2.payment.get_payout_info?module=97&type=1
func (s *PaymentServiceOp[T]) GetPayoutInfo(ctx context.Context, sid uint64, req GetPayoutInfoRequest, tok string) (*GetPayoutInfoResponse, error) {
	path := "/payment/get_payout_info"
	resp := new(GetPayoutInfoResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// GetShopInstallmentStatus Get the installment state of shop.
// Path: /api/v2/payment/get_shop_installment_status
// https://open.shopee.com/documents/v2/v2.payment.get_shop_installment_status?module=97&type=1
func (s *PaymentServiceOp[T]) GetShopInstallmentStatus(ctx context.Context, sid uint64, tok string) (*GetShopInstallmentStatusResponse, error) {
	path := "/payment/get_shop_installment_status"
	resp := new(GetShopInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, nil, resp)
	return resp, err
}

// GetWalletTransactionList {"content":"<p>Use this API to get the transaction records of wallet. Only applicable for local shops</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this API to get the transaction records of wallet. Only applicable for local shops"}]}]}
// Path: /api/v2/payment/get_wallet_transaction_list
// https://open.shopee.com/documents/v2/v2.payment.get_wallet_transaction_list?module=97&type=1
func (s *PaymentServiceOp[T]) GetWalletTransactionList(ctx context.Context, sid uint64, req GetWalletTransactionListRequest, tok string) (*GetWalletTransactionListResponse, error) {
	path := "/payment/get_wallet_transaction_list"
	resp := new(GetWalletTransactionListResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// SetItemInstallmentStatus Set item installment.Only for TH、TW.
// Path: /api/v2/payment/set_item_installment_status
// https://open.shopee.com/documents/v2/v2.payment.set_item_installment_status?module=97&type=1
func (s *PaymentServiceOp[T]) SetItemInstallmentStatus(ctx context.Context, sid uint64, req SetItemInstallmentStatusRequest, tok string) (*SetItemInstallmentStatusResponse, error) {
	path := "/payment/set_item_installment_status"
	resp := new(SetItemInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}

// SetShopInstallmentStatus Sets the staging capability of shop level.
// Path: /api/v2/payment/set_shop_installment_status
// https://open.shopee.com/documents/v2/v2.payment.set_shop_installment_status?module=97&type=1
func (s *PaymentServiceOp[T]) SetShopInstallmentStatus(ctx context.Context, sid uint64, req SetShopInstallmentStatusRequest, tok string) (*SetShopInstallmentStatusResponse, error) {
	path := "/payment/set_shop_installment_status"
	resp := new(SetShopInstallmentStatusResponse)
	err := s.client.WithShop(sid, tok).Post(ctx, path, req, resp)
	return resp, err
}
