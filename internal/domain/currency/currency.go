package currency

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type Currency string

const (
	KGS Currency = "KGS"
	KZT Currency = "KZT"
	USD Currency = "USD"
)

const DefaultCurrency = KGS

func (c Currency) String() string {
	return string(c)
}

func (c Currency) IsValid() bool {
	switch c {
	case KGS, USD, KZT:
		return true
	}
	return false
}

func (c Currency) Validate() error {
	if !c.IsValid() {
		return fmt.Errorf("unsupported currency: %s", c)
	}
	return nil
}

func Parse(s string) (Currency, error) {
	c := Currency(strings.ToUpper(strings.TrimSpace(s)))
	if !c.IsValid() {
		return DefaultCurrency, fmt.Errorf("unsupported currency: %s", s)
	}
	return c, nil
}

func MustParse(s string) Currency {
	c, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

func (c Currency) Value() (driver.Value, error) {
	return string(c), nil
}

func (c *Currency) Scan(value interface{}) error {
	if value == nil {
		*c = DefaultCurrency
		return nil
	}

	switch v := value.(type) {
	case string:
		*c = Currency(v)
	case []byte:
		*c = Currency(v)
	default:
		return errors.New("invalid type for Currency")
	}

	if !c.IsValid() {
		*c = DefaultCurrency
	}

	return nil
}

func (c Currency) Name() string {
	switch c {
	case KGS:
		return "сом"
	case USD:
		return "доллар"
	case KZT:
		return "тенге"
	default:
		return string(c)
	}
}

func (c Currency) MinorUnits() int {
	return 100
}

func (c Currency) Format(amount int64) string {
	minor := int64(c.MinorUnits())
	whole := amount / minor
	fraction := amount % minor

	if fraction == 0 {
		return fmt.Sprintf("%s %d", c.Name(), whole)
	}

	decimalPlaces := 2
	if c.MinorUnits() == 1 {
		decimalPlaces = 0
	}

	return fmt.Sprintf("%s %d.%0*d", c.Name(), whole, decimalPlaces, fraction)
}
