package notifications

type NotificationProxy interface {
	Send(webhookURL, message string) (err error)
}
