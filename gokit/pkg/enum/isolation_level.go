package enum

import (
	"database/sql/driver"
)

// nolint: revive
const (
	IsolationLevel_NONE             IsolationLevel = ""
	IsolationLevel_READ_UNCOMMITTED IsolationLevel = "READ UNCOMMITTED"
	IsolationLevel_READ_COMMITTED   IsolationLevel = "READ COMMITTED"
	IsolationLevel_REPEATABLE_READ  IsolationLevel = "REPEATABLE READ"
	IsolationLevel_SERIALIZABLE     IsolationLevel = "SERIALIZABLE"
)

type IsolationLevel string

func (e IsolationLevel) String() string {
	return string(e)
}

func (e *IsolationLevel) Scan(value any) error {
	*e = IsolationLevel(value.([]byte)) //nolint: forcetypeassert

	return nil
}

func (e IsolationLevel) Value() (driver.Value, error) {
	return e.String(), nil
}
