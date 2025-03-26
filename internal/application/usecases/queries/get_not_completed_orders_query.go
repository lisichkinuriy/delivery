package queries

type GetNotCompletedOrdersQuery struct {
	isSet bool
}

func NewGetNotCompletedOrdersQuery() (GetNotCompletedOrdersQuery, error) {
	return GetNotCompletedOrdersQuery{
		isSet: true,
	}, nil
}

func (q GetNotCompletedOrdersQuery) IsEmpty() bool {
	return !q.isSet
}
