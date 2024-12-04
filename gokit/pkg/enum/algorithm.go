package enum

import (
	"database/sql/driver"
)

// nolint: revive
const (
	Algorithm_NONE  Algorithm = ""
	Algorithm_ECDSA Algorithm = "ECDSA"
	Algorithm_EDDSA Algorithm = "EDDSA"
)

type Algorithm string

func (e Algorithm) String() string {
	return string(e)
}

func (e *Algorithm) Scan(value any) error {
	*e = Algorithm(value.([]byte)) //nolint: forcetypeassert

	return nil
}

func (e Algorithm) Value() (driver.Value, error) {
	return e.String(), nil
}
