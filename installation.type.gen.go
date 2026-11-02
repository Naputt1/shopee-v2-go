package goshopee

type ReceiveInstallationOrderRequest struct {
	InstallationOrderId int64  `json:"installation_order_id"` // [Required] <p>Unique ID of the installation order. Passed to ISP on installation order creation.</p>
	Reason              string `json:"reason"`                // [Required] <p>Reason for the installation status update.<br />Currently only support empty string.</p>
}
type ReceiveInstallationOrderResponse struct {
	BaseResponse // Common response fields
}
type ValidateInstallationOtpCodeRequest struct {
	InstallationOrderId int64  `json:"installation_order_id"` // [Required] <p>Unique ID of the installation order. Passed to ISP on installation order creation.</p>
	InstallationOtp     string `json:"installation_otp"`      // [Required] <p>OTP Code from buyer side.</p>
}
type ValidateInstallationOtpCodeResponse struct {
	BaseResponse // Common response fields
}
