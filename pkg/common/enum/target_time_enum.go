package enum

type ProjectTimeType int

const (
	ONE_WEEKS ProjectTimeType = iota
	ONE_MONTH
	THREE_MONTH
	LONG_TERM
)

func (enum ProjectTimeType) Value() string {
	return [...]string{
		"ONE_WEEKS",
		"ONE_MONTH",
		"THREE_MONTH",
		"LONG_TERM",
	}[enum]
}

func GetProjectTimeType(value string) string {
	serviceMap := map[string]string{
		ONE_WEEKS.Value():   " 1 Minggu",
		ONE_MONTH.Value():   " 1 Bulan",
		THREE_MONTH.Value(): " 2 Sampai 3 Bulan",
		LONG_TERM.Value():   " Infinity",
	}

	return serviceMap[value]
}
