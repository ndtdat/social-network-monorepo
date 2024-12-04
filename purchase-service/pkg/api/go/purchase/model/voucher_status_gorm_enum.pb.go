// Code generated by protoc-gen-go-gorm-enum. DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"fmt"
)

func (x *VoucherStatus) Scan(value any) error {
	*x = VoucherStatus(VoucherStatus_value[string(value.([]byte))])

	return nil
}

func (x VoucherStatus) Value() (driver.Value, error) {
	return x.String(), nil
}

func VoucherStatus_FromString(str string) (VoucherStatus, error) {
	value, ok := VoucherStatus_value[str]
	if !ok {
		return VoucherStatus(0), fmt.Errorf("cannot parse VoucherStatus from %s", str)
	}

	return VoucherStatus(value), nil
}

func VoucherStatus_MustParseFromString(str string) VoucherStatus {
	value, ok := VoucherStatus_value[str]
	if !ok {
		panic(fmt.Errorf("cannot parse VoucherStatus from %s", str))
	}

	return VoucherStatus(value)
}
