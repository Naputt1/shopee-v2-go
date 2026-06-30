package goshopee

type AcceptOfferRequest struct {
	ReturnSn string `json:"return_sn"` // [Required] The serial number of return.
}
type AcceptOfferResponse struct {
	BaseResponse                         // Common response fields
	Response     AcceptOfferResponseData `json:"response"` //
}
type AcceptOfferResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
}
type Activity struct {
	ActivityId      int64           `json:"activity_id"`      // [Required] The id of activity.
	ActivityType    string          `json:"activity_type"`    // [Required] The type of activity.
	OriginalPrice   string          `json:"original_price"`   // [Required] activity's origin price
	DiscountedPrice string          `json:"discounted_price"` // [Required] activity's discount price
	Items           []ActivityItems `json:"items"`            // [Required]
	RefundAmount    float64         `json:"refund_amount"`    // [Required] <p>item's refund amount for bundle deal cases, only for shops whitelisted for Partial Qty RR.</p>
}
type ActivityItems struct {
	ItemId            int64  `json:"item_id"`            // [Required] The id of item.
	VariationId       int64  `json:"variation_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	QuantityPurchased int64  `json:"quantity_purchased"` // [Required] item's quantity purchase
	OriginalPrice     string `json:"original_price"`     // [Required] item's origin price
}
type BuyerVideos struct {
	ThumbnailUrl string `json:"thumbnail_url"` // [Required] <p>The thumbnail url of video</p>
	VideoUrl     string `json:"video_url"`     // [Required] <p>The url of video</p>
}
type CancelDisputeRequest struct {
	Email    string `json:"email"`     // [Required] <p>The operator's email address. For operation record keeping purposes (same as v2.returns.dispute API).</p>
	ReturnSn string `json:"return_sn"` // [Required] <p>Shopee's unique serial number identifier for a Return Refund request.</p><p><br /></p><p>Note:&nbsp;Sellers can only cancel compensation disputes, not normal disputes. This means that sellers can only cancel disputes <b>when the return_status is ACCEPTED and the compensation_status is COMPENSATION_REQUESTED</b>.</p>
}
type CancelDisputeResponse struct {
	BaseResponse                           // Common response fields
	Response     CancelDisputeResponseData `json:"response"` //
}
type CancelDisputeResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>Shopee's unique serial number identifier for a Return Refund request.</p>
	Message  string `json:"message"`   // [Required] <p>To indicate whether the cancel dispute operation is successful or failed.</p>
}
type ConfirmRequest struct {
	ReturnSn string `json:"return_sn"` // [Required] The serial number of return.
}
type ConfirmResponse struct {
	BaseResponse                     // Common response fields
	Response     ConfirmResponseData `json:"response"` //
}
type ConfirmResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] The identifier for an API request for error tracking
}
type ConvertImageRequest struct {
	ReturnSn    string      `json:"return_sn"`    // [Required] <p>The serial number of return.<br /></p>
	UploadImage interface{} `json:"upload_image"` // [Required] <p>The proof picture to be uploaded must be within 10MB in size, and the format only supports .jpg, .jpeg, and .png. Only one picture is allowed to be uploaded per request. If multiple pictures are uploaded, only the first picture will be processed.<br /></p>
}
type ConvertImageResponse struct {
	BaseResponse                          // Common response fields
	Response     ConvertImageResponseData `json:"response"` //
}
type ConvertImageResponseData struct {
	Url       string `json:"url"`       // [Required] <p>The link uploaded to the image server can be used with the upload_proof interface.</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The image thumbnail.</p>
}
type DisputeReason struct {
	DisputeReason      string           `json:"dispute_reason"`       // [Required] <p>The dispute_reason for one specific case. See Data Definition - DisputeReason.<br /></p>
	DisputeRequirement string           `json:"dispute_requirement"`  // [Required] <p>Indicate the importance of uploading required proof.<br /></p>
	SampleEvidence     []SampleEvidence `json:"sample_evidence"`      // [Required] <p>The URL of sample evidence to upload.<br /></p>
	EvidenceModuleList []EvidenceModule `json:"evidence_module_list"` // [Required] <p>The associated evidence module list for current dispute reason.</p>
}
type DisputeRequest struct {
	DisputeReasonId   int64                 `json:"dispute_reason_id"`             // [Required] <p>The dispute reason id.<br /></p><p>Please call&nbsp;v2.returns.get_return_dispute_reason to get it.</p>
	DisputeTextReason *string               `json:"dispute_text_reason,omitempty"` // [Optional] <p>The content of dispute reason.</p>
	Email             string                `json:"email"`                         // [Required] <p>The email address.</p>
	ImageList         []DisputeRequestImage `json:"image_list,omitempty"`          // [Optional] <p>Determines whether image submission is mandatory for the dispute request - mandatory input field for all dispute reasons except "Did not receive the return product".</p>
	ReturnSn          string                `json:"return_sn"`                     // [Required] The serial number of return.
}
type DisputeRequestImage struct {
	ModuleIndex int64    `json:"module_index"` // [Required] <p>The module_index of current evidence module returned by get_return_dispute_reason API.</p>
	Requirement string   `json:"requirement"`  // [Required] <p>The requirement content of current evidence module returned by get_return_dispute_reason API.</p>
	ImageUrl    []string `json:"image_url"`    // [Required] <p>The image URLs of current evidence module. It is recommended to pass in the URL returned by convert_image API.<br /></p>
}
type DisputeResponse struct {
	BaseResponse                     // Common response fields
	Response     DisputeResponseData `json:"response"` //
}
type DisputeResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
	Msg      string `json:"msg"`       // [Required]
}
type EvidenceModule struct {
	ModuleIndex int64  `json:"module_index"` // [Required] <p>The index of current evidence module.</p>
	Requirement string `json:"requirement"`  // [Required] <p>The specific requirement of current evidence module.</p>
	IsRequired  bool   `json:"is_required"`  // [Required] <p>Indicate if current evidence module is mandatory or not.<br /></p>
}
type FollowUpAction struct {
	ItemId               int64    `json:"item_id"`                 // [Required] <p>Unique identifier of the item.</p>
	ModelId              int64    `json:"model_id"`                // [Required] <p>Unique identifier of the model under the item.</p>
	Qty                  int64    `json:"qty"`                     // [Required] <p>Quantity of items or models under the same current status.</p>
	CurrentStatus        int64    `json:"current_status"`          // [Required] <p>Current status for the item/model within the warehouse.</p><p><br /></p><p>Applicable values:</p><p>1：Dispose</p><p>2：Return to Seller</p><p>7：Received and Putaway</p><p>8：Return to Buyer</p><p>9：Shortage</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, the following values will not be returned for now, and will be returned once Resell becomes applicable to Return Refund parcels in the future:</p><p>3：Putaway for Resell</p><p>4：Resell Outbound</p><p>5：Resell Failed</p><p>6：Resell Exit</p>
	RelatedOrderSnList   []string `json:"related_order_sn_list"`   // [Required] <p>List of order_sn generated from the Resell process. Returned only when current_status = 4 (Resell Outbound).</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, this field will remain empty for now, and valid values will be returned once Resell becomes applicable to Return Refund parcels in the future.</p>
	ResellFailedNextStep string   `json:"resell_failed_next_step"` // [Required] <p>Next step after a Resell failure. Returned only when current_status = 5 (Resell Failed).</p><p><br /></p><p>Note: Since Resell is currently applicable only to Failed Delivery parcels, this field will remain empty for now, and valid values will be returned once Resell becomes applicable to Return Refund parcels in the future.</p>
}
type GetAvailableSolutionsRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] The serial number of return.
}
type GetAvailableSolutionsResponse struct {
	BaseResponse                                   // Common response fields
	Response     GetAvailableSolutionsResponseData `json:"response"` //
}
type GetAvailableSolutionsResponseData struct {
	ReturnSn          string             `json:"return_sn"`           // [Required] <p>The serial number of return.<br /></p>
	OfferReturnRefund *OfferReturnRefund `json:"offer_return_refund"` // [Required]
	OfferRefund       *OfferReturnRefund `json:"offer_refund"`        // [Required]
}
type GetReturnDetailRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] The serial number of return.
}
type GetReturnDetailResponse struct {
	BaseResponse                             // Common response fields
	Response     GetReturnDetailResponseData `json:"response"` // Amount of the refund.
}
type GetReturnDetailResponseData struct {
	Image                               []string                          `json:"image"`                                   // [Required] Image URLs of return.
	BuyerVideos                         []BuyerVideos                     `json:"buyer_videos"`                            // [Required]
	Reason                              string                            `json:"reason"`                                  // [Required] <p>Indicates the original return reason submitted by the buyer when initiating the return request.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason.</p><p><br /></p><p> </p><p>Note: There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;If the platform updates the return reason during this process, the reassessed outcome will be provided separately in the <b>reassessed_request_reason</b> field.</p>
	TextReason                          string                            `json:"text_reason"`                             // [Required] Reason that buyer provide.
	ReassessedRequestReason             string                            `json:"reassessed_request_reason"`               // [Required] <p>Indicates the return reason reassessed by the platform as more suitable.</p><p><br /></p><p>There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason. If no reassessment has been made, the value will be NONE.</p>
	ReturnSn                            string                            `json:"return_sn"`                               // [Required] The serial number of return.
	RefundAmount                        float64                           `json:"refund_amount"`                           // [Required] Amount of the refund.
	Currency                            string                            `json:"currency"`                                // [Required] Currency of the return.
	CreateTime                          int64                             `json:"create_time"`                             // [Required] The time of return create.
	UpdateTime                          int64                             `json:"update_time"`                             // [Required] The time of modify return.
	Status                              string                            `json:"status"`                                  // [Required] Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	DueDate                             int64                             `json:"due_date"`                                // [Required] The last time seller deal with this return.
	TrackingNumber                      string                            `json:"tracking_number"`                         // [Required] The tracking number assigned by the shipping carrier for item shipment.
	DisputeReason                       int64                             `json:"dispute_reason"`                          // [Required] The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeTextReason                   string                            `json:"dispute_text_reason"`                     // [Required] The reason that seller provide. While the return has been disputed, this field is useful.
	NeedsLogistics                      bool                              `json:"needs_logistics"`                         // [Required] Items to be sent back to seller. Can be either integrated/non-integrated.
	AmountBeforeDiscount                float64                           `json:"amount_before_discount"`                  // [Required] Order price before discount.
	User                                *User                             `json:"user"`                                    // [Required]
	Item                                []GetReturnDetailResponseDataItem `json:"item"`                                    // [Required]
	OrderSn                             string                            `json:"order_sn"`                                // [Required] Shopee's unique identifier for an order.
	ReturnShipDueDate                   int64                             `json:"return_ship_due_date"`                    // [Required] The due date for buyer to ship order.
	ReturnSellerDueDate                 int64                             `json:"return_seller_due_date"`                  // [Required] The due date for seller to deal with this return when buyer have shipped order.
	Activity                            []Activity                        `json:"activity"`                                // [Required]
	SellerProof                         *SellerProof                      `json:"seller_proof"`                            // [Required]
	SellerCompensation                  *SellerCompensation               `json:"seller_compensation"`                     // [Required]
	Negotiation                         *Negotiation                      `json:"negotiation"`                             // [Required]
	LogisticsStatus                     LogisticsStatus                   `json:"logistics_status"`                        // [Required] <p>To indicate the reverse logistics status. See "Data Definition - LogisticsStatus".</p><p><br /></p><p>Note:&nbsp;</p><p>- This is a legacy field that only reflects the reverse logistics status of Normal RR. To determine whether the RR is a Normal RR, check if return_refund_request_type = 0.</p><p>- If you need the reverse logistics status for Normal RR, In-transit RR, or Return-on-the-Spot, please use the newly released field reverse_logistic_status instead.</p>
	ReverseLogisticStatus               string                            `json:"reverse_logistic_status"`                 // [Required] <p>To indicate the latest reverse logistic status of a return, referring to the current status of the buyer shipping the return parcel back to the validation point (seller or warehouse),&nbsp;including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>See "Data Definition - ReverseLogisticsStatus" as status displayed for Normal RR and In-transit RR or Return-on-the-Spot are different.</p>
	ReturnPickupAddress                 *ReturnPickupAddress              `json:"return_pickup_address"`                   // [Required] To indicate the buyer's pickup address
	VirtualContactNumber                string                            `json:"virtual_contact_number"`                  // [Required] <p>[Only for TW non-integrated channel] The virtual phone number to contact the recipient.</p>
	PackageQueryNumber                  string                            `json:"package_query_number"`                    // [Required] <p>[Only for TW non-integrated channel] The query number used in virtual phone number calls to contact the recipient of this return.</p>
	ReturnAddress                       *ReturnAddress                    `json:"return_address"`                          // [Required]
	ReturnRefundType                    string                            `json:"return_refund_type"`                      // [Required] <p>To indicate whether the return is RRBOC (Return/Refund request raised before Order Complete) or RRAOC&nbsp;(Return/Refund request raised after Order Complete).</p>
	ReturnSolution                      int64                             `json:"return_solution"`                         // [Required] <p>To indicate the most updated solution of the Return/Refund request (NOTE: this is not the solution during negotiation). Applicable value:&nbsp;</p><p>- 0: Return and Refund</p><p>- 1: Refund Only<br /></p>
	IsSellerArrange                     bool                              `json:"is_seller_arrange"`                       // [Required] <p>To indicate whether the return_sn is using the “Seller Arrange” return method. This would only be True for TW and BR.</p>
	IsShippingProofMandatory            bool                              `json:"is_shipping_proof_mandatory"`             // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	HasUploadedShippingProof            bool                              `json:"has_uploaded_shipping_proof"`             // [Required] <p>To indicate whether seller has already uploaded shipping proof for this return.</p>
	IsReverseLogisticsChannelIntegrated bool                              `json:"is_reverse_logistics_channel_integrated"` // [Required] <p>To indicate whether the reverse logistic channel type selected is integrated or non-integrated.</p>
	ReverseLogisticChannelName          string                            `json:"reverse_logistic_channel_name"`           // [Required] <p>To indicate reverse logistic carrier name.</p>
	ReturnRefundRequestType             int64                             `json:"return_refund_request_type"`              // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType                      string                            `json:"validation_type"`                         // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	IsArrivedAtWarehouse                int64                             `json:"is_arrived_at_warehouse"`                 // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Indicates the parcel’s check-in status at the warehouse. This field helps sellers quickly determine whether the parcel has arrived at the warehouse or has been rejected.&nbsp;</p><p><br /></p><p>Applicable values:</p><p>1: Pending Inbound</p><p>2: Rejected</p><p>3: Inbound</p><p>4: Cancelled</p>
	FollowUpActionList                  []FollowUpAction                  `json:"follow_up_action_list"`                   // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Warehouse handling actions for each item in the parcel.</p>
}
type GetReturnDetailResponseDataItem struct {
	ModelId      int64    `json:"model_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	Name         string   `json:"name"`           // [Required] Name of item in local language.
	Images       []string `json:"images"`         // [Required] Image URLs of item.
	Amount       int64    `json:"amount"`         // [Required] Amount of this item.
	ItemPrice    float64  `json:"item_price"`     // [Required] The price of item.
	IsAddOnDeal  bool     `json:"is_add_on_deal"` // [Required] To indicate if this item belongs to an addon deal.
	IsMainItem   bool     `json:"is_main_item"`   // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId  int64    `json:"add_on_deal_id"` // [Required] The unique identity of an addon deal.
	ItemId       int64    `json:"item_id"`        // [Required] The id of item.
	ItemSku      string   `json:"item_sku"`       // [Required] The sku of item.
	VariationSku string   `json:"variation_sku"`  // [Required] the variation sku of item
	RefundAmount float64  `json:"refund_amount"`  // [Required] <p>item's refund amount.&nbsp;only for shops whitelisted for Partial Qty RR.</p><p>If not available, refer to item_price</p>
}
type GetReturnDisputeReasonRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.</p>
}
type GetReturnDisputeReasonResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetReturnDisputeReasonResponseData `json:"response"` // <p>Detail informations you are querying.<br /></p>
}
type GetReturnDisputeReasonResponseData struct {
	DisputeReasonList []DisputeReason `json:"dispute_reason_list"` // [Required] <p>The dispute_reason and associated evidence list.<br /></p>
}
type GetReturnListRequest struct {
	CreateTimeFrom           *int64  `json:"create_time_from,omitempty" url:"create_time_from,omitempty"`                     // [Optional] The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo             *int64  `json:"create_time_to,omitempty" url:"create_time_to,omitempty"`                         // [Optional] The create_time_from and create_time_to fields specify a date range for retrieving orders (based on the order create time). The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	NegotiationStatus        *string `json:"negotiation_status,omitempty" url:"negotiation_status,omitempty"`                 // [Optional] This is for filtering return request by counter status. See "Data Definition - NegotiationStatus"
	PageNo                   int64   `json:"page_no" url:"page_no"`                                                           // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize                 int64   `json:"page_size" url:"page_size"`                                                       // [Required] if many items are available to retrieve, you may need to call GetReturnList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	SellerCompensationStatus *string `json:"seller_compensation_status,omitempty" url:"seller_compensation_status,omitempty"` // [Optional] This is for filtering return request by compensation status. See "Data Definition - SellerCompensationStatus"
	SellerProofStatus        *string `json:"seller_proof_status,omitempty" url:"seller_proof_status,omitempty"`               // [Optional] <p>This is for filtering return request by proof status. See "Data Definition - SellerProofStatus"</p>
	Status                   *string `json:"status,omitempty" url:"status,omitempty"`                                         // [Optional] This is for filtering return request by return status. See "Data Definition - ReturnStatus"
	UpdateTimeFrom           *int64  `json:"update_time_from,omitempty" url:"update_time_from,omitempty"`                     // [Optional] <p>The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the last return updated time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. update_time_from should be &gt;= create_time_from<br /></p>
	UpdateTimeTo             *int64  `json:"update_time_to,omitempty" url:"update_time_to,omitempty"`                         // [Optional] <p>The update_time_from and update_time_to fields specify a date range for retrieving orders (based on the last return updated time). The update_time_from field is the starting date range. The maximum date range that may be specified with the update_time_from and update_time_to fields is 15 days. update_time_from should be &gt;= create_time_from<br /></p>
}
type GetReturnListResponse struct {
	BaseResponse                           // Common response fields
	Response     GetReturnListResponseData `json:"response"` // Amount of the refund.
}
type GetReturnListResponseData struct {
	More   bool     `json:"more"`   // [Required] Whether has next page
	Return []Return `json:"return"` // [Required]
}
type GetReverseTrackingInfoRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>Shopee's unique identifier for a return/refund request (serial number of return).</p>
}
type GetReverseTrackingInfoResponse struct {
	BaseResponse                                    // Common response fields
	Response     GetReverseTrackingInfoResponseData `json:"response"` //
}
type GetReverseTrackingInfoResponseData struct {
	ReturnSn                        string                                           `json:"return_sn"`                           // [Required] <p>Shopee's unique identifier for a return/refund request (serial number of return).<br /></p>
	ReturnRefundRequestType         int64                                            `json:"return_refund_request_type"`          // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType                  string                                           `json:"validation_type"`                     // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	ReverseLogisticsStatus          string                                           `json:"reverse_logistics_status"`            // [Required] <p>To indicate the latest reverse logistic status of a return, referring to the current status of the buyer shipping the return parcel back to the validation point (seller or warehouse),&nbsp;including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>See "Data Definition - ReverseLogisticsStatus" as status displayed for Normal RR and In-transit RR or Return-on-the-Spot are different.</p><p><br /></p><p>Note: If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller).</b> Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	ReverseLogisticsUpdateTime      int64                                            `json:"reverse_logistics_update_time"`       // [Required] <p>The last update time of the reverse logistics status including Normal RR, In-transit RR, and Return-on-the-Spot.</p><p><br /></p><p>Note: If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller).</b> Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	EstimatedDeliveryDateMax        int64                                            `json:"estimated_delivery_date_max"`         // [Required] <p>The maximum estimated delivery date for the reverse logistics. This is calculated by Shopee Logistics Services once buyer ships out if there is historical tracking data available from third party logistics provider.&nbsp;</p><p><br /></p><p>Note: Only available for <b>Normal RR with integrated reverse logistics.</b></p>
	EstimatedDeliveryDateMin        int64                                            `json:"estimated_delivery_date_min"`         // [Required] <p>The minimum estimated delivery date for the reverse logistics.&nbsp;This is calculated by Shopee Logistics Services once buyer ships out if there is historical tracking data available from third party logistics provider.</p><p><br /></p><p>Note: Only available for <b>Normal RR with integrated reverse logistics.</b></p>
	CollectionPinCode               string                                           `json:"collection_pin_code"`                 // [Required] <p>The collection Pin Code to enter for seller to collect parcel in collection point or locker.</p><p><br /></p><p>Note: Only available for <b>TW region</b>.</p>
	TrackingNumber                  string                                           `json:"tracking_number"`                     // [Required] <p>The tracking number for the reverse logistics&nbsp;(the logistics tracking number provided when the buyer ships the item back).</p><p><br /></p><p>Note:&nbsp;</p><p>- Only available for <b>Normal RR with integrated reverse logistics.</b></p><p>- If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller)</b>. Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	TrackingInfo                    []GetReverseTrackingInfoResponseDataTrackingInfo `json:"tracking_info"`                       // [Required] <p>The detailed tracking information list for the reverse logistics.<b></b></p><p><b><br /></b></p><p>Note:&nbsp;</p><p>- Only available for <b>Normal RR with integrated reverse logistics</b>, with the tracking information pushed by third party logistics provider to Shopee.</p><p>- If <b>validation_type = seller_validation</b>, there is <b>only one segment of reverse logistics (The buyer ships the return parcel directly back to the seller)</b>. Please use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain the reverse logistics tracking information.</p>
	PostReturnLogisticsStatus       string                                           `json:"post_return_logistics_status"`        // [Required] <p>Post-return logistics status, referring to the current status of the warehouse shipping the return parcel back to the seller in warehouse validation mode. See&nbsp;"Data Definition - Post Return Logistics Status".</p><p><br /></p><p>Note:&nbsp;</p><p>- Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and<b> validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>- If <b>validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics</b>:&nbsp;</p><p style="padding-left:2em;">- <b>The buyer first ships the return parcel back to the warehouse.</b> Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">- <b>The warehouse then ships the return parcel back to the seller.</b> Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API <b>returns information for both the first and second segments.</b> For Local Returns, if the second segment exists, the API <b>prioritizes and returns only the second segment information.</b></p>
	PostReturnLogisticsUpdateTime   int64                                            `json:"post_return_logistics_update_time"`   // [Required] <p>The last update time of the post-return logistics status where warehouse sends return parcel from warehouse to seller.</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>-&nbsp;If <b>validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments</b>.&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
	RtsTrackingNumber               string                                           `json:"rts_tracking_number"`                 // [Required] <p>The tracking number for the post-return logistics (the logistics tracking number used when the warehouse ships the parcel back to the seller). RTS stands for "Return to Seller".</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p>-<b>&nbsp;If validation_type = warehouse_validation</b> AND the warehouse <b>uses an integrated logistics channel</b> to ship the return parcel back to the seller, there are <b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments.</b>&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
	PostReturnLogisticsTrackingInfo []GetReverseTrackingInfoResponseDataTrackingInfo `json:"post_return_logistics_tracking_info"` // [Required] <p>Only&nbsp;available for <b>Normal RR with return_solution = 0 (Return and Refund)</b> and <b>validation_type = warehouse_validation</b>, and the warehouse ships the return parcel back to seller <b>using integrated reverse logistics.</b></p><p><b><br /></b></p><p>In this scenario, the tracking logistics from warehouse to seller is called "post-return logistics", with the tracking information pushed by third party logistics provider to Shopee.</p><p><br /></p><p>Note:&nbsp;</p><p>-&nbsp;If <b>validation_type = warehouse_validation</b>&nbsp;AND the warehouse&nbsp;<b>uses an integrated logistics channel</b>&nbsp;to ship the return parcel back to the seller, there are&nbsp;<b>two segments of reverse logistics:&nbsp;</b></p><p style="padding-left:2em;">-&nbsp;<b>The buyer first ships the return parcel back to the warehouse.</b>&nbsp;Use the fields reverse_logistics_status, reverse_logistics_update_time, tracking_number, and tracking_info to obtain tracking information for this first segment.</p><p style="padding-left:2em;">-&nbsp;<b>The warehouse then ships the return parcel back to the seller.</b>&nbsp;Use the fields post_return_logistics_status, post_return_logistics_update_time, rts_tracking_number, and post_return_logistics_tracking_info to obtain tracking information for this second segment (post-return logistics).</p><p style="padding-left:2em;">- For Cross-Border Returns, if the second segment exists, the API&nbsp;<b>returns information for both the first and second segments.</b>&nbsp;For Local Returns, if the second segment exists, the API&nbsp;<b>prioritizes and returns only the second segment information.</b></p>
}
type GetReverseTrackingInfoResponseDataTrackingInfo struct {
	UpdateTime          int64    `json:"update_time"`          // [Required] <p>The timestamps when reverse logistics info has been updated for Normal RR from warehouse to seller, pushed from third party logistics provider to Shopee.</p>
	TrackingDescription string   `json:"tracking_description"` // [Required] <p>The description of reverse logistics tracking info&nbsp;for Normal RR from warehouse to seller, pushed by third party logistics provider to Shopee.</p><p><br /></p><p>These would match the tracking description displayed to sellers on Seller Center return/refund detail page.</p>
	EpopImageList       []string `json:"epop_image_list"`      // [Required] <p>Image URLs of electronic proof of pickup (ePOP) after return parcel has been picked up from the warehouse for Normal RR with warehouse validation.</p>
	EpodImageList       []string `json:"epod_image_list"`      // [Required] <p>Image URLs of electronic proof of delivery (ePOD) after return parcel has been delivered to the seller for Normal RR with warehouse validation.</p>
}
type GetShippingCarrierRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.<br /></p>
}
type GetShippingCarrierResponse struct {
	BaseResponse                                // Common response fields
	Response     GetShippingCarrierResponseData `json:"response"` //
}
type GetShippingCarrierResponseData struct {
	IsShippingProofMandatory      bool                      `json:"is_shipping_proof_mandatory"`       // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	HasUploadedSellerArrangeProof bool                      `json:"has_uploaded_seller_arrange_proof"` // [Required] <p>To indicate whether seller has already uploaded shipping proof for this return.</p>
	ShippingProofTemplate         []ShippingProofTemplate   `json:"shipping_proof_template"`           // [Required] <p>To display list of request parameters needed to upload shipping proof.</p>
	ReverseLogisticsCarrierList   []ReverseLogisticsCarrier `json:"reverse_logistics_carrier_list"`    // [Required] <p>The list of logistics carriers available for sellers to choose.</p>
}
type ImageInfo struct {
	ImageId *string `json:"image_id,omitempty"` // [Optional] <p>Unique image_id.</p>
}
type Negotiation struct {
	NegotiationStatus  string  `json:"negotiation_status"`   // [Required] To indicate whether the seller can negotiate with the buyer. See "Data Definition - NegotiationStatus"
	LatestSolution     string  `json:"latest_solution"`      // [Required] To indicate what is the offer solution. See "Data Definition - ReturnSolution"
	LatestOfferAmount  float64 `json:"latest_offer_amount"`  // [Required] To indicate the refund amount in the latest offer solution
	LatestOfferCreator string  `json:"latest_offer_creator"` // [Required] To indicate which party made the latest offer
	CounterLimit       int64   `json:"counter_limit"`        // [Required] To indicate the remaining counter limit
	OfferDueDate       int64   `json:"offer_due_date"`       // [Required] To indicate offer_due_date
}
type OfferRequest struct {
	ProposedAdjustedRefundAmount *float64 `json:"proposed_adjusted_refund_amount,omitempty"` // [Optional] The new refund amount to be offered
	ProposedSolution             string   `json:"proposed_solution"`                         // [Required] The new solution to be offered. See "Data Definition - ReturnSolution"
	ReturnSn                     string   `json:"return_sn"`                                 // [Required] The serial number of return.
}
type OfferResponse struct {
	BaseResponse                   // Common response fields
	Response     OfferResponseData `json:"response"` //
}
type OfferResponseData struct {
	ReturnSn string `json:"return_sn"` // [Required] <p>The serial number of return.</p>
}
type OfferReturnRefund struct {
	Eligibility            bool    `json:"eligibility"`              // [Required] <p>To indicate whether Refund solution is available for sellers to select.</p>
	RefundAmountAdjustable bool    `json:"refund_amount_adjustable"` // [Required] <p>To indicate whether refund is adjustable for Refund solution.</p>
	MaxRefundAmount        float64 `json:"max_refund_amount"`        // [Required] <p>The max refund amount for Refund solution.&nbsp;Returned when refund_amount_adjustable is true.</p>
	MinRefundAmount        float64 `json:"min_refund_amount"`        // [Required] <p>The min refund amount for Refund solution.&nbsp;Returned when refund_amount_adjustable is true.</p>
}
type QueryProofRequest struct {
	ReturnSn string `json:"return_sn" url:"return_sn"` // [Required] <p>The serial number of return.</p>
}
type QueryProofResponse struct {
	BaseResponse                        // Common response fields
	Response     QueryProofResponseData `json:"response"` //
}
type QueryProofResponseData struct {
	Image       []ResponseDataImage `json:"image"`       // [Required]
	Video       []ResponseDataImage `json:"video"`       // [Required]
	Description string              `json:"description"` // [Required] <p>The text description in the dispute proof.<br /></p>
}
type ResponseDataImage struct {
	Url       string `json:"url"`       // [Required] <p>Uploaded proof image link, it is recommended to pass in the return url of api called convert_image.</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The proof image&nbsp;thumbnail.</p>
}
type Return struct {
	Image                    []string         `json:"image"`                       // [Required] Image URLs of return.
	Reason                   string           `json:"reason"`                      // [Required] <p>Indicates the original return reason submitted by the buyer when initiating the return request.</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason.</p><p><br /></p><p>Note: There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;If the platform updates the return reason during this process, the reassessed outcome will be provided separately in the&nbsp;<b>reassessed_request_reason</b>&nbsp;field.</p>
	TextReason               string           `json:"text_reason"`                 // [Required] Reason that buyer provide.
	ReassessedRequestReason  string           `json:"reassessed_request_reason"`   // [Required] <p>Indicates the return reason reassessed by the platform as more suitable.</p><p><br /></p><p>There may be cases where Shopee Agent updates the return request with a "Reassessed Return Reason" after reviewing more details about the buyer's return request and potentially after requesting evidence from the seller.&nbsp;</p><p><br /></p><p>Applicable values: See Data Definition- ReturnReason and Reassessed Request Reason. If no reassessment has been made, the value will be NONE.</p>
	ReturnSn                 string           `json:"return_sn"`                   // [Required] The serial number of return.
	RefundAmount             float64          `json:"refund_amount"`               // [Required] Amount of the refund.
	Currency                 string           `json:"currency"`                    // [Required] Currency of the return.
	CreateTime               int64            `json:"create_time"`                 // [Required] The time of return create.
	UpdateTime               int64            `json:"update_time"`                 // [Required] The time of modify return.
	Status                   string           `json:"status"`                      // [Required] Enumerated type that defines the current status of the return. Applicable values: See Data Definition- ReturnStatus.
	DueDate                  int64            `json:"due_date"`                    // [Required] The last time seller deal with this return.
	TrackingNumber           string           `json:"tracking_number"`             // [Required] The tracking number assigned by the shipping carrier for item shipment.
	DisputeReason            []string         `json:"dispute_reason"`              // [Required] The reason of seller dispute return. While the return has been disputed, this field is useful. Applicable values: See Data Definition- ReturnDisputeReason.
	DisputeTextReason        []string         `json:"dispute_text_reason"`         // [Required] The reason that seller provide. While the return has been disputed, this field is useful.
	NeedsLogistics           bool             `json:"needs_logistics"`             // [Required] Items to be sent back to seller. Can be either integrated/non-integrated.
	AmountBeforeDiscount     float64          `json:"amount_before_discount"`      // [Required] Order price before discount.
	User                     *User            `json:"user"`                        // [Required]
	Item                     []ReturnItem     `json:"item"`                        // [Required]
	OrderSn                  string           `json:"order_sn"`                    // [Required] Shopee's unique identifier for an order.
	ReturnShipDueDate        int64            `json:"return_ship_due_date"`        // [Required] The due date for buyer to ship order.
	ReturnSellerDueDate      int64            `json:"return_seller_due_date"`      // [Required] The due date for seller to deal with this return when buyer have shipped order.
	NegotiationStatus        string           `json:"negotiation_status"`          // [Required] Counter status. See "Data Definition - NegotiationStatus"
	SellerProofStatus        string           `json:"seller_proof_status"`         // [Required] <p>Proof status. See "Data Definition - SellerProofStatus"</p>
	SellerCompensationStatus string           `json:"seller_compensation_status"`  // [Required] Compensation status. See "Data Definition - SellerCompensationStatus"
	ReturnRefundType         string           `json:"return_refund_type"`          // [Required] <p>To indicate whether the return is RRBOC (Return/Refund request raised before Order Complete) or RRAOC&nbsp;(Return/Refund request raised after Order Complete).<br /></p>
	ReturnSolution           int64            `json:"return_solution"`             // [Required] <p>To indicate the most updated solution of the Return/Refund request (NOTE: this is not the solution during negotiation). Applicable value:&nbsp;</p><p>- 0: Return and Refund</p><p>- 1: Refund Only</p>
	IsSellerArrange          bool             `json:"is_seller_arrange"`           // [Required] <p>To indicate whether the return_sn is using the “Seller Arrange” return method. This would only be True for TW and BR.</p>
	IsShippingProofMandatory bool             `json:"is_shipping_proof_mandatory"` // [Required] <p>To indicate whether uploading shipping proof is mandatory for seller to confirm "Arrange Pickup" when is_seller_arrange = true.</p>
	ReturnRefundRequestType  int64            `json:"return_refund_request_type"`  // [Required] <p>To indicate the type of return refund request, whether it is a Normal RR request, an In-transit RR request, and a Return on the Spot:&nbsp;</p><p style>0:&nbsp;Normal RR (RR is raised by the buyer after delivery done / estimated delivery date)</p><p style>1: In-transit RR (RR is raised by the buyer while item is still in-transit to buyer)</p><p style>2: Return-on-the-Spot (RR is raised by the driver after buyer rejected parcel at delivery)</p><p><br /></p><p>For more details, see Data Definition- Return Refund Request Type.</p>
	ValidationType           string           `json:"validation_type"`             // [Required] <p>To indicate whether seller or warehouse will expect to receive the return parcel from buyer and validate the condition of the parcel:&nbsp;</p><p>- seller_validation&nbsp;</p><p>- warehouse_validation</p><p><br /></p><p>For more details, see Data Definition- ValidationType.</p>
	IsArrivedAtWarehouse     int64            `json:"is_arrived_at_warehouse"`     // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b> Indicates the parcel’s check-in status at the warehouse. This field helps sellers quickly determine whether the parcel has arrived at the warehouse or has been rejected.&nbsp;</p><p><br /></p><p>Applicable values:</p><p>1: Pending Inbound</p><p>2: Rejected</p><p>3: Inbound</p><p>4: Cancelled</p>
	FollowUpActionList       []FollowUpAction `json:"follow_up_action_list"`       // [Required] <p><b>[Only for validation_type =&nbsp;warehouse_validation]</b>&nbsp;Warehouse handling actions for each item in the parcel.</p>
}
type ReturnAddress struct {
	WhsId string `json:"whs_id"` // [Required] <p>To indicate the warehouse id where item will be returned to.&nbsp;<br /></p><p>Please call&nbsp;v2.shop.get_warehouse_detail to check the detailed warehouse information the item returned to with the field "location_id" of the&nbsp;v2.shop.get_warehouse_detail match to the field"whs_id"of the v2.return.get_return_detail.<br /></p><p><br /></p><p>For fulfillment by Shopee (FBS) &amp; multi warehouse sellers, R/R orders will be returned back to the nearest warehouse of buyer address instead of going back to only 1 default return address like a normal seller.If it's a normal seller, then the field will be response empty.<br /></p>
}
type ReturnItem struct {
	ModelId      int64    `json:"model_id"`       // [Required] Shopee's unique identifier for a variation of an item.
	Name         string   `json:"name"`           // [Required] Name of item in local language.
	Images       []string `json:"images"`         // [Required] Image URLs of item.
	Amount       int64    `json:"amount"`         // [Required] Amount of this item.
	ItemPrice    float64  `json:"item_price"`     // [Required] The price of item.
	IsAddOnDeal  bool     `json:"is_add_on_deal"` // [Required] To indicate if this item belongs to an addon deal.
	IsMainItem   bool     `json:"is_main_item"`   // [Required] To indicate if this item is main item or sub item. True means main item, false means sub item.
	AddOnDealId  int64    `json:"add_on_deal_id"` // [Required] The unique identity of an addon deal.
	ItemId       int64    `json:"item_id"`        // [Required] The id of item.
	ItemSku      string   `json:"item_sku"`       // [Required] The sku of item.
	VariationSku string   `json:"variation_sku"`  // [Required] The variation sku of item
}
type ReturnPickupAddress struct {
	Address  string `json:"address"`  // [Required] <p>To indicate receiver's address</p>
	Name     string `json:"name"`     // [Required] <p>To indicate receiver's name<br /></p>
	Phone    string `json:"phone"`    // [Required] <p>To indicate receiver's phone<br />[Only for TW non-integrated channel] Will return "****" when the "virtual_contact_number" is available<br /></p>
	Town     string `json:"town"`     // [Required] <p>To indicate receiver's town</p>
	District string `json:"district"` // [Required] <p>To indicate receiver's district<br /></p>
	City     string `json:"city"`     // [Required] <p>To indicate receiver's city</p>
	State    string `json:"state"`    // [Required] <p>To indicate receiver's state</p>
	Region   string `json:"region"`   // [Required] <p>To indicate receiver's region</p>
	Zipcode  string `json:"zipcode"`  // [Required] <p>To indicate receiver's zip code<br /></p>
}
type ReverseLogisticsCarrier struct {
	ReverseLogisticsCarrierId   int64  `json:"reverse_logistics_carrier_id"`   // [Required] <p>To indicate the id of the non-integrated reverse logistics channel used by seller.</p>
	ReverseLogisticsCarrierName string `json:"reverse_logistics_carrier_name"` // [Required] <p>To indicate the selected carrier name from the list of carrier names provided.</p>
}
type SampleEvidence struct {
	Type      int64  `json:"type"`      // [Required] <p>The type of sample evidence. Applicable value:&nbsp;</p><p>- 1: Image</p>
	Url       string `json:"url"`       // [Required] <p>The link of sample evidence.</p>
	Thumbnail string `json:"thumbnail"` // [Required] <p>The link of the thumbnail of sample evidence.</p>
}
type SellerCompensation struct {
	SellerCompensationStatus  string  `json:"seller_compensation_status"`   // [Required] To indicate whether the seller is eligible for raising a compensation request. See "Data Definition - SellerCompensationStatus"
	SellerCompensationDueDate int64   `json:"seller_compensation_due_date"` // [Required] To indicate the deadline for requesting the compensation
	CompensationAmount        float64 `json:"compensation_amount"`          // [Required] To indicate the compensation amount that the agent decided
}
type SellerProof struct {
	SellerProofStatus      string `json:"seller_proof_status"`      // [Required] <p>To indicate whether the seller needs to provide evidence when the return status is RETURN_JUDING, RETURN_SELLER_DISPUTE and RETURN_ACCEPTED. Applicable values: See Data Definition- SellerProofStatus.</p>
	SellerEvidenceDeadline int64  `json:"seller_evidence_deadline"` // [Required] <p>To indicate the deadline for submitting the evidence.</p>
}
type ShippingProofTemplate struct {
	IsTrackingNumberRequired     bool `json:"is_tracking_number_required"`      // [Required] <p>To indicate whether it is mandatory to provide tracking number when uploading shipping proof.</p>
	IsShippingImageFileMandatory bool `json:"is_shipping_image_file_mandatory"` // [Required] <p>To indicate whether it is mandatory to provide shipping image file(s) when uploading shipping proof.</p>
}
type UploadProofRequest struct {
	Description *string             `json:"description,omitempty"` // [Optional] <p>text description in the dispute proof<br /></p>
	Photo       []ResponseDataImage `json:"photo,omitempty"`       // [Optional]
	ReturnSn    string              `json:"return_sn"`             // [Required] <p>The serial number of return.</p>
}
type UploadProofResponse struct {
	BaseResponse             // Common response fields
	Response     interface{} `json:"response"` // Actual response data
}
type UploadShippingProofRequest struct {
	ImageIdList                 []ImageInfo `json:"image_id_list,omitempty"`                  // [Optional] <p>List of image_id of shipping proof image. Required when&nbsp;is_shipping_image_file_mandatory = true in v2.returns.get_shipping_carrier. Max: 3.</p><p>You can call the v2.media.upload_image to upload image and get the image_id, for this scenario, please pass the business = 2 and scene = 1.</p>
	Remarks                     *string     `json:"remarks,omitempty"`                        // [Optional] <p>Optional remarks</p>
	ReturnSn                    string      `json:"return_sn"`                                // [Required] <p>The serial number of return.</p>
	ReverseLogisticsCarrierId   int64       `json:"reverse_logistics_carrier_id"`             // [Required] <p>Unique ID of non-integrated reverse logistics channel used by seller.</p>
	ReverseLogisticsCarrierName *string     `json:"reverse_logistics_carrier_name,omitempty"` // [Optional] <p>Non-integrated reverse logistics channel name used by seller.</p>
	TrackingNumber              *string     `json:"tracking_number,omitempty"`                // [Optional] <p>Tracking number used in seller arrange. Required when&nbsp;is_tracking_number_required = true in v2.returns.get_shipping_carrier.</p>
}
type UploadShippingProofResponse struct {
	BaseResponse             // Common response fields
	Response     interface{} `json:"response"` // Actual response data
}
type User struct {
	Username string `json:"username"` // [Required] <p>Buyer's nickname, will be masked as "****" if it is a non-integrated return in TW region.</p>
	Email    string `json:"email"`    // [Required] <p>Buyer's email, will be empty if it is a non-integrated return in TW region.<br /></p>
	Portrait string `json:"portrait"` // [Required] <p>Buyer's portrait, will be empty if it is a non-integrated return in TW region.</p>
}
