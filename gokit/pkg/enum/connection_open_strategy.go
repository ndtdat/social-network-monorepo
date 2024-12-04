package enum

import (
	"database/sql/driver"
)

// nolint: revive
const (
	ConnectionOpenStrategy_NONE        ConnectionOpenStrategy = ""
	ConnectionOpenStrategy_ROUND_ROBIN ConnectionOpenStrategy = "round_robin"
	ConnectionOpenStrategy_IN_ORDER    ConnectionOpenStrategy = "in_order"
)

type ConnectionOpenStrategy string

func (c ConnectionOpenStrategy) String() string {
	return string(c)
}

func (c *ConnectionOpenStrategy) Scan(value any) error {
	*c = ConnectionOpenStrategy(value.([]byte)) //nolint: forcetypeassert

	return nil
}

func (c ConnectionOpenStrategy) Value() (driver.Value, error) {
	return c.String(), nil
}
