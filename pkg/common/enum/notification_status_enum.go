package enum

type NotificationStatusType int

const (
	READ NotificationStatusType = iota
	UNREAD
)

func (enum NotificationStatusType) Value() string {
	return [...]string{
		"READ",
		"UNREAD",
	}[enum]
}
