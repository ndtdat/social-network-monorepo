package enum

import (
	"database/sql/driver"
)

// nolint: revive
const (
	Environment_NONE        Environment = ""
	Environment_LOCAL       Environment = "LOCAL"
	Environment_DOCKER      Environment = "DOCKER"
	Environment_DEVELOPMENT Environment = "DEVELOPMENT"
	Environment_STAGING     Environment = "STAGING"
	Environment_PRODUCTION  Environment = "PRODUCTION"
)

type Environment string

func (e Environment) String() string {
	return string(e)
}

func (e *Environment) Scan(value any) error {
	*e = Environment(value.([]byte)) //nolint: forcetypeassert

	return nil
}

func (e Environment) Value() (driver.Value, error) {
	return e.String(), nil
}
