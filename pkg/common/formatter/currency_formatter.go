package formatter

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CurrencyDotSeparator(lang language.Tag, currencyCode string, amount float64) string {
	p := message.NewPrinter(lang)

	return p.Sprintf("%s. %f", currencyCode, amount)
}
