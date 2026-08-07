package goshopee

type BuyerPaymentInfo struct {
	BuyerPaymentMethod    string  `json:"buyer_payment_method"`     // [Required] <p>The payment method used by buyer.</p>
	BuyerServiceFee       float64 `json:"buyer_service_fee"`        // [Required] <p>An additional service fee charged by shopee to Buyer at checkout. currently only applicable to ID region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerTaxAmount        float64 `json:"buyer_tax_amount"`         // [Required] <p>The tax amount paid by buyer.</p><p>currently this is a custom tax charged to CB orders in TW,CL,MX</p>
	BuyerTotalAmount      float64 `json:"buyer_total_amount"`       // [Required] <p>The total amount paid by buyer at checkout moment.<br /></p>
	ShopeevipSubtotal     float64 `json:"shopeevip_subtotal"`       // [Required] <p>The subscription fee paid by buyer for ShopeeVIP.</p>
	CreditCardPromotion   float64 `json:"credit_card_promotion"`    // [Required] <p>The promotion provided by credit card.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IcmsTaxAmount         float64 `json:"icms_tax_amount"`          // [Required] <p>The icms tax paid by buyer. this is only applicable to BR region<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ImportTaxAmount       float64 `json:"import_tax_amount"`        // [Required] <p>The import tax paid by buyer.&nbsp;<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InitialBuyerTxnFee    float64 `json:"initial_buyer_txn_fee"`    // [Required] <p>Transaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. ).&nbsp;Most regions would have this fee charged to buyer at checkout depending on the fee rules applied in each region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InsurancePremium      float64 `json:"insurance_premium"`        // [Required] <p>The insurance premium paid by buyer. Currently only applicable to some regions like PH, MY, ID, VN, SG and TH it is an initial value (will not be updated after RR/cancellation)</p>
	IofTaxAmount          float64 `json:"iof_tax_amount"`           // [Required] <p>The iof tax paid by buyer.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IsPaidByCreditCard    bool    `json:"is_paid_by_credit_card"`   // [Required] <p>Whether this order is paid by credit card. it's related to payment channel used at checkout.</p><p>Value: false,true</p>
	MerchantSubtotal      float64 `json:"merchant_subtotal"`        // [Required] <p>The total item price paid by buyer at checkout.</p><p>it is an initial value and will not be updated after RR/cancellation<br /></p>
	SellerVoucher         float64 `json:"seller_voucher"`           // [Required] <p>The voucher provided by seller to offset some value needs to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShippingFee           float64 `json:"shipping_fee"`             // [Required] <p>The shipping fee paid by buyer. (Only display for non cb sip affiliate shop.&nbsp;<br /></p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShippingFeeSstAmount  float64 `json:"shipping_fee_sst_amount"`  // [Required] <p>The shipping fee sst paid by buyer. Currently apply to MY region only&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShopeeVoucher         float64 `json:"shopee_voucher"`           // [Required] <p>The voucher provided by Shopee to offset the amount need to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ShopeeCoinsRedeemed   float64 `json:"shopee_coins_redeemed"`    // [Required] <p>This is an amount of coin used by buyer at checkout to offset some amount to be paid.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerPaidPackagingFee float64 `json:"buyer_paid_packaging_fee"` // [Required] <p>The fee charged to the buyer for packaging materials</p>
	TradeInBonus          float64 `json:"trade_in_bonus"`           // [Required] <p>The total amount of all trade-in bonuses applied to a transaction. This value is the summation of the bonuses contributed by all four parties: Shopee, Seller, Bank and Trade-in Partner.</p>
	BulkyHandlingFee      float64 `json:"bulky_handling_fee"`       // [Required] <p>A fee charged to the buyer for the special handling and transportation required for items that exceed a specified weight or dimension threshold. Only for ID local seller</p>
	DiscountPix           float64 `json:"discount_pix"`             // [Required] <p>The discount provided by PIX&nbsp;(Only applicable for BR region)</p>
	BcrsDeposit           float64 `json:"bcrs_deposit"`             // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA). This will be an initial value and will not update after RR/cancellation.</p>
	AdsVoucherDiscount    float64 `json:"ads_voucher_discount"`     // [Required] <p>The voucher provided by Ads team to offset the amount need to be paid by buyer.</p>
}
type Escrow struct {
	OrderSn           string  `json:"order_sn"`            // [Required] <p>Shopee's unique identifier for an order.</p>
	PayoutAmount      float64 `json:"payout_amount"`       // [Required] The settlement amount
	EscrowReleaseTime int64   `json:"escrow_release_time"` // [Required] The release time
}
type EscrowDetail struct {
	OrderSn           string                        `json:"order_sn"`             // [Required] <p>Shopee's unique identifier for an order.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	BuyerUserName     string                        `json:"buyer_user_name"`      // [Required] <p>The username of buyer.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ReturnOrderSnList []string                      `json:"return_order_sn_list"` // [Required] <p>The list of the serial number of return.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	OrderIncome       *EscrowDetailOrderIncome      `json:"order_income"`         // [Required]
	BuyerPaymentInfo  *EscrowDetailBuyerPaymentInfo `json:"buyer_payment_info"`   // [Required] <p>The buyer payment info at order checkout moment. (snapshot value)</p>
}
type EscrowDetailBuyerPaymentInfo struct {
	BuyerPaymentMethod    string  `json:"buyer_payment_method"`     // [Required] <p>The payment method used by buyer.</p>
	BuyerServiceFee       string  `json:"buyer_service_fee"`        // [Required] <p>An additional service fee charged by shopee to Buyer at checkout. currently only applicable to ID region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	BuyerTaxAmount        float64 `json:"buyer_tax_amount"`         // [Required] <p>The tax amount paid by buyer.</p><p>currently this is a custom tax charged to CB orders in TW,CL,MX</p>
	BuyerTotalAmount      float64 `json:"buyer_total_amount"`       // [Required] <p>The total amount paid by buyer at checkout moment.<br /></p>
	ShopeevipSubtotal     float64 `json:"shopeevip_subtotal"`       // [Required] <p>The subscription fee paid by buyer for ShopeeVIP.</p>
	BcrsDiscount          float64 `json:"bcrs_discount"`            // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA). This will be an initial value and will not update after RR/cancellation.</p>
	CreditCardPromotion   float64 `json:"credit_card_promotion"`    // [Required] <p>The promotion provided by credit card.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	IcmsTaxAmount         float64 `json:"icms_tax_amount"`          // [Required] <p>The icms tax paid by buyer. this is only applicable to BR region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	ImportTaxAmount       float64 `json:"import_tax_amount"`        // [Required] <p>The import tax paid by buyer.&nbsp;<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InitialBuyerTxnFee    float64 `json:"initial_buyer_txn_fee"`    // [Required] <p>Transaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. ).&nbsp;Most regions would have this fee charged to buyer at checkout depending on the fee rules applied in each region.<br />it is an initial value (will not be updated after RR/cancellation)<br /></p>
	InsurancePremium      float64 `json:"insurance_premium"`        // [Required] <p>The insurance premium paid by buyer. Currently only applicable to some regions like ID &amp; BR</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	IofTaxAmount          float64 `json:"iof_tax_amount"`           // [Required] <p>The iof tax paid by buyer. this is only applicable for BR region.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	IsPaidByCreditCard    bool    `json:"is_paid_by_credit_card"`   // [Required] <p>Whether this order is paid by credit card. it's related to payment channel used at checkout<br /></p>
	MerchantSubtotal      float64 `json:"merchant_subtotal"`        // [Required] <p>The total item price paid by buyer at checkout.</p><p>it is an initial value and will not be updated after RR/cancellation</p>
	SellerVoucher         float64 `json:"seller_voucher"`           // [Required] <p>The voucher provided by seller to offset some value needs to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShippingFee           float64 `json:"shipping_fee"`             // [Required] <p>The shipping fee paid by buyer. (Only display for non cb sip affiliate shop.&nbsp;<br /></p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShippingFeeSstAmount  float64 `json:"shipping_fee_sst_amount"`  // [Required] <p>The shipping fee sst paid by buyer. Currently apply to MY region only&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShopeeVoucher         float64 `json:"shopee_voucher"`           // [Required] <p>The voucher provided by Shopee to offset the amount need to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	ShopeeCoinsRedeemed   float64 `json:"shopee_coins_redeemed"`    // [Required] <p>This is an amount of coin used by buyer at checkout to offset some amount to be paid.&nbsp;</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
	BuyerPaidPackagingFee float64 `json:"buyer_paid_packaging_fee"` // [Required] <p>The fee charged to the buyer for packaging materials.</p>
	TradeInBonus          float64 `json:"trade_in_bonus"`           // [Required] <p>The total amount of all trade-in bonuses applied to a transaction. This value is the summation of the bonuses contributed by all four parties: Shopee, Seller, Bank and Trade-in Partner.</p>
	BulkyHandlingFee      float64 `json:"bulky_handling_fee"`       // [Required] <p>A fee charged to the buyer for the special handling and transportation required for items that exceed a specified weight or dimension threshold.Only for local ID sellers.</p>
	DiscountPix           float64 `json:"discount_pix"`             // [Required] <p>The discount provided by PIX&nbsp;(Only applicable for BR region)</p>
	BcrsDeposit           float64 `json:"bcrs_deposit"`             // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA). This will be an initial value and will not update after RR/cancellation.</p>
	AdsVoucherDiscount    float64 `json:"ads_voucher_discount"`     // [Required] <p>The voucher provided by Ads team to offset the amount need to be paid by buyer.</p><p>it is an initial value (will not be updated after RR/cancellation)</p>
}
type EscrowDetailOrderIncome struct {
	EscrowAmount                                       float64               `json:"escrow_amount"`                                             // [Required] <p>The total amount that the seller is expected to receive for the order and will change before order is completed.&nbsp;</p><p>For non cb sip affiliate shop (new formula):&nbsp;</p><p>escrow_amount=&nbsp;original_cost_of_goods_sold-original_shopee_discount+seller_return_refund+ shopee_discounts- voucher_from_seller- seller_coin_cash_back+ buyer_paid_shipping_fee- actual_shipping_fee+ shopee_shipping_rebate+ shipping_fee_discount_from_3pl- reverse_shipping_fee+ rsf_seller_protection_fee_claim_amount+ fsf_seller_protection_fee_claim_amount- final_return_to_seller_shipping_fee- seller_transaction_fee- service_fee- commission_fee- campaign_fee- shipping_seller_protection_fee_premium_amount- delivery_seller_protection_fee_premium_amount-final_escrow_product_gst- order_ams_commission fee- escrow_tax-sales_tax_on_lvg-reverse_shipping_fee_sst-shipping_fee_sst-withholding_tax-overseas_return_service_fee-vat_on_imported_goods - withholding_vat_tax - withholding_pit_tax -&nbsp;withholding_cit_tax - seller_order_processing_fee + buyer_paid_packaging_fee - trade_in_bonus_seller - fbs_fee -&nbsp;ads_escrow_top_up_fee_or_technical_support_fee - th_import_duty</p><br /><p><br /></p><p>For cb sip affiliate shop:&nbsp;escrow_amount=escrow_amount_pri * exchange_rate</p><p><br /></p><p>note: Return refund amount = if adjustable RR, will equal to drc_adjustable_refund</p>
	OrderOriginalPrice                                 float64               `json:"order_original_price"`                                      // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPrice                                      float64               `json:"original_price"`                                            // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OrderSellingPrice                                  float64               `json:"order_selling_price"`                                       // [Required] <p>This field value = sum of item unit price.selling price comes from the sum up of each item's unit price. For non-bundle deal case, this value will be same like order_original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo.It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	BcrsDeposit                                        float64               `json:"bcrs_deposit"`                                              // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA). This value might be updated after RR/cancellation.</p>
	OrderSellerDiscount                                float64               `json:"order_seller_discount"`                                     // [Required] <p>The total discount seller provided for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerDiscount                                     float64               `json:"seller_discount"`                                           // [Required] <p>Final sum of each item seller discount of a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderDiscountedPrice                               float64               `json:"order_discounted_price"`                                    // [Required] <p>The total discounted price for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	ShopeeDiscount                                     float64               `json:"shopee_discount"`                                           // [Required] <p>Final sum of each item Shopee discount of a specific order. This amount will return remaining rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )<br /></p>
	VoucherFromSeller                                  float64               `json:"voucher_from_seller"`                                       // [Required] <p>Final value of voucher provided by Seller for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	VoucherFromShopee                                  float64               `json:"voucher_from_shopee"`                                       // [Required] <p>Final value of voucher provided by Shopee for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	VoucherFromExternalParty                           float64               `json:"voucher_from_external_party"`                               // [Required] <p>Final value of voucher provided by External Party for the order. (Only display for non cb sip affiliate shop.)</p>
	Coins                                              float64               `json:"coins"`                                                     // [Required] <p>This value indicates the total amount offset when the buyer consumed Shopee Coins upon checkout. (Only display for non cb sip affiliate shop. )<br /></p>
	CrossBorderTax                                     float64               `json:"cross_border_tax"`                                          // [Required] <p>Amount incurred by Buyer for purchasing items outside of home country. Amount may change after Return Refund. (Only display for non cb sip affiliate shop. )<br /></p>
	PaymentPromotion                                   float64               `json:"payment_promotion"`                                         // [Required] <p>The amount offset via payment promotion. (Only display for non cb sip affiliate shop. )<br /></p>
	CommissionFee                                      float64               `json:"commission_fee"`                                            // [Required] <p>The commission fee charged by Shopee platform if applicable.</p><p>For CB SIP affiliate shop:&nbsp;commission_fee=commission_fee_pri * exchange_rate</p>
	ServiceFee                                         float64               `json:"service_fee"`                                               // [Required] <p>Amount charged by Shopee to seller for additional services.</p><p>For CB SIP affiliate shop:&nbsp;service_fee=service_fee_pri * exchange_rate</p>
	SellerTransactionFee                               float64               `json:"seller_transaction_fee"`                                    // [Required] <p>Tansaction fee paid by seller for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerLostCompensation                             float64               `json:"seller_lost_compensation"`                                  // [Required] <p>Compensation to seller in case of lost parcel. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerCoinCashBack                                 float64               `json:"seller_coin_cash_back"`                                     // [Required] <p>Value of coins provided by Seller for purchasing with his or her store for the order. (Only display for non cb sip affiliate shop. )<br /></p>
	EscrowTax                                          float64               `json:"escrow_tax"`                                                // [Required] <p>Cross-border tax imposed by the Indonesian government on sellers. (Only display for non cb sip affiliate shop. )<br /></p>
	FinalShippingFee                                   float64               `json:"final_shipping_fee"`                                        // [Required] <p>Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. (Only display for non cb sip affiliate shop. )<br /></p>
	ActualShippingFee                                  float64               `json:"actual_shipping_fee"`                                       // [Required] <p>The final shipping cost of order and it is positive. For Non-integrated logistics channel is 0. (Only display for non cb sip affiliate shop. )<br /></p>
	ShopeeShippingRebate                               float64               `json:"shopee_shipping_rebate"`                                    // [Required] <p>The platform shipping subsidy to the seller. (Only display for non cb sip affiliate shop. )<br /></p>
	ShippingFeeSst                                     float64               `json:"shipping_fee_sst"`                                          // [Required] <p>The Service Tax charged on Seller Paid Shipping Fee for forward shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. Definition of Seller Paid Shipping Fee is Actual Shipping Fee subtracted by sum of Shipping Fee Paid by Buyer and Shipping Rebate From Shopee. (Only applicable for non cb sip affiliate shop)</p>
	ShippingFeeDiscountFromPl                          float64               `json:"shipping_fee_discount_from_pl"`                             // [Required] <p>The discount of shipping fee from 3PL. Currently only applicable to ID. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerShippingDiscount                             float64               `json:"seller_shipping_discount"`                                  // [Required] <p>The shipping discount defined by seller. (Only display for non cb sip affiliate shop. )<br /></p>
	EstimatedShippingFee                               float64               `json:"estimated_shipping_fee"`                                    // [Required] <p>The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard. (Only display for non cb sip affiliate shop. )<br /></p>
	SellerVoucherCode                                  float64               `json:"seller_voucher_code"`                                       // [Required] <p>The list of voucher code provided by seller. (Only display for non cb sip affiliate shop. )<br /></p>
	DrcAdjustableRefund                                float64               `json:"drc_adjustable_refund"`                                     // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center.<br /></p>
	RefundAmountToBuyer                                float64               `json:"refund_amount_to_buyer"`                                    // [Required] <p>Amount returned to Seller in the event of Partial Return.<br /></p>
	CostOfGoodsSold                                    float64               `json:"cost_of_goods_sold"`                                        // [Required] <p>Final amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OriginalCostOfGoodsSold                            float64               `json:"original_cost_of_goods_sold"`                               // [Required] <p>Amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )<br /></p>
	OriginalShopeeDiscount                             float64               `json:"original_shopee_discount"`                                  // [Required] <p>Sum of each item Shopee discount of a specific order. This amount will return initial rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )<br /></p>
	Items                                              []OrderIncomeItems    `json:"items"`                                                     // [Required] <p>This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.<br /></p>
	EscrowAmountPri                                    float64               `json:"escrow_amount_pri"`                                         // [Required] <p>The total amount in the prim currency that the seller is expected to receive for the order and will change before order completed . escrow_amount_pri=original_price_pri-seller_return_refund_pri-commission_fee_pri-service_fee_pri-drc_adjustable_refund_pri.(Only display for cb sip affiliate order. )<br /></p>
	BuyerTotalAmountPri                                float64               `json:"buyer_total_amount_pri"`                                    // [Required] <p>The total amount that paid by buyer in the primary currency. (Only display for cb sip affiliate order. )</p>
	OriginalPricePri                                   float64               `json:"original_price_pri"`                                        // [Required] <p>The original price of the item before ANY promotion/discount in the primary currency. It returns the subtotal of that specific item if quantity exceeds 1.(Only display for cb sip affiliate order. )</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	SellerReturnRefundPri                              float64               `json:"seller_return_refund_pri"`                                  // [Required] <p>Amount returned to Seller in the event of Partial Return in the primary currency. (Only display for cb sip affiliate shop. )</p>
	CommissionFeePri                                   float64               `json:"commission_fee_pri"`                                        // [Required] <p>The commission fee charged by Shopee platform if applicable in the primary currency. (Only display for cb sip affiliate shop. )</p>
	ServiceFeePri                                      float64               `json:"service_fee_pri"`                                           // [Required] <p>Amount charged by Shopee to seller for additional services in the primary currency. (Only display for cb sip affiliate shop. )</p>
	DrcAdjustableRefundPri                             float64               `json:"drc_adjustable_refund_pri"`                                 // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center in the primary currency&nbsp;after proration. (Only applicable for cb sip affiliate shop.)</p>
	PriCurrency                                        string                `json:"pri_currency"`                                              // [Required] <p>The currency of the country or region where the shop that real seller operates. (Only display for cb sip affiliate shop. )</p>
	AffCurrency                                        string                `json:"aff_currency"`                                              // [Required] <p>The currency of the country or region where shop opened in. (Only display for cb sip affiliate shop. )</p>
	ExchangeRate                                       float64               `json:"exchange_rate"`                                             // [Required] <p>Exchange rate from primary shop currency to affiliate shop currency.<br /></p>
	ReverseShippingFee                                 float64               `json:"reverse_shipping_fee"`                                      // [Required] <p>Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.<br /></p>
	ReverseShippingFeeSst                              float64               `json:"reverse_shipping_fee_sst"`                                  // [Required] <p>The Service Tax charged on Reverse Shipping Fee for reverse shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. (Only applicable for non cb sip affiliate shop)<br /></p>
	FinalProductProtection                             float64               `json:"final_product_protection"`                                  // [Required] <p>The total amount of product protection purchased during placing an order.&nbsp;</p>
	CreditCardPromotion                                float64               `json:"credit_card_promotion"`                                     // [Required] <p>This value indicate the offset via credit card promotion.<br /></p>
	CreditCardTransactionFee                           float64               `json:"credit_card_transaction_fee"`                               // [Required] <p>This value indicate the total transaction fee.</p><p>credit_card_transaction_fee=buyer_transaction_fee+seller_transaction_fee</p>
	FinalProductVatTax                                 float64               `json:"final_product_vat_tax"`                                     // [Required] <p>Value-added Tax is required for online purchases based on EU Value-added Tax regulations . (Only display for non cb sip affiliate shop. )<br /></p>
	FinalShippingVatTax                                float64               `json:"final_shipping_vat_tax"`                                    // [Required] <p>Value-added Tax for product price is required for online purchases based on EU Value-added Tax regulations. (Only applicable for non cb sip affiliate shop. )<br /></p>
	CampaignFee                                        float64               `json:"campaign_fee"`                                              // [Required] <p>The campaign fee charged by Shopee platform. Only available for some local Indonesian shops.<br /></p>
	SipSubsidy                                         float64               `json:"sip_subsidy"`                                               // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. (Only available for CB SIP A Shops)<br /></p>
	SipSubsidyPri                                      float64               `json:"sip_subsidy_pri"`                                           // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. This value is in the primary currency. (Only available for CB SIP A Shops)<br /></p>
	RsfSellerProtectionFeeClaimAmount                  float64               `json:"rsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The insurance claim amount if seller opt in to insurance program. this covers Reverse shipping Fee in the case of RR. As per Jun 2024:<br />- For ID &amp; MY Local: After Extension on coverage to FSF due to RR. this claim amount will consist of FSF + RSF claim.<br />- For PH local: This will only cover RSF claim<br /><br />will be updated if there's any RR/cancellation<br /></p>
	RsfSellerProtectionFeePremiumAmount                float64               `json:"rsf_seller_protection_fee_premium_amount"`                  // [Required] <p>Amount charged by Shopee to seller for additional services. Only apply for PH local orders.<br /></p>
	FinalEscrowProductGst                              float64               `json:"final_escrow_product_gst"`                                  // [Required] <p>Goods and Service Tax for product price is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore)<br /></p>
	FinalEscrowShippingGst                             float64               `json:"final_escrow_shipping_gst"`                                 // [Required] <p>Goods and Service Tax for shipping fee is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore.)<br /></p>
	DeliverySellerProtectionFeePremiumAmount           float64               `json:"delivery_seller_protection_fee_premium_amount"`             // [Required] <p>[Currently apply to ID &amp; local orders only] An insurance premium charged to seller at the time parcel is picked up by 3PL for insurance in case of parcel lost/damaged by 3PL.<br /></p>
	OrderAmsCommissionFee                              float64               `json:"order_ams_commission_fee"`                                  // [Required] <p>The amount of affiliate commission fee for this order. Applicable for orders sold via the Affiliate Program.<br /></p>
	BuyerPaymentMethod                                 float64               `json:"buyer_payment_method"`                                      // [Required] <p>The payment method buyer used when do the order checkout.<br /></p>
	InstalmentPlan                                     float64               `json:"instalment_plan"`                                           // [Required] <p>The instalment plan buyer chosen when do the order checkout. Only applicable when payment method support instalment.<br /></p>
	SalesTaxOnLvg                                      float64               `json:"sales_tax_on_lvg"`                                          // [Required] <p>Sales Tax on Low Value Goods (LVG) is required for imported goods (overseas orders) based on Malaysia SST regulations&nbsp;for selective orders (e.g. CB LVG orders in MY)<br /></p>
	WithholdingTax                                     float64               `json:"withholding_tax"`                                           // [Required] <p>Only for PH and ID local shops.<br /><br /></p><p><b>PH:&nbsp;</b>According to regulations issued by Bureau of Internal Revenue in PH, the Withholding Tax is applied to the gross remittances sent by Shopee to online suppliers of goods and services.<br /></p><p><br /></p><p><b>ID:&nbsp;</b>According to regulations issued by Directorate General of Taxation in ID, the Withholding Tax is applied to the income stated in the invoice generated via Shopee related to Seller and/or Merchants' sales in Shopee's platform.</p>
	OverseasReturnServiceFee                           float64               `json:"overseas_return_service_fee"`                               // [Required] <p>This is overseas return service fee charged to sellers who register ORS program.(Only applicable for non cb sip affiliate shop)<br /></p>
	ProratedCoinsValueOffsetReturnItems                float64               `json:"prorated_coins_value_offset_return_items"`                  // [Required] <p>This is the prorated value from cash equivalent of coin offset due to adjustable RR.This field will only be updated when there is an adjustable RR. If it's a full RR or normal order will response 0.<br /></p>
	ProratedShopeeVoucherOffsetReturnItems             float64               `json:"prorated_shopee_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from shopee voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedSellerVoucherOffsetReturnItems             float64               `json:"prorated_seller_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from seller voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoBankOffsetReturnItems   float64               `json:"prorated_payment_channel_promo_bank_offset_return_items"`   // [Required] <p>This is the prorated value from bank payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoShopeeOffsetReturnItems float64               `json:"prorated_payment_channel_promo_shopee_offset_return_items"` // [Required] <p>This is the prorated value from shopee payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	FsfSellerProtectionFeeClaimAmount                  float64               `json:"fsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The claim amount given to seller if seller opt in to shipping fee service program. this covers Forward Shipping Fee cost in the case of cancelled due to Failed delivery. only apply to PH Local as per Jun 2024.<br /></p>
	ShippingSellerProtectionFeeAmount                  float64               `json:"shipping_seller_protection_fee_amount"`                     // [Required] <p>Service fee charged to seller in MY,ID,PH,BR Local (as per Jun 2024) due to additional program purchased<br /></p>
	FinalReturnToSellerShippingFee                     float64               `json:"final_return_to_seller_shipping_fee"`                       // [Required] <p>The amount of fee charged to seller (if any) for the failed delivery order&nbsp;</p>
	VatOnImportedGoods                                 float64               `json:"vat_on_imported_goods"`                                     // [Required] <p>7% VAT is charged for imported goods entering Thailand<br /></p>
	WithholdingVatTax                                  float64               `json:"withholding_vat_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold VAT tax applicable to all VN business household and VN individual sellers</p>
	WithholdingPitTax                                  float64               `json:"withholding_pit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to all VN business household and VN individual sellers</p>
	WithholdingCitTax                                  float64               `json:"withholding_cit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to corporate sellers selling in VN region</p>
	TaxRegistrationCode                                string                `json:"tax_registration_code"`                                     // [Required] <p>For VN Withholding Tax. This is the Tax Registration Number (TRN) from Seller Info (Business information) of the seller at the point of order creation</p>
	SellerOrderProcessingFee                           float64               `json:"seller_order_processing_fee"`                               // [Required] <p>Order Processing Fee is the amount charged to sellers for every order created.</p>
	BuyerPaidPackagingFee                              float64               `json:"buyer_paid_packaging_fee"`                                  // [Required] <p>The fee charged to the buyer for packaging materials.</p>
	TradeInBonusSeller                                 float64               `json:"trade_in_bonus_seller"`                                     // [Required] <p>The discount provided by Seller/Retailers for buyers who opt for trade-in.</p>
	FbsFee                                             float64               `json:"fbs_fee"`                                                   // [Required] <p>Fulfilled by Shopee (FBS) Fee applied to this order, covering costs such as handling and storage and packaging. Only for PH Local Orders.</p>
	NetCommissionFee                                   float64               `json:"net_commission_fee"`                                        // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net commission fee applied to the order, calculated as the sum of all net commission fee items.</p><p>-only for BR local sellers.</p>
	NetServiceFee                                      float64               `json:"net_service_fee"`                                           // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net service fee applied to the order, calculated as the sum of all net service fee items.</p><p>-only for BR local sellers.</p>
	NetCommissionFeeInfoList                           *NetCommissionFeeInfo `json:"net_commission_fee_info_list"`                              // [Required] <p>Returns a breakdown of the net commission fees.</p><p>-only for BR local sellers.</p>
	NetServiceFeeInfoList                              *NetServiceFeeInfo    `json:"net_service_fee_info_list"`                                 // [Required] <p>Returns a breakdown of the net service fees.</p><p>-only for BR local sellers.</p>
	SellerProductRebate                                *SellerProductRebate  `json:"seller_product_rebate"`                                     // [Required] <p>The shopee rebate borne by seller.</p><p>-only for BR local sellers.</p>
	PixDiscount                                        float64               `json:"pix_discount"`                                              // [Required] <p>[BR only]Final sum of pix discount of a specific order. (Only display for non cb sip affiliate shop.)</p>
	ProratedPixDiscountOffsetReturnItems               float64               `json:"prorated_pix_discount_offset_return_items"`                 // [Required] <p>[BR only]This is the prorated value from pix discount due to adjustable RR. This field will only be updated when there is an adjustable RR. If it's a full RR or normal order, will response 0.</p>
	AdsEscrowTopUpFeeOrTechnicalSupportFee             float64               `json:"ads_escrow_top_up_fee_or_technical_support_fee"`            // [Required] <p>Includes both ads escrow top up fee (auto escrow top up to your ads balance) and technical support fee (charged by Shopee)</p><p>The actual fee type included in this field varies depending on the seller type and selling region, and may represent one of the following in Shopee Seller Center:</p><p>Ads Escrow Top-Up Fee</p><p>For local MY TH SG VN PH ID sellers and CNCB / JPCB / KRCB sellers selling in PH and ID</p><p>For JPCB sellers selling in SG, MY, TH, and VN</p><p>Technical Support Fee</p><p>For CNCB sellers selling in SG, MY, TH, and VN</p><p>Traffic Growth Fee</p><p>For KRCB sellers selling in SG, MY, TH, and VN</p>
	ThImportDuty                                       float64               `json:"th_import_duty"`                                            // [Required] <p>[TH only] Import Duty collected for imported goods entering Thailand</p>
	RemainingVoucher                                   float64               `json:"remaining_voucher"`                                         // [Required] <p>[Only for BR local shop]Represents the portion of Shopee voucher that is not consumed after fee offset.</p>
}
type GenerateIncomeReportRequest struct {
	ReleaseTimeFrom int64 `json:"release_time_from"` // [Required] <p>Start time in epoch<br /></p>
	ReleaseTimeTo   int64 `json:"release_time_to"`   // [Required] <p>End time in epoch</p>
}
type GenerateIncomeReportResponse struct {
	BaseResponse                                  // Common response fields
	Response     GenerateIncomeReportResponseData `json:"response"`      // Response data
	Msg          string                           `json:"msg,omitempty"` // <p>error message</p>
}
type GenerateIncomeReportResponseData struct {
	Id int64 `json:"id"` // [Required] <p>Identifier of income report file.</p>
}
type GenerateIncomeStatementRequest struct {
	ReleaseTimeFrom int64 `json:"release_time_from" url:"release_time_from"` // [Required] <p>The release_time_from must be</p><p><b>- Monday</b>&nbsp;(local time) for a weekly report</p><p>- <b>The 1st day</b> (local time) of a Month for a monthly report</p>
	ReleaseTimeTo   int64 `json:"release_time_to" url:"release_time_to"`     // [Required] <p>The release_time_to must be</p><p>- <b>Sunday</b> (local time) for a weekly report</p><p>-&nbsp;<b>The last day</b>&nbsp;(local time) of a Month for a monthly report</p>
	StatementType   int64 `json:"statement_type" url:"statement_type"`       // [Required] <p>STATEMENT_TYPE_WEEKLY = 1;</p><p>STATEMENT_TYPE_MONTHLY = 2;</p><p><br /></p><p>Local seller Income statement requires this value to be set.</p><p>CB seller income statement does not require this.</p>
}
type GenerateIncomeStatementResponse struct {
	BaseResponse                                     // Common response fields
	Response     GenerateIncomeStatementResponseData `json:"response"` // Response data
}
type GenerateIncomeStatementResponseData struct {
	Id int64 `json:"id"` // [Required] <p>Identifier of income statement file.</p>
}
type GetBillingTransactionInfoRequest struct {
	BillingTransactionInfoType int64    `json:"billing_transaction_info_type"`  // [Required] <p>Billing transaction types: 1: TO_RELEASE, 2:&nbsp;RELEASED</p>
	Cursor                     string   `json:"cursor"`                         // [Required] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
	EncryptedPayoutIds         []string `json:"encrypted_payout_ids,omitempty"` // [Optional] <p>encrypted_payout_id get from API: v2.get_payout_info</p><p><br /></p><p><br /></p><p>when&nbsp;encrypted_payout_id provided and billing_transaction_info_type=2, we will return the "released" billing items under this payout.</p><p>when&nbsp;encrypted_payout_id not provided, we will return the "to release" billing items under hasn't form payout yet<br /></p><p><br /></p><p>Max length: 100</p>
	PageSize                   int64    `json:"page_size"`                      // [Required] <p>Number of pages returned max:100<br /></p>
}
type GetBillingTransactionInfoResponse struct {
	BaseResponse                                       // Common response fields
	Response     GetBillingTransactionInfoResponseData `json:"response"` // Response data
}
type GetBillingTransactionInfoResponseData struct {
	Transactions *Transactions `json:"transactions"` // [Required]
	More         bool          `json:"more"`         // [Required]
	NextCursor   string        `json:"next_cursor"`  // [Required]
}
type GetEscrowDetailBatchRequest struct {
	OrderSnList []string `json:"order_sn_list"` // [Required] <p>Shopee's unique identifier for an order.</p><p>limit [1,50]&nbsp;<br /></p><p>The number of recommended requests ranges from 1 to 20 orders.</p>
}
type GetEscrowDetailBatchResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetEscrowDetailBatchResponseData `json:"response"` // Response data
}
type GetEscrowDetailBatchResponseData struct {
	EscrowDetail *EscrowDetail `json:"escrow_detail"` // [Required] <p>The escrow detail for one order</p>
}
type GetEscrowDetailRequest struct {
	OrderSn string `json:"order_sn"` // [Required] Shopee's unique identifier for an order.
}
type GetEscrowDetailResponse struct {
	BaseResponse                             // Common response fields
	Response     GetEscrowDetailResponseData `json:"response"` // Response data
}
type GetEscrowDetailResponseData struct {
	OrderSn           string            `json:"order_sn"`             // [Required]  Shopee's unique identifier for an order.
	BuyerUserName     string            `json:"buyer_user_name"`      // [Required] The username of buyer.
	ReturnOrderSnList []string          `json:"return_order_sn_list"` // [Required] The list of the serial number of return.
	OrderIncome       *OrderIncome      `json:"order_income"`         // [Required]
	BuyerPaymentInfo  *BuyerPaymentInfo `json:"buyer_payment_info"`   // [Required] <p>The buyer payment info at order checkout moment. (snapshot value)</p>
}
type GetEscrowListRequest struct {
	PageNo          *int64 `json:"page_no,omitempty"`   // [Optional] The page number  min:1  default:1
	PageSize        *int64 `json:"page_size,omitempty"` // [Optional] Number of pages returned  max:100  default:40
	ReleaseTimeFrom int64  `json:"release_time_from"`   // [Required] Query start time
	ReleaseTimeTo   int64  `json:"release_time_to"`     // [Required] Query end time
}
type GetEscrowListResponse struct {
	BaseResponse                           // Common response fields
	Response     GetEscrowListResponseData `json:"response"` // Response data
}
type GetEscrowListResponseData struct {
	EscrowList []Escrow `json:"escrow_list"` // [Required] <p>The list of escrow order sn.</p>
	More       bool     `json:"more"`        // [Required] <p>This is to indicate whether the escrow list is more than one page. If this value is true, you may want to continue to check next page to retrieve escrow orders.<br /></p>
}
type GetIncomeDetailRequest struct {
	Cursor       *string `json:"cursor,omitempty"` // [Optional] <p>Pagination token for the next set of results. Use an empty string "" for the first request.</p>
	DateFrom     string  `json:"date_from"`        // [Required] <p>Start date (YYYY-MM-DD) of the income reference period. This field is only used for Income Status = Released, the other statuses will display all records currently in that status.<br /></p><p><br /></p><p>For income Status = Released,&nbsp;For Released → Payout released date:</p><p>1. date_to must be later than date_from</p><p>2. date range cannot exceed 14 days</p><p>3. Input must follow valid date format.</p>
	DateTo       string  `json:"date_to"`          // [Required] <p>End date (YYYY-MM-DD) of the income reference period. Must be later than date_from. This field is only used for Income Status = Released, the other statuses will display all records currently in that status.<br /></p><p><br /></p><p>For income Status = Released,&nbsp;For Released → Payout released date:</p><p>1. date_to must be later than date_from</p><p>2. date range cannot exceed 14 days</p><p>3. Input must follow valid date format.</p>
	IncomeStatus int64   `json:"income_status"`    // [Required] <p>Status of Seller Income payout (Enum - Desc)</p><p><br /></p><p>Local</p><p>1 -Released</p><p>2 - Pending</p><p>CB</p><p>0 - To Release</p><p>1 -&nbsp;Released</p><p>2 - Pending</p>
	PageSize     int64   `json:"page_size"`        // [Required] <p>Number of income detail records to retrieve per page</p>
}
type GetIncomeDetailResponse struct {
	BaseResponse                   // Common response fields
	IncomeDetailList *IncomeDetail `json:"income_detail_list,omitempty"` // <p>List of income detail records returned for the specified time range and status.</p>
}
type GetIncomeOverviewRequest struct {
	IncomeStatus *int64 `json:"income_status,omitempty" url:"income_status,omitempty"` // [Optional] <p><b><u>Status of Seller Income payout (Enum - Desc)</u></b></p><p><br /></p><p><b><u>Local Shop</u></b></p><p>1 -Released</p><p>2 - Pending</p><p><b><u><br /></u></b></p><p><b><u>CB Shop</u></b></p><p>0 - To Release</p><p>1 -&nbsp;Released</p><p>2 - Pending</p><p><br /></p><p>Note: By default, if Income Status was not provided in the request params (non mandatory), API response will return all values for all Income status based on either Local/CB</p>
}
type GetIncomeOverviewResponse struct {
	BaseResponse                               // Common response fields
	Response     GetIncomeOverviewResponseData `json:"response"` // Response data
}
type GetIncomeOverviewResponseData struct {
	LatestPayoutDate string       `json:"latest_payout_date"` // [Required] <p>The latest payout date for the released income. Format: YYYY-MM-DD. Only for CN shops.</p>
	TotalIncome      *TotalIncome `json:"total_income"`       // [Required] <p>Object containing total income components.</p>
}
type GetIncomeReportRequest struct {
	IncomeReportId int64 `json:"income_report_id"` // [Required] <p>The identifier for income report file request.<br /></p>
}
type GetIncomeReportResponse struct {
	BaseResponse                             // Common response fields
	Response     GetIncomeReportResponseData `json:"response"`      // Response data
	Msg          string                      `json:"msg,omitempty"` // <p>Error Message</p>
}
type GetIncomeReportResponseData struct {
	Id            int64  `json:"id"`             // [Required] <p>The identifier for income statement file request.<br /></p>
	FileName      string `json:"file_name"`      // [Required] <p>Income report file name.<br /></p>
	Status        int64  `json:"status"`         // [Required] <p>STATUS_INVALID = 0;</p><p>STATUS_PROCESSING = 1;</p><p>STATUS_DOWNLOADABLE = 2;</p><p>STATUS_DOWNLOADED = 3;</p><p>STATUS_FAILED = 4;</p>
	GeneratedTime int64  `json:"generated_time"` // [Required] <p>File generation time.<br /></p>
	FileLink      string `json:"file_link"`      // [Required] <p>Link to download income report file.<br /></p>
}
type GetIncomeStatementRequest struct {
	IncomeStatementId int64 `json:"income_statement_id" url:"income_statement_id"` // [Required] <p>The identifier for income statement file request.<br />return from the API&nbsp;<span style="font-size:14px;">v2.payment.generate_income_statement</span></p><br />
}
type GetIncomeStatementResponse struct {
	BaseResponse                                // Common response fields
	Response     GetIncomeStatementResponseData `json:"response"` // Response data
}
type GetIncomeStatementResponseData struct {
	Id            int64  `json:"id"`             // [Required] <p>The identifier for income statement file request.<br /></p>
	FileName      string `json:"file_name"`      // [Required] <p>Income statement file name.</p>
	Status        int64  `json:"status"`         // [Required] <p>STATUS_INVALID = 0;</p><p>STATUS_PROCESSING = 1;</p><p>STATUS_DOWNLOADABLE = 2;</p><p>STATUS_DOWNLOADED = 3;</p><p>STATUS_FAILED = 4;</p>
	GeneratedTime int64  `json:"generated_time"` // [Required] <p>File generation time.</p>
	FileLink      string `json:"file_link"`      // [Required] <p>Link to download income statement file.</p>
}
type GetItemInstallmentStatusRequest struct {
	ItemIdList []int64 `json:"item_id_list"` // [Required] Item id array, Max :100
}
type GetItemInstallmentStatusResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetItemInstallmentStatusResponseData `json:"response"` // Response data
}
type GetItemInstallmentStatusResponseData struct {
	ItemInstallmentList []ItemInstallment `json:"item_installment_list"` // [Required]
	ItemPlanAhoraList   []ItemPlanAhora   `json:"item_plan_ahora_list"`  // [Required] Only applicable for local AR sellers.
}
type GetPaymentMethodListResponse struct {
	BaseResponse                                  // Common response fields
	Response     GetPaymentMethodListResponseData `json:"response"` // Response data
}
type GetPaymentMethodListResponseData struct {
	PaymentMethod []string `json:"payment_method"` // [Required]
	Region        string   `json:"region"`         // [Required]
}
type GetPayoutDetailRequest struct {
	PageNo         int64 `json:"page_no"`          // [Required] The page number  min:1  default:1
	PageSize       int64 `json:"page_size"`        // [Required] Number of pages returned  max:100
	PayoutTimeFrom int64 `json:"payout_time_from"` // [Required] <p>Strat time. Maximum time range is 15 days</p>
	PayoutTimeTo   int64 `json:"payout_time_to"`   // [Required] End time
}
type GetPayoutDetailResponse struct {
	BaseResponse                             // Common response fields
	Response     GetPayoutDetailResponseData `json:"response"` // Response data
}
type GetPayoutDetailResponseData struct {
	More       bool     `json:"more"`        // [Required]
	PayoutList []Payout `json:"payout_list"` // [Required]
}
type GetPayoutInfoRequest struct {
	Cursor         string `json:"cursor"`           // [Required] <p>Specifies the starting entry of data to return in the current call. Default is "". If data is more than one page, the offset can be some entry to start next call.<br /></p>
	PageSize       int64  `json:"page_size"`        // [Required] <p>Number of pages returned max:100<br /></p>
	PayoutTimeFrom int64  `json:"payout_time_from"` // [Required] <p>Start time. Maximum time range is 15 days<br /></p>
	PayoutTimeTo   int64  `json:"payout_time_to"`   // [Required] <p>Payout End time<br /></p>
}
type GetPayoutInfoResponse struct {
	BaseResponse                           // Common response fields
	Response     GetPayoutInfoResponseData `json:"response"` // Response data
}
type GetPayoutInfoResponseData struct {
	PayoutList *ResponseDataPayout `json:"payout_list"` // [Required]
	More       bool                `json:"more"`        // [Required] <p>True or False<br /></p>
	NextCursor string              `json:"next_cursor"` // [Required] <p>used for next batch data query. will return empty when all data been returned<br /></p>
}
type GetShopInstallmentStatusResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetShopInstallmentStatusResponseData `json:"response"` // Response data
}
type GetShopInstallmentStatusResponseData struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required] <p>The installment status for the shop</p>
}
type GetWalletTransactionListRequest struct {
	CreateTimeFrom     *int64  `json:"create_time_from,omitempty"`     // [Optional] The create_time_from field is the starting date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	CreateTimeTo       *int64  `json:"create_time_to,omitempty"`       // [Optional] The create_time_to field is the ending date range. The maximum date range that may be specified with the create_time_from and create_time_to fields is 15 days.
	MoneyFlow          *string `json:"money_flow,omitempty"`           // [Optional] <p>It's to indicate whether user wants to only return :&nbsp;</p><p>MONEY_IN = addition&nbsp;<br />MONEY_OUT = Deduction</p><p><br /></p><p>if not specified, we will return all</p><p>Note special case for TW JKO Pay, we will ignore Money_flow&nbsp;</p>
	PageNo             int64   `json:"page_no"`                        // [Required] Specifies the starting entry of data to return in the current call. Default is 0. if data is more than one page, the offset can be some entry to start next call.
	PageSize           int64   `json:"page_size"`                      // [Required] If many transactions are available to retrieve, you may need to call GetTransactionList multiple times to retrieve all the data. Each result set is returned as a page of entries. Default is 40. Use the Pagination filters to control the maximum number of entries (<= 100) to retrieve per page (i.e., per call), the offset number to start next call. This integer value is usUed to specify the maximum number of entries to return in a single ""page"" of data.
	TransactionTabType *string `json:"transaction_tab_type,omitempty"` // [Optional] <p><b>NOTE: Only 1 'transaction tab type' value should be passed in.</b></p><p>Passing in more than 1 value (eg: comma separated values) will return default response. This is because the request param treats the value passed in as a single string.</p><p><br /></p><p>This to indicates the updated filtering type that client can use to specify which transaction type we want to return. it will have :&nbsp;<br /><br />Default<br />wallet_order_income<br />wallet_adjustment_filter</p><p>wallet_wallet_payment</p><p>wallet_refund_from_order</p><p>wallet_withdrawals</p><p>fast_escrow_repayment</p><p>fast_pay</p><p>seller_loan<br />corporate_loan<br />pix_transactions_filter</p><p>open_finance_transactions_filter&nbsp;<br /><br />Note for BR, wallet txn type that linked to&nbsp;pix_transactions_filter&nbsp; and&nbsp;open_finance_transactions_filter are classified as&nbsp;default&nbsp;&nbsp;type tab instead. therefore for Open API client who wants to query these 2 trx can put default as the filter in this type<br /></p>
	TransactionType    *string `json:"transaction_type,omitempty"`     // [Optional] <p>Transaction type APIs:&nbsp;<br /></p><p>ESCROW_VERIFIED_ADD = 101;&nbsp; // Escrow has been verified and paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ESCROW_VERIFIED_MINUS = 102; // Escrow has been verified and charged from seller as escrow amount is negative&nbsp;&nbsp;&nbsp;&nbsp;</p><p>WITHDRAWAL_CREATED = 201; // The seller has created a withdrawal, so it’s deducted from balance &nbsp; &nbsp; &nbsp;<br /></p><p>WITHDRAWAL_COMPLETED = 202; // The withdrawal has been completed, so the ongoing amount decreases.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>WITHDRAWAL_CANCELLED = 203; // The withdrawal has been canceled, so the amount is added back to the seller balance. Ongoing amount decreases as well.&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ADJUSTMENT_ADD = 401; // One adjustment item has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ADJUSTMENT_MINUS = 402; // One adjustment item has been charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;FBS_ADJUSTMENT_ADD = 404; //One adjustment item related to Shopee fulfillment order is added to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FBS_ADJUSTMENT_MINUS = 405; // One adjustment item related to Shopee fulfillment order is deducted from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>ADJUSTMENT_CENTER_ADD = 406; // One adjustment item has been added to seller wallet&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;ADJUSTMENT_CENTER_DEDUCT = 407; // One adjustment item has been deducted from seller wallet&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FSF_COST_PASSING_DEDUCT = 408; FSF cost passing for canceled/invalid orders&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;PERCEPTION_VAT_TAX_DEDUCT = 409; Extra charge for perception regime VAT tax (Argentina)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;PERCEPTION_TURNOVER_TAX_DEDUCT = 410; Extra charge for perception regime turnover tax (Argentina)&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>PAID_ADS_CHARGE = 450; // Paid ads are charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>PAID_ADS_REFUND = 451; // Paid ads are refunded to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>FAST_ESCROW_DISBURSE = 452; // ADD. // The first disbursement of fast escrow has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>AFFILIATE_ADS_SELLER_FEE = 455; // DEDUCT // Affiliate ads seller fee is charged from seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;AFFILIATE_ADS_SELLER_FEE_REFUND = 456; // ADD // Affiliate ads seller fee is refunded to seller</p><p>FAST_ESCROW_DEDUCT = 458; // Fast escrow is deducted from seller balance in the event of return and refund&nbsp;</p><p>FAST_ESCROW_DISBURSE_REMAIN = 459; // The second disbursement of fast escrow has been paid to seller&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p><p>AFFILIATE_FEE_DEDUCT = 460; // Affiliate MKT fee is charged from seller for using affiliate MKT services<br /></p>
	WalletType         *string `json:"wallet_type,omitempty"`          // [Optional] This field indicates the wallet type.
}
type GetWalletTransactionListResponse struct {
	BaseResponse                                      // Common response fields
	Response     GetWalletTransactionListResponseData `json:"response"` // Response data
}
type GetWalletTransactionListResponseData struct {
	TransactionList []Transaction `json:"transaction_list"` // [Required]
	More            bool          `json:"more"`             // [Required]
}
type IncomeDetail struct {
	NextPage             *Pagination           `json:"next_page"`               // [Required] <p>Contains pagination metadata for fetching the next page.</p>
	IncomeDetailListItem *IncomeDetailListItem `json:"income_detail_list_item"` // [Required] <p>List of income detail objects</p>
}
type IncomeDetailListItem struct {
	PaymentMethod         string  `json:"payment_method"`          // [Required] <p>Payment channel or method used for the order</p>
	OrderSn               string  `json:"order_sn"`                // [Required] <p>Unique order serial number associated with the income record.</p>
	Description           string  `json:"description"`             // [Required] <p>Type of income or billing item — e.g., Order Income, Adjustment etc</p>
	Status                string  `json:"status"`                  // [Required] <p>Status description of the order income or payout.</p>
	Currency              string  `json:"currency"`                // [Required] <p>Currency in which the income was transacted.</p>
	EstimatedEscrowAmount float64 `json:"estimated_escrow_amount"` // [Required] <p>Estimated escrow amount pending release for the order.</p>
	EstimatedPayoutTime   int64   `json:"estimated_payout_time"`   // [Required] <p>Estimated payout time (Unix timestamp). Applicable for Pending/To Release status.</p>
	ToReleaseAmount       float64 `json:"to_release_amount"`       // [Required] <p>Amount that is queued for release to seller (Cross Border only).</p>
	CreationDate          int64   `json:"creation_date"`           // [Required] <p>Order creation timestamp (Unix format).</p>
	ReleasedAmount        float64 `json:"released_amount"`         // [Required] <p>Amount successfully released to the seller.</p>
	ActualPayoutTime      int64   `json:"actual_payout_time"`      // [Required] <p>Actual payout time (Unix timestamp) when funds were transferred.</p>
}
type ItemInstallment struct {
	ItemId     int64   `json:"item_id"`     // [Required] Item unique id
	TenureList []int64 `json:"tenure_list"` // [Required] The tenures of item support installment. [] represents with no installment
}
type ItemPlanAhora struct {
	ItemId               int64 `json:"item_id"`                // [Required] Only applicable for local AR sellers.
	ParticipatePlanAhora bool  `json:"participate_plan_ahora"` // [Required] Only applicable for local AR sellers.
}
type ItemsPromotion struct {
	PromotionType string      `json:"promotion_type"` // [Required] <p>Indicates the type of item- or package-level promotion applied to a product. Each item can be associated with at most one item promotion and one package promotion at a time.<br /><br />Item Promotions:</p><p>low_price_promotion</p><p>deep_discount</p><p>platform_sale</p><p>seller_discount</p><p>flash_sale</p><p>wholesale</p><p>welcome_package_free_gift</p><p>brand_flash_sale</p><p>in_shop_flash_sale</p><p>synced_promo</p><p>platform_streaming_price</p><p>seller_streaming_price</p><p>exclusive_streamer_price</p><p>price_bidding_with_rebate</p><p>price_bidding_without_rebate</p><p>seller_advisor_price</p><p>selling_price</p><p>settlement_price</p><p>campaign_settlement_price</p><p>local_sip_settlement_price</p><p>platform_exclusive_price</p><p>seller_exclusive_price</p><p>seller_member_exclusive_sku</p><p>item_price</p><p>order_sync_price</p><p><br />Package Promotions:<br />bundle_deal<br />add_on_deal_main<br />add_on_deal_sub<br /><br /><br /><br /><br /></p>
	PromotionId   interface{} `json:"promotion_id"`   // [Required] <p>Represents the unique identifier of a specific promotion applied to an item. Each promotion_id corresponds to a distinct promotion rule or campaign, defined under a particular promotion_type.&nbsp;</p>
}
type KitItems struct {
	OriginalProductId float64 `json:"original_product_id"` // [Required] <p>The merchant item identifier of the product within the kit&nbsp;(Only for BR Local Sellers)</p>
	OriginalModelId   float64 `json:"original_model_id"`   // [Required] <p>The merchant product model of the item within the kit&nbsp;(Only for BR Local Sellers)</p>
	TotalQty          float64 `json:"total_qty"`           // [Required] <p>The quantity of this specific component within the kit.&nbsp;(Only for BR Local Sellers)</p>
	OriginalPrice     float64 `json:"original_price"`      // [Required] <p>The price of the item when it is listed as a standalone item.</p>
	ProportionalPrice float64 `json:"proportional_price"`  // [Required] <p>The price of the item when it is listed within the kit (i.e. proportionally distributed)&nbsp;(Only for BR Local Sellers)</p>
}
type NetCommissionFeeInfo struct {
	RuleId          float64 `json:"rule_id"`           // [Required] <p>The unique identifier of the commission rule applied to calculate the net commission fee.</p>
	FeeAmount       float64 `json:"fee_amount"`        // [Required] <p>The net commission fee amount calculated based on the corresponding commission rule.</p>
	RuleDisplayName string  `json:"rule_display_name"` // [Required] <p>The display name of the commission rule for reference and readability.</p>
}
type NetServiceFeeInfo struct {
	RuleId          float64 `json:"rule_id"`           // [Required] <p>The unique identifier of the service fee rule applied to calculate the net service fee.</p>
	FeeAmount       float64 `json:"fee_amount"`        // [Required] <p>The net service fee amount calculated based on the corresponding service fee rule.</p>
	RuleDisplayName string  `json:"rule_display_name"` // [Required] <p>The display name of the service fee rule for reference and readability.</p>
	Category        string  `json:"category"`          // [Required] <p>The category of the service fee, indicating the type of service the fee is charged for.</p>
}
type OfflineAdjustment struct {
	AdjustmentAmount float64 `json:"adjustment_amount"` // [Required] The amount of offline adjustments.
	Module           string  `json:"module"`            // [Required] The reason for offline adjustment.
	Remark           string  `json:"remark"`            // [Required] The remark for the reason.
	Scenario         string  `json:"scenario"`          // [Required] The scenario of adjustment.
	AdjustmentLevel  string  `json:"adjustment_level"`  // [Required] Dimension of offline adjustment. Available value: shop, order.
	OrderSn          string  `json:"order_sn"`          // [Required] Shopee's unique identifier for an order.
}
type OrderAdjustment struct {
	Amount           float64 `json:"amount"`            // [Required] <p>adjustment transaction amount.<br /></p>
	Date             int64   `json:"date"`              // [Required] <p>adjustment transaction complete date.<br /></p>
	Currency         string  `json:"currency"`          // [Required] <p>order level adjustment transaction's currency type.<br /></p>
	AdjustmentReason string  `json:"adjustment_reason"` // [Required] <p>Reason for adjustment.&nbsp;</p><p>Possible cases could be:</p><p>- Return Refund deduction or compensation</p><p>- logistic issue deduction or compensation</p><p>- marketing fee deduction</p><p>- payment related fee<br /></p>
}
type OrderIncome struct {
	EscrowAmount                                       float64               `json:"escrow_amount"`                                             // [Required] <p>The total amount that the seller is expected to receive for the order and will change before order is completed.&nbsp;</p><p>For non cb sip affiliate shop (new formula):&nbsp;</p><p>escrow_amount=&nbsp;original_cost_of_goods_sold-original_shopee_discount+seller_return_refund+ shopee_discount- voucher_from_seller- seller_coin_cash_back+ buyer_paid_shipping_fee- actual_shipping_fee+ shopee_shipping_rebate+ shipping_fee_discount_from_3pl- reverse_shipping_fee+ rsf_seller_protection_fee_claim_amount- final_return_to_seller_shipping_fee- seller_transaction_fee- service_fee- commission_fee- campaign_fee- shipping_seller_protection_fee_amount- delivery_seller_protection_fee_premium_amount-final_escrow_product_gst- order_ams_commission_fee- escrow_tax-sales_tax_on_lvg-reverse_shipping_fee_sst-shipping_fee_sst-withholding_tax-overseas_return_service_fee-vat_on_imported_goods - withholding_vat_tax - withholding_pit_tax - withholding_cit_tax - seller_order_processing_fee&nbsp;+ buyer_paid_packaging_fee - trade_in_bonus_by_seller - fbs_fee -&nbsp;ads_escrow_top_up_fee_or_technical_support_fee - th_import_duty<br /></p><p><br /></p><p>For cb sip affiliate shop:&nbsp;escrow_amount=escrow_amount_pri * exchange_rate</p><p><br /></p><p>note: Return refund amount = if adjustable RR, will equal to drc_adjustable_refund<br /></p>
	BuyerTotalAmount                                   float64               `json:"buyer_total_amount"`                                        // [Required] <p>The snapshot value of total amount that paid by buyer at checkout moment.<br /></p>
	OrderOriginalPrice                                 float64               `json:"order_original_price"`                                      // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPrice                                      float64               `json:"original_price"`                                            // [Required] The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.
	OrderDiscountedPrice                               float64               `json:"order_discounted_price"`                                    // [Required] <p>The total discounted price for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderSellingPrice                                  float64               `json:"order_selling_price"`                                       // [Required] <p>This field value = sum of item unit price.selling price comes from the sum up of each item's unit price. For non-bundle deal case, this value will be same like order_original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo.It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	OrderSellerDiscount                                float64               `json:"order_seller_discount"`                                     // [Required] <p>The total discount seller provided for this order. It returns the subtotal of that specific item if quantity exceeds 1. (Only display for non cb sip affiliate shop. )<br /></p>
	BcrsDeposit                                        float64               `json:"bcrs_deposit"`                                              // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA). This value might be updated after RR/cancellation.</p>
	SellerDiscount                                     float64               `json:"seller_discount"`                                           // [Required] Final sum of each item seller discount of a specific order. (Only display for non cb sip affiliate shop. )
	ShopeeDiscount                                     float64               `json:"shopee_discount"`                                           // [Required] <p>Final sum of each item Shopee discount of a specific order. This amount will Return remaining rebate value to seller. (Only display for non cb sip affiliate order. )</p>
	VoucherFromSeller                                  float64               `json:"voucher_from_seller"`                                       // [Required] Final value of voucher provided by Seller for the order. (Only display for non cb sip affiliate shop. )
	VoucherFromShopee                                  float64               `json:"voucher_from_shopee"`                                       // [Required] Final value of voucher provided by Shopee for the order. (Only display for non cb sip affiliate shop. )
	VoucherFromExternalParty                           float64               `json:"voucher_from_external_party"`                               // [Required] <p>Final value of voucher provided by External Party for the order. (Only display for non cb sip affiliate shop.)</p>
	Coins                                              float64               `json:"coins"`                                                     // [Required] This value indicates the total amount offset when the buyer consumed Shopee Coins upon checkout. (Only display for non cb sip affiliate shop. )
	BuyerPaidShippingFee                               float64               `json:"buyer_paid_shipping_fee"`                                   // [Required] The shipping fee paid by buyer. (Only display for non cb sip affiliate shop. )
	BuyerTransactionFee                                float64               `json:"buyer_transaction_fee"`                                     // [Required] Tansaction fee paid by buyer for the order. (Only display for non cb sip affiliate shop. )
	CrossBorderTax                                     float64               `json:"cross_border_tax"`                                          // [Required] <p>Amount incurred by Buyer for purchasing items outside of home country or region. Amount may change after Return Refund. (Only display for non cb sip affiliate shop. )</p>
	PaymentPromotion                                   float64               `json:"payment_promotion"`                                         // [Required] The amount offset via payment promotion. (Only display for non cb sip affiliate shop. )
	CommissionFee                                      float64               `json:"commission_fee"`                                            // [Required] <p>The commission fee charged by Shopee platform if applicable.&nbsp;</p><p><br /></p><p>For cb sip affiliate shop:&nbsp;commission_fee=commission_fee_pri * exchange_rate</p>
	ServiceFee                                         float64               `json:"service_fee"`                                               // [Required] <p>Amount charged by Shopee to seller for additional services.</p><p><br /></p><p>For cb sip affiliate shop:&nbsp;service_fee=service_fee_pri * exchange_rate</p><p><br /></p><p>For tw shop, there will be pre-order service fee included</p>
	SellerTransactionFee                               float64               `json:"seller_transaction_fee"`                                    // [Required] Tansaction fee paid by seller for the order. (Only display for non cb sip affiliate shop. )
	SellerLostCompensation                             float64               `json:"seller_lost_compensation"`                                  // [Required] Compensation to seller in case of lost parcel. (Only display for non cb sip affiliate shop. )
	SellerCoinCashBack                                 float64               `json:"seller_coin_cash_back"`                                     // [Required] Value of coins provided by Seller for purchasing with his or her store for the order. (Only display for non cb sip affiliate shop. )
	EscrowTax                                          float64               `json:"escrow_tax"`                                                // [Required] Cross-border tax imposed by the Indonesian government on sellers. (Only display for non cb sip affiliate shop. )
	EstimatedShippingFee                               float64               `json:"estimated_shipping_fee"`                                    // [Required] The estimated shipping fee is an estimation calculated by Shopee based on specific logistics courier's standard. (Only display for non cb sip affiliate shop. )
	FinalShippingFee                                   float64               `json:"final_shipping_fee"`                                        // [Required] Final adjusted amount that seller has to bear as part of escrow. This amount could be negative or positive. (Only display for non cb sip affiliate shop. )
	ActualShippingFee                                  float64               `json:"actual_shipping_fee"`                                       // [Required] The final shipping cost of order and it is positive. For Non-integrated logistics channel is 0. (Only display for non cb sip affiliate shop. )
	ShippingFeeSst                                     float64               `json:"shipping_fee_sst"`                                          // [Required] <p>The Service Tax charged on Seller Paid Shipping Fee for forward shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. Definition of Seller Paid Shipping Fee is Actual Shipping Fee subtracted by sum of Shipping Fee Paid by Buyer and Shipping Rebate From Shopee. (Only applicable for non cb sip affiliate shop)<br /></p>
	OrderChargeableWeight                              int64                 `json:"order_chargeable_weight"`                                   // [Required] <p>For CB shop, display weight used to calculate actual_shipping_fee for this order.<br /></p>
	ShopeeShippingRebate                               float64               `json:"shopee_shipping_rebate"`                                    // [Required] The platform shipping subsidy to the seller. (Only display for non cb sip affiliate shop. )
	ShippingFeeDiscountFrom3pl                         float64               `json:"shipping_fee_discount_from_3pl"`                            // [Required] The discount of shipping fee from 3PL. Currently only applicable to ID. (Only display for non cb sip affiliate shop. )
	SellerShippingDiscount                             float64               `json:"seller_shipping_discount"`                                  // [Required] The shipping discount defined by seller. (Only display for non cb sip affiliate shop. )
	SellerVoucherCode                                  []string              `json:"seller_voucher_code"`                                       // [Required] The list of voucher code provided by seller. (Only display for non cb sip affiliate shop. )
	DrcAdjustableRefund                                float64               `json:"drc_adjustable_refund"`                                     // [Required] The adjustable refund amount from Shopee Dispute Resolution Center.
	CostOfGoodsSold                                    float64               `json:"cost_of_goods_sold"`                                        // [Required] Final amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
	OriginalCostOfGoodsSold                            float64               `json:"original_cost_of_goods_sold"`                               // [Required] Amount paid by the buyer for the items in a specific order. (Only display for non cb sip affiliate shop. )
	OriginalShopeeDiscount                             float64               `json:"original_shopee_discount"`                                  // [Required] <p>Sum of each item Shopee discount of a specific order. This amount will return initial rebate value (i.e. remaining Shopee Item Rebate +&nbsp;remaining Shopee PIX Rebate) to seller. (Only display for non cb sip affiliate order. )</p>
	SellerReturnRefund                                 float64               `json:"seller_return_refund"`                                      // [Required] Amount returned to Seller in the event of Partial Return.
	Items                                              []OrderIncomeItems    `json:"items"`                                                     // [Required] This object contains the detailed breakdown for all the items in this order, including regular items(non-activity) and activity items.
	EscrowAmountPri                                    float64               `json:"escrow_amount_pri"`                                         // [Required] <p>The total amount in the prim currency that the seller is expected to receive for the order and will change before order completed . escrow_amount_pri=original_price_pri-seller_return_refund_pri-commission_fee_pri-service_fee_pri-drc_adjustable_refund_pri.(Only display for cb sip affiliate order. )</p>
	BuyerTotalAmountPri                                float64               `json:"buyer_total_amount_pri"`                                    // [Required] <p>The total amount that paid by buyer in the primary currency. (Only display for cb sip affiliate order. )</p>
	OriginalPricePri                                   float64               `json:"original_price_pri"`                                        // [Required] <p>The original price of the item before ANY promotion/discount in the primary currency. It returns the subtotal of that specific item if quantity exceeds 1.(Only display for cb sip affiliate order. )</p>
	SellerReturnRefundPri                              float64               `json:"seller_return_refund_pri"`                                  // [Required] <p>Amount returned to Seller in the event of Partial Return in the primary currency. (Only display for cb sip affiliate shop. )</p>
	CommissionFeePri                                   float64               `json:"commission_fee_pri"`                                        // [Required] <p>The commission fee charged by Shopee platform if applicable in the primary currency. (Only display for cb sip affiliate shop. )</p>
	ServiceFeePri                                      float64               `json:"service_fee_pri"`                                           // [Required] <p>Amount charged by Shopee to seller for additional services in the primary currency. (Only display for cb sip affiliate shop. )</p>
	DrcAdjustableRefundPri                             float64               `json:"drc_adjustable_refund_pri"`                                 // [Required] <p>The adjustable refund amount from Shopee Dispute Resolution Center in the primary currency&nbsp;after proration. (Only applicable for cb sip affiliate shop.)</p>
	PriCurrency                                        string                `json:"pri_currency"`                                              // [Required] <p>The currency of the country or region where the shop that real seller operates. (Only display for cb sip affiliate shop. )</p>
	AffCurrency                                        string                `json:"aff_currency"`                                              // [Required] <p>The currency of the country or region where shop opened in. (Only display for cb sip affiliate shop. )</p>
	ExchangeRate                                       float64               `json:"exchange_rate"`                                             // [Required] Exchange rate from primary shop currency to affiliate shop currency.
	ReverseShippingFee                                 float64               `json:"reverse_shipping_fee"`                                      // [Required] Shopee charges the reverse shipping fee for the returned order.The value of this field will be non-negative.
	ReverseShippingFeeSst                              float64               `json:"reverse_shipping_fee_sst"`                                  // [Required] <p>The Service Tax charged on Reverse Shipping Fee for reverse shipping, based on Malaysia SST regulations&nbsp;for shipping fees on all orders. (Only applicable for non cb sip affiliate shop)<br /></p>
	FinalProductProtection                             float64               `json:"final_product_protection"`                                  // [Required] <p>The total amount of product protection purchased during placing an order.&nbsp;</p>
	CreditCardPromotion                                float64               `json:"credit_card_promotion"`                                     // [Required] This value indicate the offset via credit card promotion.
	CreditCardTransactionFee                           float64               `json:"credit_card_transaction_fee"`                               // [Required] <p>This value indicate the total transaction fee.</p><p>credit_card_transaction_fee=buyer_transaction_fee+seller_transaction_fee<br /></p>
	FinalProductVatTax                                 float64               `json:"final_product_vat_tax"`                                     // [Required] Value-added Tax is required for online purchases based on EU Value-added Tax regulations . (Only display for non cb sip affiliate shop. )
	FinalShippingVatTax                                float64               `json:"final_shipping_vat_tax"`                                    // [Required] <p>Value-added Tax for product price is required for online purchases based on EU Value-added Tax regulations. (Only applicable for non cb sip affiliate shop. )<br /></p>
	CampaignFee                                        float64               `json:"campaign_fee"`                                              // [Required] <p>The campaign fee charged by Shopee platform. Only available for some local Indonesian shops.<br /></p>
	SipSubsidy                                         float64               `json:"sip_subsidy"`                                               // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. (Only available for CB SIP A Shops)<br /></p>
	SipSubsidyPri                                      float64               `json:"sip_subsidy_pri"`                                           // [Required] <p>The SIP subsidy amount is the difference between the item settlement price set by seller and item price actually paid by buyer. This value is in the primary currency. (Only available for CB SIP A Shops)<br /></p>
	RsfSellerProtectionFeeClaimAmount                  float64               `json:"rsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The insurance claim amount if seller opt in to insurance program. this covers Reverse shipping Fee in the case of RR. As per Jun 2024:<br />- For ID &amp; MY Local: After Extension on coverage to FSF due to RR. this claim amount will consist of FSF + RSF claim.<br />- For PH local: This will only cover RSF claim<br /><br />will be updated if there's any RR/cancellation<br /></p>
	ShippingSellerProtectionFeeAmount                  float64               `json:"shipping_seller_protection_fee_amount"`                     // [Required] <p>Service fee charged to seller in MY,ID,PH,BR Local (as per Jun 2024) due to additional program purchased.<br /></p>
	FinalEscrowProductGst                              float64               `json:"final_escrow_product_gst"`                                  // [Required] <p>Goods and Service Tax for product price is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore)<br /></p>
	FinalEscrowShippingGst                             float64               `json:"final_escrow_shipping_gst"`                                 // [Required] <p>Goods and Service Tax for shipping fee is required for imported goods (overseas orders) based on Singapore GST regulations. (Only applicable for non cb sip affiliate shop selling in Singapore.)<br /></p>
	DeliverySellerProtectionFeePremiumAmount           float64               `json:"delivery_seller_protection_fee_premium_amount"`             // [Required] <p>[Currently apply to ID &amp; local orders only] An insurance premium charged to seller at the time parcel is picked up by 3PL for insurance in case of parcel lost/damaged by 3PL.<br /></p>
	OrderAdjustment                                    []OrderAdjustment     `json:"order_adjustment"`                                          // [Required] <p>Order level adjustment transaction information.<br /></p><p>If the&nbsp;order without adjustment, no returned of the field.</p>
	TotalAdjustmentAmount                              float64               `json:"total_adjustment_amount"`                                   // [Required] <p>Total adjustment made to the order.<br /></p>
	EscrowAmountAfterAdjustment                        float64               `json:"escrow_amount_after_adjustment"`                            // [Required] <p>Final income seller can get from this order after deduct the order-level adjustment.<br /></p>
	OrderAmsCommissionFee                              float64               `json:"order_ams_commission_fee"`                                  // [Required] <p>The amount of affiliate commission fee for this order. Applicable for orders sold via the Affiliate Program.<br /></p>
	BuyerPaymentMethod                                 string                `json:"buyer_payment_method"`                                      // [Required] <p>The payment method buyer used when do the order checkout.<br /></p>
	InstalmentPlan                                     string                `json:"instalment_plan"`                                           // [Required] <p>The instalment plan buyer chosen when do the order checkout. Only applicable when payment method support instalment.<br /></p>
	SalesTaxOnLvg                                      float64               `json:"sales_tax_on_lvg"`                                          // [Required] <p>Sales Tax on Low Value Goods (LVG) is required for imported goods (overseas orders) based on Malaysia SST regulations&nbsp;for selective orders (e.g. CB LVG orders in MY)<br /></p>
	FinalReturnToSellerShippingFee                     float64               `json:"final_return_to_seller_shipping_fee"`                       // [Required] <p>The amount of fee charged to seller (if any) for the failed delivery order.</p>
	WithholdingTax                                     float64               `json:"withholding_tax"`                                           // [Required] <p>Only for PH and ID local shops.</p><p><br /></p><p><b>PH</b>: According to regulations issued by Bureau of Internal Revenue in PH, the Withholding Tax is applied to the gross remittances sent by Shopee to online suppliers of goods and services.<br /><br /><b>ID</b>: According to regulations issued by Directorate General of Taxation in ID, the Withholding Tax is applied to the income stated in the invoice generated via Shopee related to Seller and/or Merchants' sales in Shopee's platform.</p>
	OverseasReturnServiceFee                           float64               `json:"overseas_return_service_fee"`                               // [Required] <p>This is overseas return service fee charged to sellers who register ORS program.(Only applicable for non cb sip affiliate shop)<br /></p>
	ProratedCoinsValueOffsetReturnItems                float64               `json:"prorated_coins_value_offset_return_items"`                  // [Required] <p>This is the prorated value from cash equivalent of coin offset due to adjustable RR.This field will only be updated when there is an adjustable RR. If it's a full RR or normal order will response 0.<br /></p>
	ProratedShopeeVoucherOffsetReturnItems             float64               `json:"prorated_shopee_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from shopee voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedSellerVoucherOffsetReturnItems             float64               `json:"prorated_seller_voucher_offset_return_items"`               // [Required] <p>This is the prorated refund value from seller voucher discount due to adjustable RR.This field will only be updated when there is an adjustable RR.&nbsp;If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoBankOffsetReturnItems   float64               `json:"prorated_payment_channel_promo_bank_offset_return_items"`   // [Required] <p>This is the prorated value from bank payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	ProratedPaymentChannelPromoShopeeOffsetReturnItems float64               `json:"prorated_payment_channel_promo_shopee_offset_return_items"` // [Required] <p>This is the prorated value from shopee payment channel promo due to adjustable RR.This field will only be updated when there is an adjustable RR.If it's a full RR or normal order will response 0.<br /></p>
	FsfSellerProtectionFeeClaimAmount                  float64               `json:"fsf_seller_protection_fee_claim_amount"`                    // [Required] <p>The claim amount given to seller if seller opt in to shipping fee service program. this covers Forward Shipping Fee cost in the case of cancelled due to Failed delivery.&nbsp;<br /></p>
	VatOnImportedGoods                                 float64               `json:"vat_on_imported_goods"`                                     // [Required] <p>7% VAT is charged for imported goods entering Thailand.<br /></p><p>8% VAT is charged for imported goods entering Vietnam</p>
	TenureInfoList                                     *TenureInfo           `json:"tenure_info_list"`                                          // [Required]
	WithholdingVatTax                                  float64               `json:"withholding_vat_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold VAT tax applicable to all VN business household and VN individual sellers</p>
	WithholdingPitTax                                  float64               `json:"withholding_pit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to all VN business household and VN individual sellers</p>
	WithholdingCitTax                                  float64               `json:"withholding_cit_tax"`                                       // [Required] <p>By VN law, E-commerce platforms need to Withhold Personal Income Tax applicable to corporate sellers selling in VN region</p>
	TaxRegistrationCode                                string                `json:"tax_registration_code"`                                     // [Required] <p>For VN Withholding Tax. This is the Tax Registration Number (TRN) from Seller Info (Business information) of the seller at the point of order creation</p>
	SellerOrderProcessingFee                           float64               `json:"seller_order_processing_fee"`                               // [Required] <p>Order Processing Fee is the amount charged to sellers for every order created.</p>
	BuyerPaidPackagingFee                              float64               `json:"buyer_paid_packaging_fee"`                                  // [Required] <p>The fee charged to the buyer for packaging materials</p>
	TradeInBonusBySeller                               float64               `json:"trade_in_bonus_by_seller"`                                  // [Required] <p>The discount provided by Seller/ Retailers for buyers who opt for trade-in.</p>
	FbsFee                                             float64               `json:"fbs_fee"`                                                   // [Required] <p>Fulfilled by Shopee (FBS) Fee applied to this order, covering costs such as handling and storage and packaging. Only for PH Local Orders.</p>
	NetCommissionFee                                   float64               `json:"net_commission_fee"`                                        // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net commission fee applied to the order, calculated as the sum of all net commission fee items.</p><p>-only for BR local sellers.</p><br />
	NetServiceFee                                      float64               `json:"net_service_fee"`                                           // [Required] <p>The respective fee amounts after the proportional rebate deduction.The total net service fee applied to the order, calculated as the sum of all net service fee items.</p><p>-only for BR local sellers.</p>
	NetCommissionFeeInfoList                           *NetCommissionFeeInfo `json:"net_commission_fee_info_list"`                              // [Required] <p>Returns a breakdown of the net commission fees.</p><p>-only for BR local sellers.</p>
	NetServiceFeeInfoList                              *NetServiceFeeInfo    `json:"net_service_fee_info_list"`                                 // [Required] <p>Returns a breakdown of the net service fees.<br /></p><p>-only for BR local sellers.</p>
	SellerProductRebate                                *SellerProductRebate  `json:"seller_product_rebate"`                                     // [Required] <p>The shopee rebate borne by seller.</p><p>-only for BR local sellers.</p>
	PixDiscount                                        float64               `json:"pix_discount"`                                              // [Required] <p>[BR only]Final sum of pix discount of a specific order. (Only display for non cb sip affiliate shop.)</p>
	ProratedPixDiscountOffsetReturnItems               float64               `json:"prorated_pix_discount_offset_return_items"`                 // [Required] <p>[BR only]This is the prorated value from pix discount due to adjustable RR. This field will only be updated when there is an adjustable RR. If it's a full RR or normal order, will response 0.</p>
	AdsEscrowTopUpFeeOrTechnicalSupportFee             float64               `json:"ads_escrow_top_up_fee_or_technical_support_fee"`            // [Required] <p>Includes both ads escrow top up fee (auto escrow top up to your ads balance) and technical support fee (charged by Shopee)</p><p>The actual fee type included in this field varies depending on the seller type and selling region, and may represent one of the following in Shopee Seller Center:</p>   <p>Ads Escrow Top-Up Fee</p>   <p>For local MY TH SG VN PH ID sellers and CNCB / JPCB / KRCB sellers selling in PH and ID</p>   <p>For JPCB sellers selling in SG, MY, TH, and VN</p>     <p>Technical Support Fee</p>   <p>For CNCB sellers selling in SG, MY, TH, and VN</p>     <p>Traffic Growth Fee</p>   <p>For KRCB sellers selling in SG, MY, TH, and VN</p>
	ThImportDuty                                       float64               `json:"th_import_duty"`                                            // [Required] <p>[TH only] Import Duty collected for imported goods entering Thailand.</p>
	RemainingVoucher                                   float64               `json:"remaining_voucher"`                                         // [Required] <p>[Only for BR local shop]Represents the portion of Shopee voucher that is not consumed after fee offset.</p>
}
type OrderIncomeItems struct {
	ItemId                    int64            `json:"item_id"`                      // [Required] <p>ID of item<br /></p>
	ItemName                  string           `json:"item_name"`                    // [Required] <p>Name of item<br /></p>
	ItemSku                   string           `json:"item_sku"`                     // [Required] <p>A item SKU (stock keeping unit) is an identifier defined by a seller, sometimes called parent SKU. Item SKU can be assigned to an item in Shopee Listings.<br /></p>
	ModelId                   int64            `json:"model_id"`                     // [Required] <p>ID of the model that belongs to the same item.<br /></p>
	ModelName                 string           `json:"model_name"`                   // [Required] <p>Name of the model that belongs to the same item. A seller can offer variations of the same item. For example, the seller could create a fixed-priced listing for a t-shirt design and offer the shirt in different colors and sizes. In this case, each color and size combination is a separate variation. Each variation can have a different quantity and price.<br /></p>
	ModelSku                  string           `json:"model_sku"`                    // [Required] <p>A model SKU (stock keeping unit) is an identifier defined by a seller. It is only intended for the seller's use. Many sellers assign a SKU to an item of a specific type, size, and color, which are variations of one item in Shopee Listings.<br /></p>
	LineItemId                int64            `json:"line_item_id"`                 // [Required] <p>The identity of order item. In case the order item is a bundle deal, this value will be unique to distinguish the order item.</p>
	OriginalPrice             float64          `json:"original_price"`               // [Required] <p>The original price of the item before ANY promotion/discount in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1.<br /></p>
	OriginalPricePri          float64          `json:"original_price_pri"`           // [Required] <p>The agreed settlement price of items used as settlement basis, amount is in the primary currency. (Only display for CB SIP affiliate shop.)<br /></p>
	SellingPrice              float64          `json:"selling_price"`                // [Required] <p>For non-bundle deal case, this value will be same like item original_price; For bundle deal case, this value will be price of sum of item price before bundle deal promo but after item promo. It returns the subtotal of that specific item if quantity exceeds 1 (Only display for non cb sip affiliate shop.)<br /></p>
	DiscountedPrice           float64          `json:"discounted_price"`             // [Required] <p>The after-discount price of the item in the listing currency. It returns the subtotal of that specific item if quantity exceeds 1. If there is no discount, this value will be the same as that of original_price.<br /></p>
	BcrsDeposit               float64          `json:"bcrs_deposit"`                 // [Required] <p>The deposit fee paid by buyer of $0.10 per container as part of the SG Beverage Container Return Scheme mandated by the National Environment Agency (NEA).&nbsp;This will be an initial value and will not update after RR/cancellation.</p>
	SellerDiscount            float64          `json:"seller_discount"`              // [Required] <p>The discount provided by seller for this item<br /></p>
	ShopeeDiscount            float64          `json:"shopee_discount"`              // [Required] <p>The discount provided by Shopee for this item<br /></p>
	DiscountFromCoin          float64          `json:"discount_from_coin"`           // [Required] <p>The offset of this item when the buyer consumed Shopee Coins upon checkout. In case of bundle deal item, this value will return 0. Due to technical restriction, this field will return incorrect value under bundle deal case if we don’t configure it to 0.<br /></p>
	DiscountFromVoucherShopee float64          `json:"discount_from_voucher_shopee"` // [Required] <p>The offset of this item when the buyer use Shopee voucher.&nbsp;</p>
	DiscountFromVoucherSeller float64          `json:"discount_from_voucher_seller"` // [Required] <p>The offset of this item when the buyer use seller-specific voucher.</p>
	ActivityType              string           `json:"activity_type"`                // [Required] <p>The type of the item, default is "". If the item is a bundle item the value is "bundle_deal", and if a add on deal item, the value is "add_on_deal"<br /></p>
	ActivityId                interface{}      `json:"activity_id"`                  // [Required] <p>If bundle_deal the is id of bundle deal, if add_on_deal this is id of add on deal.</p>
	IsMainItem                bool             `json:"is_main_item"`                 // [Required] <p>Meaning a main or sub item for add_on_deal.<br /></p>
	QuantityPurchased         int64            `json:"quantity_purchased"`           // [Required] <p>This value indicates the number of identical items purchased at the same time by the same buyer from one listing/item.<br /></p>
	IsB2cShopItem             bool             `json:"is_b2c_shop_item"`             // [Required] <p>Indicates whether this is a B2C owned item.<br /></p>
	AmsCommissionFee          float64          `json:"ams_commission_fee"`           // [Required] <p>The amount of affiliate commission fee. Applicable for items sold via the Affiliate Program.<br /></p>
	IsKit                     bool             `json:"is_kit"`                       // [Required] <p>Indicates if this item is a kit. (Only for BR Local Sellers)</p>
	KitItems                  *KitItems        `json:"kit_items"`                    // [Required] <p>Only applicable to BR local sellers</p>
	PromotionList             []ItemsPromotion `json:"promotion_list"`               // [Required]
}
type Pagination struct {
	Cursor   string `json:"cursor"`    // [Required] <p>Token to retrieve the next page of results. Returns empty if there is no more data.</p>
	PageSize int64  `json:"page_size"` // [Required] <p>Number of records returned per page.</p>
}
type PayOrder struct {
	OrderSn  string `json:"order_sn"`  // [Required] Shopee's unique identifier for an order.
	ShopName string `json:"shop_name"` // [Required] Name of the shop.
}
type Payout struct {
	PayoutInfo            *PayoutInfo         `json:"payout_info"`             // [Required] The information of payout.
	EscrowList            []PayoutEscrow      `json:"escrow_list"`             // [Required]
	OfflineAdjustmentList []OfflineAdjustment `json:"offline_adjustment_list"` // [Required] The list of offline adjustments.
}
type PayoutEscrow struct {
	EscrowAmount float64 `json:"escrow_amount"` // [Required] The total amount that the seller is expected to receive for the order.
	Currency     string  `json:"currency"`      // [Required] The currency used for calculating escrow amount.
	OrderSn      string  `json:"order_sn"`      // [Required] Shopee's unique identifier for an order.
}
type PayoutInfo struct {
	FromCurrency   string  `json:"from_currency"`   // [Required] The settlement currency of orders.
	PayoutCurrency string  `json:"payout_currency"` // [Required] The actual currency of payout.
	FromAmount     float64 `json:"from_amount"`     // [Required] The settlement amount.
	PayoutAmount   float64 `json:"payout_amount"`   // [Required] The actual amount of payout.
	ExchangeRate   string  `json:"exchange_rate"`   // [Required] The exchange rate.
	PayoutTime     int64   `json:"payout_time"`     // [Required] The time of payout.
	PayService     string  `json:"pay_service"`     // [Required] The service provider of seller. Available value: payoneer, pingpong, lianlian.
	PayeeId        string  `json:"payee_id"`        // [Required] Seller's account to receive the payout.
}
type ResponseDataItemPlanAhora struct {
	ItemId              int64 `json:"item_id"`               // [Required] Only applicable for local AR sellers.
	ParticipatePlanAhor bool  `json:"participate_plan_ahor"` // [Required] Only applicable for local AR sellers.
}
type ResponseDataPayout struct {
	FromCurrency      string  `json:"from_currency"`       // [Required] <p>The settlement currency of orders.<br /></p>
	PayoutCurrency    string  `json:"payout_currency"`     // [Required] <p>The actual currency of payout.<br /></p>
	FromAmount        float64 `json:"from_amount"`         // [Required] <p>The settlement amount.<br /></p>
	PayoutAmount      float64 `json:"payout_amount"`       // [Required] <p>The actual amount of payout.<br /></p>
	ExchangeRate      string  `json:"exchange_rate"`       // [Required] <p>The exchange rate.<br /></p>
	PayoutTime        int64   `json:"payout_time"`         // [Required] <p>The time of payout.<br /></p>
	PayService        string  `json:"pay_service"`         // [Required] <p>The service provider of seller. Available value: payoneer, pingpong, lianlian.<br /></p>
	PayeeId           string  `json:"payee_id"`            // [Required] <p>Seller's account to receive the payout.<br /></p>
	EncryptedPayoutId string  `json:"encrypted_payout_id"` // [Required] <p>payout id used to query API "v2.get_billing_item_info" as request parameters. User can get detailed billing items under current payout<br /></p>
}
type SellerProductRebate struct {
	Amount              float64 `json:"amount"`                // [Required] <p>This is the portion of Shopee rebate borne by seller.</p>
	CommissionFeeOffset float64 `json:"commission_fee_offset"` // [Required] <p>This is the offset to gross commission fee, reducing it to the net value.</p>
	ServiceFeeOffset    float64 `json:"service_fee_offset"`    // [Required] <p>This is the offset to gross service fee, reducing it to the net value.</p>
}
type SetItemInstallmentStatusRequest struct {
	ItemIdList           []int64 `json:"item_id_list"`                     // [Required] The id array of the item, Max :100
	ParticipatePlanAhora *bool   `json:"participate_plan_ahora,omitempty"` // [Optional] Only applicable and required for local AR sellers.
	TenureList           []int64 `json:"tenure_list"`                      // [Required] <p>Staged array, TH must be [3,6,10], TW region tenures must be in [3,6,12,24], [] means closed</p>
}
type SetItemInstallmentStatusResponse struct {
	BaseResponse                                      // Common response fields
	Response     SetItemInstallmentStatusResponseData `json:"response"` // Response data
}
type SetItemInstallmentStatusResponseData struct {
	ItemInstallmentList []ItemInstallment           `json:"item_installment_list"` // [Required]
	ItemPlanAhoraList   []ResponseDataItemPlanAhora `json:"item_plan_ahora_list"`  // [Required] Only applicable for local AR sellers.
}
type SetShopInstallmentStatusRequest struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required] <p>The status of installment contains 1 and 0.</p>
}
type SetShopInstallmentStatusResponse struct {
	BaseResponse                                      // Common response fields
	Response     SetShopInstallmentStatusResponseData `json:"response"` // Response data
}
type SetShopInstallmentStatusResponseData struct {
	InstallmentStatus int64 `json:"installment_status"` // [Required]
}
type TenureInfo struct {
	PaymentChannelName string `json:"payment_channel_name"` // [Required] <p>Name of the payment channel that buyer used in checkout.</p>
	InstalmentPlan     string `json:"instalment_plan"`      // [Required] <p>Tenure information. This will have value if payment channel used has tenure information, such as credit card.</p>
}
type TotalIncome struct {
	PendingAmount   float64 `json:"pending_amount"`    // [Required] <p>Total amount pending release (Local: orders before ESCROW_PAID; CB: orders before ESCROW_PAYOUT).</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ToReleaseAmount float64 `json:"to_release_amount"` // [Required] <p>Amount queued for release in the next payout cycle (CB only). Not applicable for Local shops.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
	ReleasedAmount  float64 `json:"released_amount"`   // [Required] <p>Total amount successfully released to the seller.</p><p>&lt;path&gt;&lt;/path&gt;<br /></p>
}
type Transaction struct {
	Status             string     `json:"status"`               // [Required] The status of the transaction，available values: FAILED,COMPLETED,PENDING,INITIAL.
	TransactionType    string     `json:"transaction_type"`     // [Required] The type of transaction.
	TxnTitle           string     `json:"txn_title"`            // [Required] <p>The transaction title sent by client (Adjustment Center) for adjustments, Only for ID local sellers for now.</p>
	Amount             float64    `json:"amount"`               // [Required] The amount of transaction.
	CurrentBalance     float64    `json:"current_balance"`      // [Required] The current balance of this account.
	CreateTime         int64      `json:"create_time"`          // [Required] The create time of the transaction.
	OrderSn            string     `json:"order_sn"`             // [Required] Shopee's unique identifier for an order.
	RefundSn           string     `json:"refund_sn"`            // [Required] The serial number of return.
	WithdrawalType     string     `json:"withdrawal_type"`      // [Required] The type of withdrawal.
	TransactionFee     float64    `json:"transaction_fee"`      // [Required] This field indicates the transaction fee.
	Description        string     `json:"description"`          // [Required] The detailed description of TOPUP SUCCESS and TOPUP FAILED.
	BuyerName          string     `json:"buyer_name"`           // [Required] The name of buyer.
	PayOrderList       []PayOrder `json:"pay_order_list"`       // [Required]
	ShopName           string     `json:"shop_name"`            // [Required] Name of the shop.
	WithdrawalId       int64      `json:"withdrawal_id"`        // [Required] <p>Withdrawal ID when transaction type is withdraw_created, withdrawal_completed, withdrawal_cancelled.</p>
	Reason             string     `json:"reason"`               // [Required] The reason for ADJUSTMENT_ADD and ADJUSTMENT_MINUS.
	RootWithdrawalId   int64      `json:"root_withdrawal_id"`   // [Required] Use this field to indicate the event where a withdrawal is split into several withdrawals due to the withdrawal limit.
	TransactionTabType string     `json:"transaction_tab_type"` // [Required] <p><b>Description:</b></p><p>A new response parameter added after:&nbsp;https://confluence.shopee.io/display/SPCT/%5BPRD%5D+%5BOpen+API%5D+Update+on+New+Open+API+to+fetch+Seller+wallet+Transaction&nbsp;<br /><br />This returns the updated transaction tab types that client can use to specify which transaction types they want to return. It will have the following tab types<br /><br />Default<br />wallet_order_income<br />wallet_adjustment_filter</p><p>wallet_wallet_payment</p><p>wallet_refund_from_order</p><p>wallet_withdrawals</p><p>fast_escrow_repayment</p><p>fast_pay</p><p>seller_loan<br />corporate_loan<br />pix_transactions_filter</p><p>open_finance_transactions_filter&nbsp;<br /><br />Note for BR, currently in SOP live configuration, wallet txn type that linked to&nbsp;pix_transactions_filter&nbsp; and&nbsp;open_finance_transactions_filter&nbsp;&nbsp;are classified as&nbsp;default&nbsp;&nbsp;type tab instead. therefore for Open API client who wants to query these 2 txn can put default as the filter in this type</p>
	MoneyFlow          string     `json:"money_flow"`           // [Required] <p>New response parameter provided after:&nbsp;https://confluence.shopee.io/display/SPCT/%5BPRD%5D+%5BOpen+API%5D+Update+on+New+Open+API+to+fetch+Seller+wallet+Transaction&nbsp;<br /><br />It's to indicate the money flow</p><p>MONEY_IN = addition&nbsp;<br />MONEY_OUT = deduction</p><p>if not specified in request, will return both</p><p>Note special case for TW JKO Pay, we will ignore Money_flow&nbsp;</p>
	OutletShopName     string     `json:"outlet_shop_name"`     // [Required] <p>The outlet shop name where this outlet transaction came from.&nbsp;(In the Original Instant Mart concept, outlet transactions are redirected to Mart.)</p>
}
type Transactions struct {
	Amount                   float64 `json:"amount"`                     // [Required] <p>each transaction's amount<br /></p>
	Currency                 string  `json:"currency"`                   // [Required] <p>transaction currency<br /></p>
	OrderSn                  string  `json:"order_sn"`                   // [Required] <p>transaction currency<br /></p>
	CostHeader               string  `json:"cost_header"`                // [Required] <p>transaction belongs to which type<br /></p>
	Scenario                 string  `json:"scenario"`                   // [Required] <p>transaction detailed scenarios<br /></p>
	Remark                   string  `json:"remark"`                     // [Required] <p>detailed description for this transactions<br /></p>
	Level                    string  `json:"level"`                      // [Required] <p>To define this transaction happen at order level or shop level. e.g. shop level adjustment<br /></p>
	BillingTransactionType   string  `json:"billing_transaction_type"`   // [Required] <p>could be Escrow (Order Income) or Adjustment (for this order)<br /></p>
	BillingTransactionStatus string  `json:"billing_transaction_status"` // [Required] <p>Will be either "To Release" or "Released".<br /></p>
}
