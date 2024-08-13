package authdomain

import "database/sql"

type OtpEntity struct {
	Code       string
	Identifier string
	Trial      int
	IsValid    int
	ExpiredAt  sql.NullTime
}
