package sqlutil

//nolint:revive
const (
	LogicOperator_NONE LogicOperator = ""
	LogicOperator_AND  LogicOperator = " AND "
	LogicOperator_OR   LogicOperator = " OR "
)

type LogicOperator string

func (l LogicOperator) String() string {
	return string(l)
}
