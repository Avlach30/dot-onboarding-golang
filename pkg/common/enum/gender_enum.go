package enum

type GenderType int

const (
	MALE GenderType = iota
	FEMALE
	UNKNOWN
)

func (enum GenderType) Value() string {
	return [...]string{
		"MALE",
		"FEMALE",
		"UNKNOWN",
	}[enum]
}
