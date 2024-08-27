package enum

type ServiceType int

const (
	WEB_APPLICATION ServiceType = iota
	WEB_COMPRO
	MOBILE_APPS
	MAINTENANCE
)

func (enum ServiceType) Value() string {
	return [...]string{
		"WEB_APPLICATION",
		"WEB_COMPRO",
		"MOBILE_APPS",
		"MAINTENANCE",
	}[enum]
}

func GetTransformServiceType(value string) string {
	serviceMap := map[string]string{
		WEB_APPLICATION.Value(): "WEB APPLICATION",
		WEB_COMPRO.Value():      "WEB COMPRO",
		MOBILE_APPS.Value():     "MOBILE APPS",
		MAINTENANCE.Value():     "MAINTENANCE",
	}

	return serviceMap[value]
}
