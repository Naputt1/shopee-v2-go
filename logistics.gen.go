package goshopee

import "io"

type LogisticsService interface {
	// BatchShipOrder Use this api to batch initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller.Only channel 90003 - Padrão in Brazil has the permission of this API.
	//
	//
	// Path: /api/v2/logistics/batch_ship_order
	// https://open.shopee.com/documents/v2/v2.logistics.batch_ship_order?module=95&type=1
	BatchShipOrder(sid uint64, req BatchShipOrderRequest, tok string) (*BatchShipOrderResponse, error)
	// BatchUpdateTpfWarehouseTrackingStatus For CB orders that fulfilled by 3PF, support 3PF Warehouse Vendors to update the tpf_tracking_status when 3PF warehouse receive the order and complete the outbound of the package.
	// CB orders that fulfilled by 3PF：
	// v2.shop.get_shop_info  - shop_fulfillment_flag in {Pure - 3PF Shop,PFF - 3PF Shop,LFF Hybrid Shop}
	// And
	// v2.order.get_order_detail -  fulfillment_flag = fulfilled_by_local_seller
	// Path: /api/v2/logistics/batch_update_tpf_warehouse_tracking_status
	// https://open.shopee.com/documents/v2/v2.logistics.batch_update_tpf_warehouse_tracking_status?module=95&type=1
	BatchUpdateTpfWarehouseTrackingStatus(sid uint64, req BatchUpdateTpfWarehouseTrackingStatusRequest, tok string) (*BatchUpdateTpfWarehouseTrackingStatusResponse, error)
	// CheckPolygonUpdateStatus Only available for Brazil sellers. Use this API to check the status of polygon file uploaded for BR Entrega Turbo channel (Channel ID: 90026) by querying the task_id returned via the v2.logistics.upload_serviceable_polygon.
	// Path: /api/v2/logistics/check_polygon_update_status
	// https://open.shopee.com/documents/v2/v2.logistics.check_polygon_update_status?module=95&type=1
	CheckPolygonUpdateStatus(sid uint64, req CheckPolygonUpdateStatusRequest, tok string) (*CheckPolygonUpdateStatusResponse, error)
	// CreateBookingShippingDocument Use this api to create shipping document task for each booking and this API is only available after retrieving the tracking number.
	// Path: /api/v2/logistics/create_booking_shipping_document
	// https://open.shopee.com/documents/v2/v2.logistics.create_booking_shipping_document?module=95&type=1
	CreateBookingShippingDocument(sid uint64, req CreateBookingShippingDocumentRequest, tok string) (*CreateBookingShippingDocumentResponse, error)
	// CreateShippingDocument Use this api to create shipping document task for each order or package and this API is only available after retrieving the tracking number.
	// Path: /api/v2/logistics/create_shipping_document
	// https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document?module=95&type=1
	CreateShippingDocument(sid uint64, req CreateShippingDocumentRequest, tok string) (*CreateShippingDocumentResponse, error)
	// CreateShippingDocumentJob This API creates a shipping document job for selected documents. The system receives requests and returns a job ID along with success and failure details.
	// Path: /api/v2/logistics/create_shipping_document_job
	// https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document_job?module=95&type=1
	CreateShippingDocumentJob(sid uint64, req CreateShippingDocumentJobRequest, tok string) (*CreateShippingDocumentJobResponse, error)
	// DeleteAddress Use this api to delete address.
	// Path: /api/v2/logistics/delete_address
	// https://open.shopee.com/documents/v2/v2.logistics.delete_address?module=95&type=1
	DeleteAddress(sid uint64, req DeleteAddressRequest, tok string) (*DeleteAddressResponse, error)
	// DeleteSpecialOperatingHour This API is used to delete a specific special operating hour for a seller. This API allows sellers to manage their operating hours by removing any special operating hours that are no longer needed. To use this API, the name of the special operating hour to be deleted should be obtained from the v2.logistics.get_operating_hours API.
	// Path: /api/v2/logistics/delete_special_operating_hour
	// https://open.shopee.com/documents/v2/v2.logistics.delete_special_operating_hour?module=95&type=1
	DeleteSpecialOperatingHour(sid uint64, req DeleteSpecialOperatingHourRequest, tok string) (*DeleteSpecialOperatingHourResponse, error)
	// DownloadBookingShippingDocument Use this api to download shipping_document. You have to call v2.logistics.create_booking_shipping_document to create a new shipping document task first and call v2.logistics.get_booking_shipping_document_result to get the task status second. If the task is READY, you can download this shipping document.
	// Path: /api/v2/logistics/download_booking_shipping_document
	// https://open.shopee.com/documents/v2/v2.logistics.download_booking_shipping_document?module=95&type=1
	DownloadBookingShippingDocument(sid uint64, req DownloadBookingShippingDocumentRequest, tok string) (*DownloadBookingShippingDocumentResponse, error)
	// DownloadShippingDocument Use this api to download shipping_document. You have to call v2.logistics.create_shipping_document to create a new shipping document task first and call v2.logistics.get_shipping_document_resut to get the task status second. If the task is READY, you can download this shipping document.
	// Path: /api/v2/logistics/download_shipping_document
	// https://open.shopee.com/documents/v2/v2.logistics.download_shipping_document?module=95&type=1
	DownloadShippingDocument(sid uint64, req DownloadShippingDocumentRequest, tok string) (*DownloadShippingDocumentResponse, error)
	// DownloadShippingDocumentJob This API allows users to download the shipping document associated with a specific job ID. It checks the job status before proceeding with the download.
	// Path: /api/v2/logistics/download_shipping_document_job
	// https://open.shopee.com/documents/v2/v2.logistics.download_shipping_document_job?module=95&type=1
	DownloadShippingDocumentJob(sid uint64, req DownloadShippingDocumentJobRequest, tok string) (*DownloadShippingDocumentJobResponse, error)
	// DownloadToLabel Use the API to download the TO label that should be attached to the carton before drop-off at the warehouse (Only for TW channel_id:30029).
	// Path: /api/v2/logistics/download_to_label
	// https://open.shopee.com/documents/v2/v2.logistics.download_to_label?module=95&type=1
	DownloadToLabel(sid uint64, req DownloadToLabelRequest, tok string) (*DownloadToLabelResponse, error)
	// GetAddressList For integrated logistics channel, use this call to get pickup address for pickup mode order.
	// Path: /api/v2/logistics/get_address_list
	// https://open.shopee.com/documents/v2/v2.logistics.get_address_list?module=95&type=1
	GetAddressList(sid uint64, tok string) (*GetAddressListResponse, error)
	// GetBookingShippingDocumentDataInfo Use this api to fetch the logistics information of a booking these info can be used for airwaybill printing. Dedicated for crossborder SLS order airwaybill. May not be applicable for local channel airwaybill. Besides, this api supports returning personal info as images.
	// Path: /api/v2/logistics/get_booking_shipping_document_data_info
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_data_info?module=95&type=1
	GetBookingShippingDocumentDataInfo(sid uint64, req GetBookingShippingDocumentDataInfoRequest, tok string) (*GetBookingShippingDocumentDataInfoResponse, error)
	// GetBookingShippingDocumentParameter Use this api to get the selectable shipping_document_type and suggested shipping_document_type.
	// Path: /api/v2/logistics/get_booking_shipping_document_parameter
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_parameter?module=95&type=1
	GetBookingShippingDocumentParameter(sid uint64, req GetBookingShippingDocumentParameterRequest, tok string) (*GetBookingShippingDocumentParameterResponse, error)
	// GetBookingShippingDocumentResult Use this api to retrieve the status of the shipping document task. Document will be available for download only after the status change to 'READY'.
	// Path: /api/v2/logistics/get_booking_shipping_document_result
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_result?module=95&type=1
	GetBookingShippingDocumentResult(sid uint64, req GetBookingShippingDocumentResultRequest, tok string) (*GetBookingShippingDocumentResultResponse, error)
	// GetBookingShippingParameter Use this api to get the parameter "info_needed" from the response to check if the booking has pickup or dropoff. This api will also return the addresses and pickup time id options for the pickup method. For dropoff, it can return branch id, sender real name etc, depending on the 3PL requirements.
	// Path: /api/v2/logistics/get_booking_shipping_parameter
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_parameter?module=95&type=1
	GetBookingShippingParameter(sid uint64, opt GetBookingShippingParameterRequest, tok string) (*GetBookingShippingParameterResponse, error)
	// GetBookingTrackingInfo Use this api to get the logistics tracking information of a booking.
	// Path: /api/v2/logistics/get_booking_tracking_info
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_tracking_info?module=95&type=1
	GetBookingTrackingInfo(sid uint64, opt GetBookingTrackingInfoRequest, tok string) (*GetBookingTrackingInfoResponse, error)
	// GetBookingTrackingNumber After arranging shipment (v2.logistics.ship_booking) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
	// Path: /api/v2/logistics/get_booking_tracking_number
	// https://open.shopee.com/documents/v2/v2.logistics.get_booking_tracking_number?module=95&type=1
	GetBookingTrackingNumber(sid uint64, opt GetBookingTrackingNumberRequest, tok string) (*GetBookingTrackingNumberResponse, error)
	// GetChannelList {"content":"<p>Use this api to get all supported logistic channels.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get all supported logistic channels."}]}]}
	// Path: /api/v2/logistics/get_channel_list
	// https://open.shopee.com/documents/v2/v2.logistics.get_channel_list?module=95&type=1
	GetChannelList(sid uint64, tok string) (*LogisticsGetChannelListResponse, error)
	// GetMartPackagingInfo [Only for ID mart seller] The API allows sellers to retrieve their current packaging fee settings.
	// Path: /api/v2/logistics/get_mart_packaging_info
	// https://open.shopee.com/documents/v2/v2.logistics.get_mart_packaging_info?module=95&type=1
	GetMartPackagingInfo(sid uint64, tok string) (*GetMartPackagingInfoResponse, error)
	// GetMassShippingParameter Use this api to check if package support pickup, dropoff, non-integrated. For pickup, return address and pickup time id options. For dropoff, return branch id, sender real name, etc. Can batch request for packages under same product_location_id and logistics_channel_id. [Please call it when packages meet: 1) fulfillment status is LOGISTICS_READY; or 2) fulfillment status is LOGISTICS_PICKUP_RETRY; or 3) fulfillment status is LOGISTICS_REQUEST_CREATED and meet Instant Order Reschedule conditions]
	// Path: /api/v2/logistics/get_mass_shipping_parameter
	// https://open.shopee.com/documents/v2/v2.logistics.get_mass_shipping_parameter?module=95&type=1
	GetMassShippingParameter(sid uint64, req GetMassShippingParameterRequest, tok string) (*GetMassShippingParameterResponse, error)
	// GetMassTrackingNumber After arranging shipment (v2.logistics.mass_ship_order) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
	// Path: /api/v2/logistics/get_mass_tracking_number
	// https://open.shopee.com/documents/v2/v2.logistics.get_mass_tracking_number?module=95&type=1
	GetMassTrackingNumber(sid uint64, req GetMassTrackingNumberRequest, tok string) (*GetMassTrackingNumberResponse, error)
	// GetOperatingHourRestrictions This API is designed to retrieve all restrictions related to inputting operating hours for the v2.logistics.update_operating_hours function. This API ensures that user are aware of any limitations or conditions that may affect their operating hours.
	// Path: /api/v2/logistics/get_operating_hour_restrictions
	// https://open.shopee.com/documents/v2/v2.logistics.get_operating_hour_restrictions?module=95&type=1
	GetOperatingHourRestrictions(sid uint64, tok string) (*GetOperatingHourRestrictionsResponse, error)
	// GetOperatingHours This API is utilized to retrieve the existing operating hours of sellers including Pickup Operating Hours,  Special Hours, Instant Operating Hours, and Shop Collection Operating Hours.
	// Path: /api/v2/logistics/get_operating_hours
	// https://open.shopee.com/documents/v2/v2.logistics.get_operating_hours?module=95&type=1
	GetOperatingHours(sid uint64, tok string) (*GetOperatingHoursResponse, error)
	// GetPauseStatus This API returns the pause status of logistics channels under the shop. Pausing allows the shop to temporarily prevent buyers from placing orders through specific logistics channels. The response includes whether a pause is currently active, the pause end time (if active), and the remaining daily pause quota in seconds (if inactive). Sellers need to refer to the support_pause field in v2.logistics.get_channel_list response to determine which channels are actually paused.
	// Path: /api/v2/logistics/get_pause_status
	// https://open.shopee.com/documents/v2/v2.logistics.get_pause_status?module=95&type=1
	GetPauseStatus(sid uint64, tok string) (*GetPauseStatusResponse, error)
	// GetShippingDocumentDataInfo Use this api to fetch the logistics information of an order, these info can be used for self-design AWB printing. Besides, this api supports returning personal info as images.
	// Path: /api/v2/logistics/get_shipping_document_data_info
	// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_data_info?module=95&type=1
	GetShippingDocumentDataInfo(sid uint64, req GetShippingDocumentDataInfoRequest, tok string) (*GetShippingDocumentDataInfoResponse, error)
	// GetShippingDocumentJobStatus This API retrieves the status of a shipping document job using the job ID provided.
	// Path: /api/v2/logistics/get_shipping_document_job_status
	// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_job_status?module=95&type=1
	GetShippingDocumentJobStatus(sid uint64, req GetShippingDocumentJobStatusRequest, tok string) (*GetShippingDocumentJobStatusResponse, error)
	// GetShippingDocumentParameter Use this api to get the selectable shipping_document_type and suggested shipping_document_type.
	// Path: /api/v2/logistics/get_shipping_document_parameter
	// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_parameter?module=95&type=1
	GetShippingDocumentParameter(sid uint64, req GetShippingDocumentParameterRequest, tok string) (*GetShippingDocumentParameterResponse, error)
	// GetShippingDocumentResult Use this api to retrieve the status of the shipping document task. Document will be available for download only after the status change to 'READY'.
	//
	//
	// Path: /api/v2/logistics/get_shipping_document_result
	// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_result?module=95&type=1
	GetShippingDocumentResult(sid uint64, req GetShippingDocumentResultRequest, tok string) (*GetShippingDocumentResultResponse, error)
	// GetShippingParameter Use this api to check if the package support pickup, dropoff, or non-integrated method. For pickup, will return addresses and pickup time id options. For dropoff, will return branch id, sender real name etc, depending on 3PL requirements. [Please call this API when packages meet: 1) fulfillment status is LOGISTICS_READY; or 2) fulfillment status is LOGISTICS_PICKUP_RETRY; or 3) fulfillment status is LOGISTICS_REQUEST_CREATED and meet Instant Order Reschedule Pickup conditions]
	// Path: /api/v2/logistics/get_shipping_parameter
	// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_parameter?module=95&type=1
	GetShippingParameter(sid uint64, opt GetShippingParameterRequest, tok string) (*GetShippingParameterResponse, error)
	// GetTrackingInfo {"content":"<p>Use this api to get the logistics tracking information of an order.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get the logistics tracking information of an order."}]}]}
	// Path: /api/v2/logistics/get_tracking_info
	// https://open.shopee.com/documents/v2/v2.logistics.get_tracking_info?module=95&type=1
	GetTrackingInfo(sid uint64, opt GetTrackingInfoRequest, tok string) (*GetTrackingInfoResponse, error)
	// GetTrackingNumber After arranging shipment (v2.logistics.ship_order) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
	// Path: /api/v2/logistics/get_tracking_number
	// https://open.shopee.com/documents/v2/v2.logistics.get_tracking_number?module=95&type=1
	GetTrackingNumber(sid uint64, opt GetTrackingNumberRequest, tok string) (*GetTrackingNumberResponse, error)
	// MassShipOrder Use this api to initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_mass_shipping_parameter to fetch all required params first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller. The API can only batch arrange shipment for multiple packages under the same product_location_id and same logistics_channel_id.
	// Path: /api/v2/logistics/mass_ship_order
	// https://open.shopee.com/documents/v2/v2.logistics.mass_ship_order?module=95&type=1
	MassShipOrder(sid uint64, req MassShipOrderRequest, tok string) (*MassShipOrderResponse, error)
	// SetAddressConfig Use this API to set address config of your shop.
	// Path: /api/v2/logistics/set_address_config
	// https://open.shopee.com/documents/v2/v2.logistics.set_address_config?module=95&type=1
	SetAddressConfig(sid uint64, req SetAddressConfigRequest, tok string) (*SetAddressConfigResponse, error)
	// SetMartPackagingInfo [Only for ID mart seller] This API allows sellers to set up their packaging fee info. Through this API, sellers can enable or disable packaging fees, and if enabled, specify the dimensions of the packaging and the associated fee. This ensures that sellers can configure their shipping costs accurately based on their packaging requirements.
	// Path: /api/v2/logistics/set_mart_packaging_info
	// https://open.shopee.com/documents/v2/v2.logistics.set_mart_packaging_info?module=95&type=1
	SetMartPackagingInfo(sid uint64, req SetMartPackagingInfoRequest, tok string) (*SetMartPackagingInfoResponse, error)
	// SetPauseStatus Use this API to set the pause status of logistics channels under the shop. Pausing allows the shop to temporarily prevent buyers from placing orders through specific logistics channels. The response includes whether a pause is currently active, the pause end time (if active), and the remaining daily pause quota in seconds (if inactive). Note: The pause may take a few moments to take effect. Please check for any additional orders that may still be placed during this window.
	// Path: /api/v2/logistics/set_pause_status
	// https://open.shopee.com/documents/v2/v2.logistics.set_pause_status?module=95&type=1
	SetPauseStatus(sid uint64, req SetPauseStatusRequest, tok string) (*SetPauseStatusResponse, error)
	// ShipBooking Use this api to initiate logistics including arrange pickup, dropoff. Should call v2.logistics.get_booking_shipping_parameter to fetch all required param first.
	// Path: /api/v2/logistics/ship_booking
	// https://open.shopee.com/documents/v2/v2.logistics.ship_booking?module=95&type=1
	ShipBooking(sid uint64, req ShipBookingRequest, tok string) (*ShipBookingResponse, error)
	// ShipOrder Use this api to initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed.
	// Path: /api/v2/logistics/ship_order
	// https://open.shopee.com/documents/v2/v2.logistics.ship_order?module=95&type=1
	ShipOrder(sid uint64, req ShipOrderRequest, tok string) (*ShipOrderResponse, error)
	// UpdateAddress Use this API to update the address of a shop.
	// Path: /api/v2/logistics/update_address
	// https://open.shopee.com/documents/v2/v2.logistics.update_address?module=95&type=1
	UpdateAddress(sid uint64, req UpdateAddressRequest, tok string) (*UpdateAddressResponse, error)
	// UpdateChannel Use this api to update shop level logistics channel's configuration.
	// Path: /api/v2/logistics/update_channel
	// https://open.shopee.com/documents/v2/v2.logistics.update_channel?module=95&type=1
	UpdateChannel(sid uint64, req UpdateChannelRequest, tok string) (*UpdateChannelResponse, error)
	// UpdateOperatingHours This API is designed to allow sellers to update their operating hours. It is essential that the values provided in this API align with the restrictions retrieved from the v2.logistics.get_operating_hour_restrictions API to ensure compliance with platform requirements. This API uses overwriting updates, when updating pickup operating hours, still need to include all parts even those not needing changes.
	// Path: /api/v2/logistics/update_operating_hours
	// https://open.shopee.com/documents/v2/v2.logistics.update_operating_hours?module=95&type=1
	UpdateOperatingHours(sid uint64, req UpdateOperatingHoursRequest, tok string) (*UpdateOperatingHoursResponse, error)
	// UpdateSelfCollectionOrderLogistics Use this api to update the order status for buyer to collect the orders directly from your pharmacy. This includes indicating that order is ready for collection, and that the order has been picked up by the buyer. You should call v2.logistics.get_order_detail or v2.logistics.get_package_detail first to get the package_number of such orders.
	// Path: /api/v2/logistics/update_self_collection_order_logistics
	// https://open.shopee.com/documents/v2/v2.logistics.update_self_collection_order_logistics?module=95&type=1
	UpdateSelfCollectionOrderLogistics(sid uint64, req UpdateSelfCollectionOrderLogisticsRequest, tok string) (*UpdateSelfCollectionOrderLogisticsResponse, error)
	// UpdateShippingOrder For pickup method only, use this api to update pickup address and pickup time for packages meet: 1) package's fulfillment status is LOGISTICS_PICKUP_RETRY; or 2) package's fulfillment status is 'LOGISTICS_REQUEST_CREATED' and meets the Instant Order Reschedule Pickup conditions.
	// Path: /api/v2/logistics/update_shipping_order
	// https://open.shopee.com/documents/v2/v2.logistics.update_shipping_order?module=95&type=1
	UpdateShippingOrder(sid uint64, req UpdateShippingOrderRequest, tok string) (*UpdateShippingOrderResponse, error)
	// UpdateTrackingStatus Only available for Brazil sellers. This API is only available for orders/parcels which are fulfilled by BR Seller Logistics channel (logistics_channel_id: 90021), Samsung (logistics_channel_id: 90025) and BR Instant Delivery channel (logistics_channel_id: 90026). The logistics_status will become LOGISTICS_REQUEST_CREATED after arrange shipment, and can call this API to update to: LOGISTICS_PICKUP_DONE, LOGISTICS_DELIVERY_DONE, LOGISTICS_DELIVERY_FAILED.
	// Path: /api/v2/logistics/update_tracking_status
	// https://open.shopee.com/documents/v2/v2.logistics.update_tracking_status?module=95&type=1
	UpdateTrackingStatus(sid uint64, req UpdateTrackingStatusRequest, tok string) (*UpdateTrackingStatusResponse, error)
	// UploadServiceablePolygon Only available for Brazil sellers. Use this API to upload KML file for shop level serviceability setting for BR Entrega Turbo channel (Channel ID: 90026). Please note that multiple Outlet Shops under the same Mart Shop cannot have overlapping service areas.
	// Path: /api/v2/logistics/upload_serviceable_polygon
	// https://open.shopee.com/documents/v2/v2.logistics.upload_serviceable_polygon?module=95&type=1
	UploadServiceablePolygon(sid uint64, filename string, tok string) (*UploadServiceablePolygonResponse, error)
	UploadServiceablePolygonFromReader(sid uint64, filename string, reader io.Reader, tok string) (*UploadServiceablePolygonResponse, error)
}

