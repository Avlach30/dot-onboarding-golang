package enum

type GenderType string

const (
	MALE    GenderType = "MALE"
	FEMALE  GenderType = "FEMALE"
	UNKNOWN GenderType = "UNKNOWN"
)

func (enum GenderType) Value() string {
	return string(enum)
}
