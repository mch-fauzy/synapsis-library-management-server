package models

const (
	SortAsc  = "asc"
	SortDesc = "desc"

	OperatorEqual   = "eq"
	OperatorBetween = "between"
	OperatorIn      = "in"
	OperatorIsNull  = "is_null"
	OperatorNot     = "not"
)

type Filter struct {
	SelectFields []string
	FilterFields []FilterField
	Pagination   Pagination
	Sorts        []Sort
}

type FilterField struct {
	Field    string
	Operator string
	Value    interface{}
}

type Pagination struct {
	Page     int
	PageSize int
}

type Sort struct {
	Field string
	Order string
}
