package sorter

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/common/pkg/api/go/common"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
)

var orderValues = []string{"asc", "desc"}

type Order string

// nolint: revive
const (
	Order_NONE Order = ""
	Order_ASC  Order = "asc"
	Order_DESC Order = "desc"
)

type Sorter struct {
	Field string
	Order Order
}

func (o Order) String() string {
	return string(o)
}

func NewSorter(field string, order Order) *Sorter {
	return &Sorter{field, order}
}

func SortersFromPb(pbSorters []*common.Sorter, validFields []string) ([]*Sorter, error) {
	var results []*Sorter
	for _, s := range pbSorters {
		field := s.GetField()
		order := s.GetOrder()
		if !util.StringInSlice(order, orderValues) {
			return nil, fmt.Errorf("%s is not a valid sort order", order)
		}

		if !util.StringInSlice(field, validFields) {
			return nil, fmt.Errorf("%s is not a valid sort field", field)
		}

		results = append(results, NewSorter(field, Order(order)))
	}

	return results, nil
}

func SortersToPb(sorters []*Sorter) []*common.Sorter {
	var results []*common.Sorter
	for _, s := range sorters {
		results = append(
			results, &common.Sorter{
				Field: s.Field,
				Order: s.Order.String(),
			},
		)
	}

	return results
}

func (s *Sorter) GetExp() string {
	return fmt.Sprintf("%s %s", s.Field, s.Order)
}

func (s *Sorter) GetExpWithTable(table string) string {
	return fmt.Sprintf("`%s`.`%s` %s", table, s.Field, s.Order)
}

func SortersToOrderBy(sorters []*Sorter) string {
	if len(sorters) == 0 {
		return ""
	}

	var result string
	for _, sort := range sorters {
		exp := sort.GetExp()
		if result != "" {
			result = fmt.Sprintf("%s,%s", result, exp)

			continue
		}

		result += exp
	}

	return fmt.Sprintf("ORDER BY %s", result)
}
