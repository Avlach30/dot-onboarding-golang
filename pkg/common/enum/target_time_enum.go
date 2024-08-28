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
		ONE_WEEKS.Value():   "Deadline - 1 Minggu",
		ONE_MONTH.Value():   "Deadline - 1 Bulan",
		THREE_MONTH.Value(): "Deadline - 3 Bulan",
		LONG_TERM.Value():   "Deadline - Infinity",
	}

	return serviceMap[value]
}
