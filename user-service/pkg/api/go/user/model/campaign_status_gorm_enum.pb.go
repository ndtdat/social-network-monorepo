// Code generated by protoc-gen-go-gorm-enum. DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"fmt"
)

func (x *CampaignStatus) Scan(value any) error {
	*x = CampaignStatus(CampaignStatus_value[string(value.([]byte))])

	return nil
}

func (x CampaignStatus) Value() (driver.Value, error) {
	return x.String(), nil
}

func CampaignStatus_FromString(str string) (CampaignStatus, error) {
	value, ok := CampaignStatus_value[str]
	if !ok {
		return CampaignStatus(0), fmt.Errorf("cannot parse CampaignStatus from %s", str)
	}

	return CampaignStatus(value), nil
}

func CampaignStatus_MustParseFromString(str string) CampaignStatus {
	value, ok := CampaignStatus_value[str]
	if !ok {
		panic(fmt.Errorf("cannot parse CampaignStatus from %s", str))
	}

	return CampaignStatus(value)
}
