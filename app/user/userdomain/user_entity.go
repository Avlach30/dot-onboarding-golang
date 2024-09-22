package userdomain

type Entity struct {
	ID             int64
	Fullname       string
	IdentityNumber string
	PhoneNumber    string
	Gender         string
	Password       string
	Role           string
	ImageURL       string
	Email          string

	Roles string
}
