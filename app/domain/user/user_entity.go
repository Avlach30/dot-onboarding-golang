package userdomain

type Entity struct {
	ID             uint
	Fullname       string
	IdentityNumber string
	PhoneNumber    string
	Gender         string
	Password       string
	Role           string
	ImageURL       string
}