type LogisticsServiceOp[T any] struct {
	client *Client[T]
}

// BatchShipOrder Use this api to batch initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller.Only channel 90003 - Padrão in Brazil has the permission of this API.
//
// Path: /api/v2/logistics/batch_ship_order
// https://open.shopee.com/documents/v2/v2.logistics.batch_ship_order?module=95&type=1
func (s *LogisticsServiceOp[T]) BatchShipOrder(sid uint64, req BatchShipOrderRequest, tok string) (*BatchShipOrderResponse, error) {
	path := "/logistics/batch_ship_order"
	resp := new(BatchShipOrderResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// BatchUpdateTpfWarehouseTrackingStatus For CB orders that fulfilled by 3PF, support 3PF Warehouse Vendors to update the tpf_tracking_status when 3PF warehouse receive the order and complete the outbound of the package.
// CB orders that fulfilled by 3PF：
// v2.shop.get_shop_info  - shop_fulfillment_flag in {Pure - 3PF Shop,PFF - 3PF Shop,LFF Hybrid Shop}
// And
// v2.order.get_order_detail -  fulfillment_flag = fulfilled_by_local_seller
// Path: /api/v2/logistics/batch_update_tpf_warehouse_tracking_status
// https://open.shopee.com/documents/v2/v2.logistics.batch_update_tpf_warehouse_tracking_status?module=95&type=1
func (s *LogisticsServiceOp[T]) BatchUpdateTpfWarehouseTrackingStatus(sid uint64, req BatchUpdateTpfWarehouseTrackingStatusRequest, tok string) (*BatchUpdateTpfWarehouseTrackingStatusResponse, error) {
	path := "/logistics/batch_update_tpf_warehouse_tracking_status"
	resp := new(BatchUpdateTpfWarehouseTrackingStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// CheckPolygonUpdateStatus Only available for Brazil sellers. Use this API to check the status of polygon file uploaded for BR Entrega Turbo channel (Channel ID: 90026) by querying the task_id returned via the v2.logistics.upload_serviceable_polygon.
// Path: /api/v2/logistics/check_polygon_update_status
// https://open.shopee.com/documents/v2/v2.logistics.check_polygon_update_status?module=95&type=1
func (s *LogisticsServiceOp[T]) CheckPolygonUpdateStatus(sid uint64, req CheckPolygonUpdateStatusRequest, tok string) (*CheckPolygonUpdateStatusResponse, error) {
	path := "/logistics/check_polygon_update_status"
	resp := new(CheckPolygonUpdateStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// CreateBookingShippingDocument Use this api to create shipping document task for each booking and this API is only available after retrieving the tracking number.
// Path: /api/v2/logistics/create_booking_shipping_document
// https://open.shopee.com/documents/v2/v2.logistics.create_booking_shipping_document?module=95&type=1
func (s *LogisticsServiceOp[T]) CreateBookingShippingDocument(sid uint64, req CreateBookingShippingDocumentRequest, tok string) (*CreateBookingShippingDocumentResponse, error) {
	path := "/logistics/create_booking_shipping_document"
	resp := new(CreateBookingShippingDocumentResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// CreateShippingDocument Use this api to create shipping document task for each order or package and this API is only available after retrieving the tracking number.
// Path: /api/v2/logistics/create_shipping_document
// https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document?module=95&type=1
func (s *LogisticsServiceOp[T]) CreateShippingDocument(sid uint64, req CreateShippingDocumentRequest, tok string) (*CreateShippingDocumentResponse, error) {
	path := "/logistics/create_shipping_document"
	resp := new(CreateShippingDocumentResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// CreateShippingDocumentJob This API creates a shipping document job for selected documents. The system receives requests and returns a job ID along with success and failure details.
// Path: /api/v2/logistics/create_shipping_document_job
// https://open.shopee.com/documents/v2/v2.logistics.create_shipping_document_job?module=95&type=1
func (s *LogisticsServiceOp[T]) CreateShippingDocumentJob(sid uint64, req CreateShippingDocumentJobRequest, tok string) (*CreateShippingDocumentJobResponse, error) {
	path := "/logistics/create_shipping_document_job"
	resp := new(CreateShippingDocumentJobResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteAddress Use this api to delete address.
// Path: /api/v2/logistics/delete_address
// https://open.shopee.com/documents/v2/v2.logistics.delete_address?module=95&type=1
func (s *LogisticsServiceOp[T]) DeleteAddress(sid uint64, req DeleteAddressRequest, tok string) (*DeleteAddressResponse, error) {
	path := "/logistics/delete_address"
	resp := new(DeleteAddressResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DeleteSpecialOperatingHour This API is used to delete a specific special operating hour for a seller. This API allows sellers to manage their operating hours by removing any special operating hours that are no longer needed. To use this API, the name of the special operating hour to be deleted should be obtained from the v2.logistics.get_operating_hours API.
// Path: /api/v2/logistics/delete_special_operating_hour
// https://open.shopee.com/documents/v2/v2.logistics.delete_special_operating_hour?module=95&type=1
func (s *LogisticsServiceOp[T]) DeleteSpecialOperatingHour(sid uint64, req DeleteSpecialOperatingHourRequest, tok string) (*DeleteSpecialOperatingHourResponse, error) {
	path := "/logistics/delete_special_operating_hour"
	resp := new(DeleteSpecialOperatingHourResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DownloadBookingShippingDocument Use this api to download shipping_document. You have to call v2.logistics.create_booking_shipping_document to create a new shipping document task first and call v2.logistics.get_booking_shipping_document_result to get the task status second. If the task is READY, you can download this shipping document.
// Path: /api/v2/logistics/download_booking_shipping_document
// https://open.shopee.com/documents/v2/v2.logistics.download_booking_shipping_document?module=95&type=1
func (s *LogisticsServiceOp[T]) DownloadBookingShippingDocument(sid uint64, req DownloadBookingShippingDocumentRequest, tok string) (*DownloadBookingShippingDocumentResponse, error) {
	path := "/logistics/download_booking_shipping_document"
	resp := new(DownloadBookingShippingDocumentResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DownloadShippingDocument Use this api to download shipping_document. You have to call v2.logistics.create_shipping_document to create a new shipping document task first and call v2.logistics.get_shipping_document_resut to get the task status second. If the task is READY, you can download this shipping document.
// Path: /api/v2/logistics/download_shipping_document
// https://open.shopee.com/documents/v2/v2.logistics.download_shipping_document?module=95&type=1
func (s *LogisticsServiceOp[T]) DownloadShippingDocument(sid uint64, req DownloadShippingDocumentRequest, tok string) (*DownloadShippingDocumentResponse, error) {
	path := "/logistics/download_shipping_document"
	resp := new(DownloadShippingDocumentResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DownloadShippingDocumentJob This API allows users to download the shipping document associated with a specific job ID. It checks the job status before proceeding with the download.
// Path: /api/v2/logistics/download_shipping_document_job
// https://open.shopee.com/documents/v2/v2.logistics.download_shipping_document_job?module=95&type=1
func (s *LogisticsServiceOp[T]) DownloadShippingDocumentJob(sid uint64, req DownloadShippingDocumentJobRequest, tok string) (*DownloadShippingDocumentJobResponse, error) {
	path := "/logistics/download_shipping_document_job"
	resp := new(DownloadShippingDocumentJobResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// DownloadToLabel Use the API to download the TO label that should be attached to the carton before drop-off at the warehouse (Only for TW channel_id:30029).
// Path: /api/v2/logistics/download_to_label
// https://open.shopee.com/documents/v2/v2.logistics.download_to_label?module=95&type=1
func (s *LogisticsServiceOp[T]) DownloadToLabel(sid uint64, req DownloadToLabelRequest, tok string) (*DownloadToLabelResponse, error) {
	path := "/logistics/download_to_label"
	resp := new(DownloadToLabelResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetAddressList For integrated logistics channel, use this call to get pickup address for pickup mode order.
// Path: /api/v2/logistics/get_address_list
// https://open.shopee.com/documents/v2/v2.logistics.get_address_list?module=95&type=1
func (s *LogisticsServiceOp[T]) GetAddressList(sid uint64, tok string) (*GetAddressListResponse, error) {
	path := "/logistics/get_address_list"
	resp := new(GetAddressListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetBookingShippingDocumentDataInfo Use this api to fetch the logistics information of a booking these info can be used for airwaybill printing. Dedicated for crossborder SLS order airwaybill. May not be applicable for local channel airwaybill. Besides, this api supports returning personal info as images.
// Path: /api/v2/logistics/get_booking_shipping_document_data_info
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_data_info?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingShippingDocumentDataInfo(sid uint64, req GetBookingShippingDocumentDataInfoRequest, tok string) (*GetBookingShippingDocumentDataInfoResponse, error) {
	path := "/logistics/get_booking_shipping_document_data_info"
	resp := new(GetBookingShippingDocumentDataInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetBookingShippingDocumentParameter Use this api to get the selectable shipping_document_type and suggested shipping_document_type.
// Path: /api/v2/logistics/get_booking_shipping_document_parameter
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_parameter?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingShippingDocumentParameter(sid uint64, req GetBookingShippingDocumentParameterRequest, tok string) (*GetBookingShippingDocumentParameterResponse, error) {
	path := "/logistics/get_booking_shipping_document_parameter"
	resp := new(GetBookingShippingDocumentParameterResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetBookingShippingDocumentResult Use this api to retrieve the status of the shipping document task. Document will be available for download only after the status change to 'READY'.
// Path: /api/v2/logistics/get_booking_shipping_document_result
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_document_result?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingShippingDocumentResult(sid uint64, req GetBookingShippingDocumentResultRequest, tok string) (*GetBookingShippingDocumentResultResponse, error) {
	path := "/logistics/get_booking_shipping_document_result"
	resp := new(GetBookingShippingDocumentResultResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetBookingShippingParameter Use this api to get the parameter "info_needed" from the response to check if the booking has pickup or dropoff. This api will also return the addresses and pickup time id options for the pickup method. For dropoff, it can return branch id, sender real name etc, depending on the 3PL requirements.
// Path: /api/v2/logistics/get_booking_shipping_parameter
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_shipping_parameter?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingShippingParameter(sid uint64, opt GetBookingShippingParameterRequest, tok string) (*GetBookingShippingParameterResponse, error) {
	path := "/logistics/get_booking_shipping_parameter"
	resp := new(GetBookingShippingParameterResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetBookingTrackingInfo Use this api to get the logistics tracking information of a booking.
// Path: /api/v2/logistics/get_booking_tracking_info
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_tracking_info?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingTrackingInfo(sid uint64, opt GetBookingTrackingInfoRequest, tok string) (*GetBookingTrackingInfoResponse, error) {
	path := "/logistics/get_booking_tracking_info"
	resp := new(GetBookingTrackingInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetBookingTrackingNumber After arranging shipment (v2.logistics.ship_booking) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
// Path: /api/v2/logistics/get_booking_tracking_number
// https://open.shopee.com/documents/v2/v2.logistics.get_booking_tracking_number?module=95&type=1
func (s *LogisticsServiceOp[T]) GetBookingTrackingNumber(sid uint64, opt GetBookingTrackingNumberRequest, tok string) (*GetBookingTrackingNumberResponse, error) {
	path := "/logistics/get_booking_tracking_number"
	resp := new(GetBookingTrackingNumberResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetChannelList {"content":"<p>Use this api to get all supported logistic channels.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get all supported logistic channels."}]}]}
// Path: /api/v2/logistics/get_channel_list
// https://open.shopee.com/documents/v2/v2.logistics.get_channel_list?module=95&type=1
func (s *LogisticsServiceOp[T]) GetChannelList(sid uint64, tok string) (*LogisticsGetChannelListResponse, error) {
	path := "/logistics/get_channel_list"
	resp := new(LogisticsGetChannelListResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetMartPackagingInfo [Only for ID mart seller] The API allows sellers to retrieve their current packaging fee settings.
// Path: /api/v2/logistics/get_mart_packaging_info
// https://open.shopee.com/documents/v2/v2.logistics.get_mart_packaging_info?module=95&type=1
func (s *LogisticsServiceOp[T]) GetMartPackagingInfo(sid uint64, tok string) (*GetMartPackagingInfoResponse, error) {
	path := "/logistics/get_mart_packaging_info"
	resp := new(GetMartPackagingInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, nil, resp)
	return resp, err
}

// GetMassShippingParameter Use this api to check if package support pickup, dropoff, non-integrated. For pickup, return address and pickup time id options. For dropoff, return branch id, sender real name, etc. Can batch request for packages under same product_location_id and logistics_channel_id. [Please call it when packages meet: 1) fulfillment status is LOGISTICS_READY; or 2) fulfillment status is LOGISTICS_PICKUP_RETRY; or 3) fulfillment status is LOGISTICS_REQUEST_CREATED and meet Instant Order Reschedule conditions]
// Path: /api/v2/logistics/get_mass_shipping_parameter
// https://open.shopee.com/documents/v2/v2.logistics.get_mass_shipping_parameter?module=95&type=1
func (s *LogisticsServiceOp[T]) GetMassShippingParameter(sid uint64, req GetMassShippingParameterRequest, tok string) (*GetMassShippingParameterResponse, error) {
	path := "/logistics/get_mass_shipping_parameter"
	resp := new(GetMassShippingParameterResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetMassTrackingNumber After arranging shipment (v2.logistics.mass_ship_order) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
// Path: /api/v2/logistics/get_mass_tracking_number
// https://open.shopee.com/documents/v2/v2.logistics.get_mass_tracking_number?module=95&type=1
func (s *LogisticsServiceOp[T]) GetMassTrackingNumber(sid uint64, req GetMassTrackingNumberRequest, tok string) (*GetMassTrackingNumberResponse, error) {
	path := "/logistics/get_mass_tracking_number"
	resp := new(GetMassTrackingNumberResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetOperatingHourRestrictions This API is designed to retrieve all restrictions related to inputting operating hours for the v2.logistics.update_operating_hours function. This API ensures that user are aware of any limitations or conditions that may affect their operating hours.
// Path: /api/v2/logistics/get_operating_hour_restrictions
// https://open.shopee.com/documents/v2/v2.logistics.get_operating_hour_restrictions?module=95&type=1
func (s *LogisticsServiceOp[T]) GetOperatingHourRestrictions(sid uint64, tok string) (*GetOperatingHourRestrictionsResponse, error) {
	path := "/logistics/get_operating_hour_restrictions"
	resp := new(GetOperatingHourRestrictionsResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetOperatingHours This API is utilized to retrieve the existing operating hours of sellers including Pickup Operating Hours,  Special Hours, Instant Operating Hours, and Shop Collection Operating Hours.
// Path: /api/v2/logistics/get_operating_hours
// https://open.shopee.com/documents/v2/v2.logistics.get_operating_hours?module=95&type=1
func (s *LogisticsServiceOp[T]) GetOperatingHours(sid uint64, tok string) (*GetOperatingHoursResponse, error) {
	path := "/logistics/get_operating_hours"
	resp := new(GetOperatingHoursResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetPauseStatus This API returns the pause status of logistics channels under the shop. Pausing allows the shop to temporarily prevent buyers from placing orders through specific logistics channels. The response includes whether a pause is currently active, the pause end time (if active), and the remaining daily pause quota in seconds (if inactive). Sellers need to refer to the support_pause field in v2.logistics.get_channel_list response to determine which channels are actually paused.
// Path: /api/v2/logistics/get_pause_status
// https://open.shopee.com/documents/v2/v2.logistics.get_pause_status?module=95&type=1
func (s *LogisticsServiceOp[T]) GetPauseStatus(sid uint64, tok string) (*GetPauseStatusResponse, error) {
	path := "/logistics/get_pause_status"
	resp := new(GetPauseStatusResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, nil)
	return resp, err
}

// GetShippingDocumentDataInfo Use this api to fetch the logistics information of an order, these info can be used for self-design AWB printing. Besides, this api supports returning personal info as images.
// Path: /api/v2/logistics/get_shipping_document_data_info
// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_data_info?module=95&type=1
func (s *LogisticsServiceOp[T]) GetShippingDocumentDataInfo(sid uint64, req GetShippingDocumentDataInfoRequest, tok string) (*GetShippingDocumentDataInfoResponse, error) {
	path := "/logistics/get_shipping_document_data_info"
	resp := new(GetShippingDocumentDataInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetShippingDocumentJobStatus This API retrieves the status of a shipping document job using the job ID provided.
// Path: /api/v2/logistics/get_shipping_document_job_status
// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_job_status?module=95&type=1
func (s *LogisticsServiceOp[T]) GetShippingDocumentJobStatus(sid uint64, req GetShippingDocumentJobStatusRequest, tok string) (*GetShippingDocumentJobStatusResponse, error) {
	path := "/logistics/get_shipping_document_job_status"
	resp := new(GetShippingDocumentJobStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetShippingDocumentParameter Use this api to get the selectable shipping_document_type and suggested shipping_document_type.
// Path: /api/v2/logistics/get_shipping_document_parameter
// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_parameter?module=95&type=1
func (s *LogisticsServiceOp[T]) GetShippingDocumentParameter(sid uint64, req GetShippingDocumentParameterRequest, tok string) (*GetShippingDocumentParameterResponse, error) {
	path := "/logistics/get_shipping_document_parameter"
	resp := new(GetShippingDocumentParameterResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetShippingDocumentResult Use this api to retrieve the status of the shipping document task. Document will be available for download only after the status change to 'READY'.
//
// Path: /api/v2/logistics/get_shipping_document_result
// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_document_result?module=95&type=1
func (s *LogisticsServiceOp[T]) GetShippingDocumentResult(sid uint64, req GetShippingDocumentResultRequest, tok string) (*GetShippingDocumentResultResponse, error) {
	path := "/logistics/get_shipping_document_result"
	resp := new(GetShippingDocumentResultResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// GetShippingParameter Use this api to check if the package support pickup, dropoff, or non-integrated method. For pickup, will return addresses and pickup time id options. For dropoff, will return branch id, sender real name etc, depending on 3PL requirements. [Please call this API when packages meet: 1) fulfillment status is LOGISTICS_READY; or 2) fulfillment status is LOGISTICS_PICKUP_RETRY; or 3) fulfillment status is LOGISTICS_REQUEST_CREATED and meet Instant Order Reschedule Pickup conditions]
// Path: /api/v2/logistics/get_shipping_parameter
// https://open.shopee.com/documents/v2/v2.logistics.get_shipping_parameter?module=95&type=1
func (s *LogisticsServiceOp[T]) GetShippingParameter(sid uint64, opt GetShippingParameterRequest, tok string) (*GetShippingParameterResponse, error) {
	path := "/logistics/get_shipping_parameter"
	resp := new(GetShippingParameterResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetTrackingInfo {"content":"<p>Use this api to get the logistics tracking information of an order.</p>","raw_content":[{"name":"paragraph","children":[{"data":"Use this api to get the logistics tracking information of an order."}]}]}
// Path: /api/v2/logistics/get_tracking_info
// https://open.shopee.com/documents/v2/v2.logistics.get_tracking_info?module=95&type=1
func (s *LogisticsServiceOp[T]) GetTrackingInfo(sid uint64, opt GetTrackingInfoRequest, tok string) (*GetTrackingInfoResponse, error) {
	path := "/logistics/get_tracking_info"
	resp := new(GetTrackingInfoResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// GetTrackingNumber After arranging shipment (v2.logistics.ship_order) for the integrated channel, use this api to get the tracking_number, which is a required parameter for creating shipping labels. The api response can return tracking_number empty, since this info is dependent from the 3PL, due to this it is allowed to keep calling the api within 5 minutes interval, until the tracking_number is returned.
// Path: /api/v2/logistics/get_tracking_number
// https://open.shopee.com/documents/v2/v2.logistics.get_tracking_number?module=95&type=1
func (s *LogisticsServiceOp[T]) GetTrackingNumber(sid uint64, opt GetTrackingNumberRequest, tok string) (*GetTrackingNumberResponse, error) {
	path := "/logistics/get_tracking_number"
	resp := new(GetTrackingNumberResponse)
	err := s.client.WithShop(sid, tok).Get(path, resp, opt)
	return resp, err
}

// MassShipOrder Use this api to initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_mass_shipping_parameter to fetch all required params first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller. The API can only batch arrange shipment for multiple packages under the same product_location_id and same logistics_channel_id.
// Path: /api/v2/logistics/mass_ship_order
// https://open.shopee.com/documents/v2/v2.logistics.mass_ship_order?module=95&type=1
func (s *LogisticsServiceOp[T]) MassShipOrder(sid uint64, req MassShipOrderRequest, tok string) (*MassShipOrderResponse, error) {
	path := "/logistics/mass_ship_order"
	resp := new(MassShipOrderResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// SetAddressConfig Use this API to set address config of your shop.
// Path: /api/v2/logistics/set_address_config
// https://open.shopee.com/documents/v2/v2.logistics.set_address_config?module=95&type=1
func (s *LogisticsServiceOp[T]) SetAddressConfig(sid uint64, req SetAddressConfigRequest, tok string) (*SetAddressConfigResponse, error) {
	path := "/logistics/set_address_config"
	resp := new(SetAddressConfigResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// SetMartPackagingInfo [Only for ID mart seller] This API allows sellers to set up their packaging fee info. Through this API, sellers can enable or disable packaging fees, and if enabled, specify the dimensions of the packaging and the associated fee. This ensures that sellers can configure their shipping costs accurately based on their packaging requirements.
// Path: /api/v2/logistics/set_mart_packaging_info
// https://open.shopee.com/documents/v2/v2.logistics.set_mart_packaging_info?module=95&type=1
func (s *LogisticsServiceOp[T]) SetMartPackagingInfo(sid uint64, req SetMartPackagingInfoRequest, tok string) (*SetMartPackagingInfoResponse, error) {
	path := "/logistics/set_mart_packaging_info"
	resp := new(SetMartPackagingInfoResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// SetPauseStatus Use this API to set the pause status of logistics channels under the shop. Pausing allows the shop to temporarily prevent buyers from placing orders through specific logistics channels. The response includes whether a pause is currently active, the pause end time (if active), and the remaining daily pause quota in seconds (if inactive). Note: The pause may take a few moments to take effect. Please check for any additional orders that may still be placed during this window.
// Path: /api/v2/logistics/set_pause_status
// https://open.shopee.com/documents/v2/v2.logistics.set_pause_status?module=95&type=1
func (s *LogisticsServiceOp[T]) SetPauseStatus(sid uint64, req SetPauseStatusRequest, tok string) (*SetPauseStatusResponse, error) {
	path := "/logistics/set_pause_status"
	resp := new(SetPauseStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// ShipBooking Use this api to initiate logistics including arrange pickup, dropoff. Should call v2.logistics.get_booking_shipping_parameter to fetch all required param first.
// Path: /api/v2/logistics/ship_booking
// https://open.shopee.com/documents/v2/v2.logistics.ship_booking?module=95&type=1
func (s *LogisticsServiceOp[T]) ShipBooking(sid uint64, req ShipBookingRequest, tok string) (*ShipBookingResponse, error) {
	path := "/logistics/ship_booking"
	resp := new(ShipBookingResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// ShipOrder Use this api to initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed.
// Path: /api/v2/logistics/ship_order
// https://open.shopee.com/documents/v2/v2.logistics.ship_order?module=95&type=1
func (s *LogisticsServiceOp[T]) ShipOrder(sid uint64, req ShipOrderRequest, tok string) (*ShipOrderResponse, error) {
	path := "/logistics/ship_order"
	resp := new(ShipOrderResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateAddress Use this API to update the address of a shop.
// Path: /api/v2/logistics/update_address
// https://open.shopee.com/documents/v2/v2.logistics.update_address?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateAddress(sid uint64, req UpdateAddressRequest, tok string) (*UpdateAddressResponse, error) {
	path := "/logistics/update_address"
	resp := new(UpdateAddressResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateChannel Use this api to update shop level logistics channel's configuration.
// Path: /api/v2/logistics/update_channel
// https://open.shopee.com/documents/v2/v2.logistics.update_channel?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateChannel(sid uint64, req UpdateChannelRequest, tok string) (*UpdateChannelResponse, error) {
	path := "/logistics/update_channel"
	resp := new(UpdateChannelResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateOperatingHours This API is designed to allow sellers to update their operating hours. It is essential that the values provided in this API align with the restrictions retrieved from the v2.logistics.get_operating_hour_restrictions API to ensure compliance with platform requirements. This API uses overwriting updates, when updating pickup operating hours, still need to include all parts even those not needing changes.
// Path: /api/v2/logistics/update_operating_hours
// https://open.shopee.com/documents/v2/v2.logistics.update_operating_hours?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateOperatingHours(sid uint64, req UpdateOperatingHoursRequest, tok string) (*UpdateOperatingHoursResponse, error) {
	path := "/logistics/update_operating_hours"
	resp := new(UpdateOperatingHoursResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateSelfCollectionOrderLogistics Use this api to update the order status for buyer to collect the orders directly from your pharmacy. This includes indicating that order is ready for collection, and that the order has been picked up by the buyer. You should call v2.logistics.get_order_detail or v2.logistics.get_package_detail first to get the package_number of such orders.
// Path: /api/v2/logistics/update_self_collection_order_logistics
// https://open.shopee.com/documents/v2/v2.logistics.update_self_collection_order_logistics?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateSelfCollectionOrderLogistics(sid uint64, req UpdateSelfCollectionOrderLogisticsRequest, tok string) (*UpdateSelfCollectionOrderLogisticsResponse, error) {
	path := "/logistics/update_self_collection_order_logistics"
	resp := new(UpdateSelfCollectionOrderLogisticsResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateShippingOrder For pickup method only, use this api to update pickup address and pickup time for packages meet: 1) package's fulfillment status is LOGISTICS_PICKUP_RETRY; or 2) package's fulfillment status is 'LOGISTICS_REQUEST_CREATED' and meets the Instant Order Reschedule Pickup conditions.
// Path: /api/v2/logistics/update_shipping_order
// https://open.shopee.com/documents/v2/v2.logistics.update_shipping_order?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateShippingOrder(sid uint64, req UpdateShippingOrderRequest, tok string) (*UpdateShippingOrderResponse, error) {
	path := "/logistics/update_shipping_order"
	resp := new(UpdateShippingOrderResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UpdateTrackingStatus Only available for Brazil sellers. This API is only available for orders/parcels which are fulfilled by BR Seller Logistics channel (logistics_channel_id: 90021), Samsung (logistics_channel_id: 90025) and BR Instant Delivery channel (logistics_channel_id: 90026). The logistics_status will become LOGISTICS_REQUEST_CREATED after arrange shipment, and can call this API to update to: LOGISTICS_PICKUP_DONE, LOGISTICS_DELIVERY_DONE, LOGISTICS_DELIVERY_FAILED.
// Path: /api/v2/logistics/update_tracking_status
// https://open.shopee.com/documents/v2/v2.logistics.update_tracking_status?module=95&type=1
func (s *LogisticsServiceOp[T]) UpdateTrackingStatus(sid uint64, req UpdateTrackingStatusRequest, tok string) (*UpdateTrackingStatusResponse, error) {
	path := "/logistics/update_tracking_status"
	resp := new(UpdateTrackingStatusResponse)
	err := s.client.WithShop(sid, tok).Post(path, req, resp)
	return resp, err
}

// UploadServiceablePolygon Only available for Brazil sellers. Use this API to upload KML file for shop level serviceability setting for BR Entrega Turbo channel (Channel ID: 90026). Please note that multiple Outlet Shops under the same Mart Shop cannot have overlapping service areas.
// Path: /api/v2/logistics/upload_serviceable_polygon
// https://open.shopee.com/documents/v2/v2.logistics.upload_serviceable_polygon?module=95&type=1
func (s *LogisticsServiceOp[T]) UploadServiceablePolygon(sid uint64, filename string, tok string) (*UploadServiceablePolygonResponse, error) {
	path := "/logistics/upload_serviceable_polygon"
	resp := new(UploadServiceablePolygonResponse)
	err := s.client.WithShop(sid, tok).Upload(path, "image", filename, resp)
	return resp, err
}

func (s *LogisticsServiceOp[T]) UploadServiceablePolygonFromReader(sid uint64, filename string, reader io.Reader, tok string) (*UploadServiceablePolygonResponse, error) {
	path := "/logistics/upload_serviceable_polygon"
	resp := new(UploadServiceablePolygonResponse)
	err := s.client.WithShop(sid, tok).UploadFromReader(path, "image", filename, reader, resp)
	return resp, err
}
