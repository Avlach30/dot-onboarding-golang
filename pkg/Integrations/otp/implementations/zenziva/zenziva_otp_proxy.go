package zenziva

import (
	"encoding/json"
	"net/http"

	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/Integrations/otp"
	"github.com/codespace-id/codespace-x/pkg/Integrations/otp/dto"
	"github.com/pkg/errors"
)

type zenziva struct {
	baseURL string
	userkey string
	passkey string
}

func NewZenziva(
	baseURL string,
	userkey string,
	passkey string,
) otp.OtpProxy {
	return &zenziva{
		baseURL: baseURL,
		userkey: userkey,
		passkey: passkey,
	}
}

// SendOTP implements otp.OtpProxy.
func (z *zenziva) SendOTP(contact string, code string) (resMsg interface{}, err error) {
	httpMethod := http.MethodPost

	request := dto.OtpRequestSmsMasking{
		Userkey: config.ZenzivaUserKey,
		Passkey: config.ZenzivaPassKey,
		To:      contact,
		Message: "Kode OTP " + code + ", Rahasiakan OTP dari siapapun, OTP berlaku 5 menit. Regards codespace.id",
	}
	reqBody, _ := json.Marshal(request)

	resp, err := otp.CallAPI(httpMethod, z.baseURL, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "zenziva.SendOTP")
	}
	if resp == nil {
		return nil, errors.Wrap(errors.New("Response From Zenziva Service Nil"), "zenziva.SendOTP.CallAPI")
	}

	return resp, nil
}
