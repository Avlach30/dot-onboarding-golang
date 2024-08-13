package otp

type OtpProxy interface {
	// SendOTP contact can be email or phonNumber
	SendOTP(contact, code string) (resMsg interface{}, err error)
}
