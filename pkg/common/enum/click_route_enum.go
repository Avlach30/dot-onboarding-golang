package enum

type ClickRouteType int

const (
	CLICKTOSCREEN ClickRouteType = iota
	CLICKTOURL
)

func (enum ClickRouteType) Value() string {
	return [...]string{
		"SCREEN",
		"URL",
	}[enum]
}
