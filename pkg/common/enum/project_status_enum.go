package enum

type ProjectStatusType int

const (
	INQUIRY ProjectStatusType = iota
	QUOTATION
	SPK
	WAITING_FOR_PAYMENT
	ON_DEVELOPMENT
	BAST_AND_GUARANTEE
	FINISHED
)

func (enum ProjectStatusType) Value() string {
	return [...]string{
		"INQUIRY",
		"QUOTATION",
		"SPK",
		"WAITING_FOR_PAYMENT",
		"ON_DEVELOPMENT",
		"BAST_AND_GUARANTEE",
		"FINISHED",
	}[enum]
}
