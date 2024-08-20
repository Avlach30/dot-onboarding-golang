package enum

type DbType int

const (
	MYSQL DbType = iota
	POSTGRES
)

func (enum DbType) Value() string {
	return [...]string{
		"MYSQL",
		"POSTGRES",
	}[enum]
}
