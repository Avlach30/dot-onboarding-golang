package authdomain

import "context"

type Usecase interface {
	ExchangeToken(ctx context.Context, phoneNumber, token string) error
}
