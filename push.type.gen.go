package goshopee

type ConfirmConsumedLostPushMessageRequest struct {
	LastMessageId int64 `json:"last_message_id"` // [Required] <p>The last_message_id returned by v2.push.get_lost_push_message.<br /></p>
}
type ConfirmConsumedLostPushMessageResponse struct {
	BaseResponse // Common response fields
}
type GetAppPushConfigResponse struct {
	BaseResponse                              // Common response fields
	Response     GetAppPushConfigResponseData `json:"response"` // <p>Detail informations you are querying.<br /></p>
}
type GetAppPushConfigResponseData struct {
	CallbackUrl       string  `json:"callback_url"`         // [Required] <p>The callback url of push mechanism. It is the address where the Shopee will send the push message to. If you don't set any callback_url before, this parameters is required.<br /></p>
	LivePushStatus    string  `json:"live_push_status"`     // [Required] <p>live push status:Normal,Warning,Suspended<br /></p>
	SuspendedTime     int64   `json:"suspended_time"`       // [Required] <p>The live push suspended time caused by low successful rate of push mechanism.Only when live push status is suspended, this parameters will response.<br /></p>
	BlockedShopId     []int64 `json:"blocked_shop_id"`      // [Required] <p>Use this filed to indicate that Shopee won't send any push message created by this shop.<br /></p>
	PushConfigOnList  []int64 `json:"push_config_on_list"`  // [Required] <p>Use this field to indicate which push config turn on, and you can receive the push message.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
	PushConfigOffList []int64 `json:"push_config_off_list"` // [Required] <p>Use this field to indicate which push config turn on, and you can receive the push message.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
}
type GetLostPushMessageResponse struct {
	BaseResponse                                // Common response fields
	Response     GetLostPushMessageResponseData `json:"response"` //
}
type GetLostPushMessageResponseData struct {
	PushMessageList []PushMessage `json:"push_message_list"` // [Required] <p>Returns the earliest 100 lost push messages that were lost within 3 days of the current time and not confirmed to have been consumed.</p>
	HasNextPage     bool          `json:"has_next_page"`     // [Required] <p>This is to indicate whether the lost push message to be consumed is more than 100. If this value is true, you may need to continue calling to get the remaining lost push messages to be consumed.<br /></p>
	LastMessageId   int64         `json:"last_message_id"`   // [Required] <p>Specifies the end entry of data returned in the current call.<br /></p>
}
type PushMessage struct {
	ShopId    int64  `json:"shop_id"`   // [Required] <p>Shopee's unique identifier for a shop.&nbsp;</p><p>If it's a partner level push (such as code: 1, 2, 12), shop_id will not be returned.</p>
	Code      int64  `json:"code"`      // [Required] <p>Shopee's unique identifier for a push notification.<br /></p>
	Timestamp int64  `json:"timestamp"` // [Required] <p>Timestamp that indicates the message was lost.<br /></p>
	Data      string `json:"data"`      // [Required] <p>Main Push message data.<br /></p>
}
type SetAppPushConfigRequest struct {
	BlockedShopIdList []int64 `json:"blocked_shop_id_list,omitempty"` // [Optional] <p>Use this filed to set shops that need to be blocked.Please input no more than 500 shop id.<br /></p>
	CallbackUrl       *string `json:"callback_url,omitempty"`         // [Optional] <p>The callback url of push mechanism. It is the address where the Shopee will send the push message to. If you don't set any callback_url before, this parameters is required.<br /></p>
	SetPushConfigOff  []int64 `json:"set_push_config_off,omitempty"`  // [Optional] <p>Turn off Push config, Shopee won't send the push message into the callback url.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
	SetPushConfigOn   []int64 `json:"set_push_config_on,omitempty"`   // [Optional] <p>Turn on push config, Shopee will send the push message into the callback url.</p><p>1=Shop authorization for partners</p><p>2=Shop deauthorization for partners</p><p>3=Order status update push</p><p>4=TrackingNo push</p><p>5=Shopee Updates</p><p>6=Banned item push</p><p>7=item promotion push</p><p>8=reserved stock change push</p><p>9=promotionn update push</p><p>10=webchat push</p><p>11=video upload push</p><p>12=openapi authorization expiry push</p><p>13=brand register result<br /></p>
}
type SetAppPushConfigResponse struct {
	BaseResponse                              // Common response fields
	Response     SetAppPushConfigResponseData `json:"response"` // <p>Detail informations you are querying.<br /></p>
}
type SetAppPushConfigResponseData struct {
	Result string `json:"result"` // [Required] <p>Use this field to indicate whether the configuration is set successfully.<br /></p>
}
