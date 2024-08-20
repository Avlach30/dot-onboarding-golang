package domain

type Entity struct {
	ID          uint
	UserID      int64
	Title       string
	Description string
	IsRead      bool
	RouteScreen string
}
