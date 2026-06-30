package goshopee

type BookingStatus string

const (
	BookingStatusCancelled BookingStatus = "CANCELLED"
	BookingStatusMatched BookingStatus = "MATCHED"
	BookingStatusProcessed BookingStatus = "PROCESSED"
	BookingStatusReadyToShip BookingStatus = "READY_TO_SHIP"
	BookingStatusShipped BookingStatus = "SHIPPED"
)
type CampaignStatus string

const (
	CampaignStatusClosed CampaignStatus = "closed"
	CampaignStatusDeleted CampaignStatus = "deleted"
	CampaignStatusEnded CampaignStatus = "ended"
	CampaignStatusExpired CampaignStatus = "expired"
	CampaignStatusOngoing CampaignStatus = "ongoing"
	CampaignStatusPaused CampaignStatus = "paused"
	CampaignStatusScheduled CampaignStatus = "scheduled"
	CampaignStatusUpcoming CampaignStatus = "upcoming"
)
type DescriptionElementFieldType string

const (
	DescriptionElementFieldTypeImage DescriptionElementFieldType = "image"
	DescriptionElementFieldTypeText DescriptionElementFieldType = "text"
)
type DescriptionType string

const (
	DescriptionTypeExtended DescriptionType = "extended"
	DescriptionTypeNormal DescriptionType = "normal"
)
type InvoiceOption string

const (
	InvoiceOptionNoInvoice InvoiceOption = "NO_INVOICE"
	InvoiceOptionNonVat InvoiceOption = "NON_VAT_INVOICES"
	InvoiceOptionVat InvoiceOption = "VAT_INVOICES"
	InvoiceOptionVatMarginScheme InvoiceOption = "VAT_MARGIN_SCHEME_INVOICES"
)
type ItemStatus string

const (
	ItemStatusBanned ItemStatus = "BANNED"
	ItemStatusNormal ItemStatus = "NORMAL"
	ItemStatusUnlist ItemStatus = "UNLIST"
)
type LogisticsStatus string

const (
	LogisticsStatusCodRejected LogisticsStatus = "LOGISTICS_COD_REJECTED"
	LogisticsStatusDeliveryDone LogisticsStatus = "LOGISTICS_DELIVERY_DONE"
	LogisticsStatusDeliveryFailed LogisticsStatus = "LOGISTICS_DELIVERY_FAILED"
	LogisticsStatusInvalid LogisticsStatus = "LOGISTICS_INVALID"
	LogisticsStatusLost LogisticsStatus = "LOGISTICS_LOST"
	LogisticsStatusNotStart LogisticsStatus = "LOGISTICS_NOT_START"
	LogisticsStatusPickupDone LogisticsStatus = "LOGISTICS_PICKUP_DONE"
	LogisticsStatusPickupFailed LogisticsStatus = "LOGISTICS_PICKUP_FAILED"
	LogisticsStatusReady LogisticsStatus = "LOGISTICS_READY"
	LogisticsStatusRequestCanceled LogisticsStatus = "LOGISTICS_REQUEST_CANCELED"
	LogisticsStatusRequestCreated LogisticsStatus = "LOGISTICS_REQUEST_CREATED"
)
type OperationType string

const (
	OperationTypeManufactorer OperationType = "2"
	OperationTypeRetailer OperationType = "1"
)
type OrderStatus string

const (
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusCompleted OrderStatus = "COMPLETED"
	OrderStatusInCancel OrderStatus = "IN_CANCEL"
	OrderStatusInvoicePending OrderStatus = "INVOICE_PENDING"
	OrderStatusProcessed OrderStatus = "PROCESSED"
	OrderStatusReadyToShip OrderStatus = "READY_TO_SHIP"
	OrderStatusShipped OrderStatus = "SHIPPED"
	OrderStatusUnpaid OrderStatus = "UNPAID"
)
type PromotionStatus string

const (
	PromotionStatusAll PromotionStatus = "all"
	PromotionStatusExpired PromotionStatus = "expired"
	PromotionStatusOngoing PromotionStatus = "ongoing"
	PromotionStatusUpcoming PromotionStatus = "upcoming"
)
type ReturnStatus string

const (
	ReturnStatusAccepted ReturnStatus = "ACCEPTED"
	ReturnStatusCancelled ReturnStatus = "CANCELLED"
	ReturnStatusClosed ReturnStatus = "CLOSED"
	ReturnStatusJudging ReturnStatus = "JUDGING"
	ReturnStatusProcessing ReturnStatus = "PROCESSING"
	ReturnStatusRefunding ReturnStatus = "REFUNDING"
	ReturnStatusRequested ReturnStatus = "REQUESTED"
	ReturnStatusSellerDispute ReturnStatus = "SELLER_DISPUTE"
)
type TaxType int

const (
	TaxTypeNoTax TaxType = 0
	TaxTypeTaxable TaxType = 1
	TaxTypeTaxFree TaxType = 2
)
type WarrantyTime string

const (
	WarrantyTimeOneYear WarrantyTime = "ONE_YEAR"
	WarrantyTimeOverTwoYears WarrantyTime = "OVER_TWO_YEARS"
	WarrantyTimeTwoYears WarrantyTime = "TWO_YEARS"
)
